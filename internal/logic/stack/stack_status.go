package stack

import (
	"context"
	"strings"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/pkg/api"
	"xompose/pkg/compose"
	"xompose/pkg/utils"

	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

func (s *sStack) StackStatus(ctx context.Context, req *model.StackGetReq) ([]*model.StackStatusItem, error) {
	stack, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return nil, err
	}
	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	xCli := compose.NewComposeService(c)
	project, err := utils.YamlToProject(stack.YamlPath)
	if err != nil {
		return nil, err
	}
	items, err := xCli.Ps(ctx, stack.StackName, api.PsOptions{
		// Project:  project,
		All:      false,
		Services: project.ServiceNames(),
	})
	if err != nil {
		return nil, err
	}
	statusItems := lo.Map(items, func(item api.ContainerSummary, _ int) *model.StackStatusItem {
		label := item.Labels[api.DependenciesLabel]
		depends := lo.Map(strings.Split(label, ","), func(x string, _ int) string {
			return strings.Split(x, ":")[0]
		})
		ports := lo.FilterMap(item.Publishers, func(port api.PortPublisher, _ int) (*model.Port, bool) {
			// filter ipv6
			if port.URL == "::" {
				return nil, false
			}
			return &model.Port{
				PublishPort: port.PublishedPort,
				TargetPort:  port.TargetPort,
				Protocol:    port.Protocol,
			}, true
		})
		return &model.StackStatusItem{
			Image:     item.Image,
			Service:   item.Service,
			State:     model.StackStatus(item.State),
			Ports:     ports,
			Networks:  item.Networks,
			Volumes:   item.Mounts,
			DependsOn: depends,
		}
	})

	return statusItems, nil
}
