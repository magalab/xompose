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
	"github.com/samber/lo"
)

func (s *sStack) StackGet(ctx context.Context, req *model.StackGetReq) (*model.StackListItem, error) {
	item, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return nil, err
	}
	// 不盘空, 有未托管的服务
	c, _ := client.NewClientWithOpts(client.FromEnv)
	xCli := compose.NewComposeService(c)
	stacks, err := xCli.List(ctx, api.ListOptions{All: true})
	if err != nil {
		return nil, err
	}
	// 最多一个
	stack := lo.Filter(stacks, func(item api.Stack, _ int) bool {
		return item.Name == req.StackName
	})
	// docker compose 也没有. 出问题了
	if len(stack) == 0 {
		return nil, gerror.NewCode(gcode.CodeNotFound, "服务不存在")
	}
	// 用 stack 的状态修改
	item.StackStatus = stack[0].Status
	return &model.StackListItem{
		StackItem: item.ToItem(),
		IsManaged: true,
	}, nil
}
