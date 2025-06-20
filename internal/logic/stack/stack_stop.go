package stack

import (
	"context"
	"xompose/internal/dao"
	"xompose/pkg/api"
	"xompose/pkg/compose"
	"xompose/pkg/utils"

	"github.com/docker/docker/client"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sStack) StackStop(ctx context.Context, stackName string) error {
	stack, err := dao.Stack.GetByName(ctx, stackName)
	if err != nil {
		return err
	}
	if stack == nil {
		return gerror.NewCode(gcode.CodeNotFound, "服务不存在")
	}
	c, _ := client.NewClientWithOpts(client.FromEnv)
	xCli := compose.NewComposeService(c)
	project, err := utils.YamlToProject(stack.YamlPath)
	if err != nil {
		return err
	}
	if err = xCli.Up(ctx,
		project,
		api.UpOptions{
			Create: api.CreateOptions{},
			Start: api.StartOptions{
				Project:  project,
				Services: []string{},
			},
		}); err != nil {
		return err
	}

	return nil
}
