package utils

import (
	"context"
	"strings"
	"xompose/pkg/api"

	"github.com/compose-spec/compose-go/v2/cli"
	"github.com/compose-spec/compose-go/v2/types"
	"github.com/gogf/gf/v2/errors/gerror"
)

func YamlToProject(yamlPath string) (*types.Project, error) {
	options, err := cli.NewProjectOptions(
		[]string{yamlPath},
		cli.WithOsEnv,
		cli.WithDotEnv,
	)
	if err != nil {
		return nil, gerror.Wrap(err, "YamlToProject")
	}
	project, err := cli.ProjectFromOptions(context.Background(), options)
	if err != nil {
		return nil, gerror.Wrap(err, "YamlToProject")
	}
	for name, s := range project.Services {
		s.CustomLabels = map[string]string{
			api.ProjectLabel:     project.Name,
			api.ServiceLabel:     name,
			api.VersionLabel:     api.ComposeVersion,
			api.WorkingDirLabel:  project.WorkingDir,
			api.ConfigFilesLabel: strings.Join(project.ComposeFiles, ","),
			api.OneoffLabel:      "False",
		}
		project.Services[name] = s
	}

	return project, nil
}
