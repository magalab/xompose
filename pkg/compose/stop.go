package compose

import (
	"context"
	"slices"
	"strings"

	"xompose/pkg/api"
)

func (s *composeService) Stop(ctx context.Context, projectName string, options api.StopOptions) error {
	// return progress.RunWithTitle(ctx, func(ctx context.Context) error {
	// 	return s.stop(ctx, strings.ToLower(projectName), options)
	// }, s.stdinfo(), "Stopping")
	return s.stop(ctx, strings.ToLower(projectName), options)
}

func (s *composeService) stop(ctx context.Context, projectName string, options api.StopOptions) error {
	containers, err := s.getContainers(ctx, projectName, oneOffExclude, true)
	if err != nil {
		return err
	}

	project := options.Project
	if project == nil {
		project, err = s.getProjectWithResources(ctx, containers, projectName)
		if err != nil {
			return err
		}
	}

	if len(options.Services) == 0 {
		options.Services = project.ServiceNames()
	}

	// w := progress.ContextWriter(ctx)
	return InReverseDependencyOrder(ctx, project, func(c context.Context, service string) error {
		if !slices.Contains(options.Services, service) {
			return nil
		}
		serv := project.Services[service]
		// return s.stopContainers(ctx, w, &serv, containers.filter(isService(service)).filter(isNotOneOff), options.Timeout)
		return s.stopContainers(ctx, &serv, containers.filter(isService(service)).filter(isNotOneOff), options.Timeout)
	})
}
