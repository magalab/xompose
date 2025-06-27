package stack

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/pkg/api"
	"xompose/pkg/compose"
	"xompose/pkg/utils"

	"github.com/docker/docker/client"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sStack) StackStart(ctx context.Context, stackName string) error {
	stack, err := dao.Stack.GetByName(ctx, stackName)
	if err != nil {
		return err
	}
	if stack == nil {
		return gerror.NewCode(gcode.CodeNotFound, "服务不存在")
	}
	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	xCli := compose.NewComposeService(c)
	project, err := utils.YamlToProject(stack.YamlPath)
	if err != nil {
		return err
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if e := xCli.Up(ctx,
			project,
			api.UpOptions{
				Create: api.CreateOptions{},
				Start: api.StartOptions{
					Project:  project,
					Services: []string{},
				},
			}); e != nil {
			return e
		}
		if e := dao.Stack.UpdateStackStatus(ctx, stackName, model.StackRunning); e != nil {
			return e
		}

		return nil
	})
}
