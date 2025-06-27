package stack

import (
	"context"
	"path/filepath"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/internal/model/do"
	"xompose/pkg/api"
	"xompose/pkg/compose"
	"xompose/pkg/utils"

	"github.com/docker/docker/client"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sStack) StackDeploy(ctx context.Context, req *model.StackAddReq) error {
	stack, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return err
	}
	if stack != nil {
		return gerror.NewCode(gcode.CodeInvalidRequest, "同名服务已存在")
	}
	stackPath := g.Cfg().MustGet(ctx, "server.stackPath").String()
	if _, err = dao.Stack.Ctx(ctx).InsertAndGetId(&do.Stack{
		StackName:   req.StackName,
		YamlContent: req.YamlContent,
		StackStatus: model.StackCreated,
		YamlPath:    filepath.Join(stackPath, req.StackName, "compose.yml"),
	}); err != nil {
		return err
	}
	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	xCli := compose.NewComposeService(c)
	project, err := utils.YamlToProject(stackPath)
	if err != nil {
		return err
	}
	if err = xCli.Up(ctx, project, api.UpOptions{
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
