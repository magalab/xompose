package compose

import (
	"fmt"

	"xompose/pkg/api"

	"github.com/docker/docker/api/types/filters"
)

func projectFilter(projectName string) filters.KeyValuePair {
	return filters.Arg("label", fmt.Sprintf("%s=%s", api.ProjectLabel, projectName))
}

func serviceFilter(serviceName string) filters.KeyValuePair {
	return filters.Arg("label", fmt.Sprintf("%s=%s", api.ServiceLabel, serviceName))
}

func networkFilter(name string) filters.KeyValuePair {
	return filters.Arg("label", fmt.Sprintf("%s=%s", api.NetworkLabel, name))
}

func oneOffFilter(b bool) filters.KeyValuePair {
	v := "False"
	if b {
		v = "True"
	}
	return filters.Arg("label", fmt.Sprintf("%s=%s", api.OneoffLabel, v))
}

func containerNumberFilter(index int) filters.KeyValuePair {
	return filters.Arg("label", fmt.Sprintf("%s=%d", api.ContainerNumberLabel, index))
}

func hasProjectLabelFilter() filters.KeyValuePair {
	return filters.Arg("label", api.ProjectLabel)
}

func hasConfigHashLabel() filters.KeyValuePair {
	return filters.Arg("label", api.ConfigHashLabel)
}
