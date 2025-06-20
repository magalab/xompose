package compose

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"xompose/pkg/api"

	"github.com/compose-spec/compose-go/v2/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/jonboulle/clockwork"
)

var stdioToStdout bool

func init() {
	out, ok := os.LookupEnv("COMPOSE_STATUS_STDOUT")
	if ok {
		stdioToStdout, _ = strconv.ParseBool(out)
	}
}

type composeService struct {
	// dockerCli command.Cli
	cli *client.Client

	clock          clockwork.Clock
	maxConcurrency int
	dryRun         bool
}

// NewComposeService create a local implementation of the compose.Service API
// func NewComposeService(dockerCli command.Cli) api.Service {
func NewComposeService(cli *client.Client) api.Service {
	// c, err := client.NewClientWithOpts(client.FromEnv)
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to create client :%v", err))
	// }

	// project, err := utils.YamlToProject(yamlPath)
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to parse yaml file :%v", err))
	// }

	return &composeService{
		cli: cli,
		// project:        project,
		clock:          clockwork.NewRealClock(),
		maxConcurrency: -1,
		dryRun:         false,
	}
}

func (s *composeService) Close() error {
	var errs []error
	// if s.dockerCli != nil {
	if s.cli != nil {
		errs = append(errs, s.cli.Close())
	}

	return errors.Join(errs...)
}

func (s *composeService) apiClient() client.APIClient {
	return s.cli
}

// func (s *composeService) configFile() *configfile.ConfigFile {
// 	return s.dockerCli.ConfigFile()
// }

func (s *composeService) MaxConcurrency(i int) {
	s.maxConcurrency = i
}

// func (s *composeService) DryRunMode(ctx context.Context, dryRun bool) (context.Context, error) {
// 	s.dryRun = dryRun
// 	if dryRun {
// 		cli, err := command.NewDockerCli()
// 		if err != nil {
// 			return ctx, err
// 		}

// 		options := flags.NewClientOptions()
// 		options.Context = s.dockerCli.CurrentContext()
// 		err = cli.Initialize(options, command.WithInitializeClient(func(cli *command.DockerCli) (client.APIClient, error) {
// 			return api.NewDryRunClient(s.apiClient(), s.dockerCli)
// 		}))
// 		if err != nil {
// 			return ctx, err
// 		}
// 		s.dockerCli = cli
// 	}
// 	return context.WithValue(ctx, api.DryRunKey{}, dryRun), nil
// }

// func (s *composeService) stdout() *streams.Out {
// 	return s.dockerCli.Out()
// }

// func (s *composeService) stdin() *streams.In {
// 	return s.dockerCli.In()
// }

// func (s *composeService) stderr() *streams.Out {
// 	return s.dockerCli.Err()
// }

// func (s *composeService) stdinfo() *streams.Out {
// 	if stdioToStdout {
// 		return s.dockerCli.Out()
// 	}
// 	return s.dockerCli.Err()
// }

func getCanonicalContainerName(c container.Summary) string {
	if len(c.Names) == 0 {
		// corner case, sometime happens on removal. return short ID as a safeguard value
		return c.ID[:12]
	}
	// Names return container canonical name /foo  + link aliases /linked_by/foo
	for _, name := range c.Names {
		if strings.LastIndex(name, "/") == 0 {
			return name[1:]
		}
	}

	return strings.TrimPrefix(c.Names[0], "/")
}

func getContainerNameWithoutProject(c container.Summary) string {
	project := c.Labels[api.ProjectLabel]
	defaultName := getDefaultContainerName(project, c.Labels[api.ServiceLabel], c.Labels[api.ContainerNumberLabel])
	name := getCanonicalContainerName(c)
	if name != defaultName {
		// service declares a custom container_name
		return name
	}
	return name[len(project)+1:]
}

// projectFromName builds a types.Project based on actual resources with compose labels set
func (s *composeService) projectFromName(containers Containers, projectName string, services ...string) (*types.Project, error) {
	project := &types.Project{
		Name:     projectName,
		Services: types.Services{},
	}
	if len(containers) == 0 {
		return project, fmt.Errorf("no container found for project %q: %w", projectName, api.ErrNotFound)
	}
	set := types.Services{}
	for _, c := range containers {
		serviceLabel, ok := c.Labels[api.ServiceLabel]
		if !ok {
			serviceLabel = getCanonicalContainerName(c)
		}
		service, ok := set[serviceLabel]
		if !ok {
			service = types.ServiceConfig{
				Name:   serviceLabel,
				Image:  c.Image,
				Labels: c.Labels,
			}
		}
		service.Scale = increment(service.Scale)
		set[serviceLabel] = service
	}
	for name, service := range set {
		dependencies := service.Labels[api.DependenciesLabel]
		if dependencies != "" {
			service.DependsOn = types.DependsOnConfig{}
			for _, dc := range strings.Split(dependencies, ",") {
				dcArr := strings.Split(dc, ":")
				condition := ServiceConditionRunningOrHealthy
				// Let's restart the dependency by default if we don't have the info stored in the label
				restart := true
				required := true
				dependency := dcArr[0]

				// backward compatibility
				if len(dcArr) > 1 {
					condition = dcArr[1]
					if len(dcArr) > 2 {
						restart, _ = strconv.ParseBool(dcArr[2])
					}
				}
				service.DependsOn[dependency] = types.ServiceDependency{Condition: condition, Restart: restart, Required: required}
			}
			set[name] = service
		}
	}
	project.Services = set

SERVICES:
	for _, qs := range services {
		for _, es := range project.Services {
			if es.Name == qs {
				continue SERVICES
			}
		}
		return project, fmt.Errorf("no such service: %q: %w", qs, api.ErrNotFound)
	}
	project, err := project.WithSelectedServices(services)
	if err != nil {
		return project, err
	}

	return project, nil
}

func increment(scale *int) *int {
	i := 1
	if scale != nil {
		i = *scale + 1
	}
	return &i
}

func (s *composeService) actualVolumes(ctx context.Context, projectName string) (types.Volumes, error) {
	opts := volume.ListOptions{
		Filters: filters.NewArgs(projectFilter(projectName)),
	}
	volumes, err := s.apiClient().VolumeList(ctx, opts)
	if err != nil {
		return nil, err
	}

	actual := types.Volumes{}
	for _, vol := range volumes.Volumes {
		actual[vol.Labels[api.VolumeLabel]] = types.VolumeConfig{
			Name:   vol.Name,
			Driver: vol.Driver,
			Labels: vol.Labels,
		}
	}
	return actual, nil
}

func (s *composeService) actualNetworks(ctx context.Context, projectName string) (types.Networks, error) {
	networks, err := s.apiClient().NetworkList(ctx, network.ListOptions{
		Filters: filters.NewArgs(projectFilter(projectName)),
	})
	if err != nil {
		return nil, err
	}

	actual := types.Networks{}
	for _, net := range networks {
		actual[net.Labels[api.NetworkLabel]] = types.NetworkConfig{
			Name:   net.Name,
			Driver: net.Driver,
			Labels: net.Labels,
		}
	}
	return actual, nil
}

type runtimeVersionCache struct {
	once sync.Once
	val  string
	err  error
}

var runtimeVersion runtimeVersionCache

func (s *composeService) RuntimeVersion(ctx context.Context) (string, error) {
	runtimeVersion.once.Do(func() {
		// version, err := s.dockerCli.Client().ServerVersion(ctx)
		version, err := s.cli.ServerVersion(ctx)
		if err != nil {
			runtimeVersion.err = err
		}
		runtimeVersion.val = version.APIVersion
	})
	return runtimeVersion.val, runtimeVersion.err
}
