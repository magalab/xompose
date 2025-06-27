package stack

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/pkg/api"
	"xompose/pkg/compose"

	"github.com/docker/docker/client"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
)

func (s *sStack) StackDown(ctx context.Context, req *model.StackGetReq) error {
	stack, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return err
	}
	if stack == nil {
		return gerror.NewCode(gcode.CodeNotFound, "服务不存在")
	}
	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	xCli := compose.NewComposeService(c)
	if err = xCli.Down(ctx, req.StackName, api.DownOptions{}); err != nil {
		return err
	}

	if err = dao.Stack.DeleteByName(ctx, req.StackName); err != nil {
		glog.Error(ctx, "删除 stack 失败", err)

		return nil
	}

	return nil
}
