package stack

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/pkg/api"
	"xompose/pkg/compose"

	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

func (s *sStack) StackList(ctx context.Context, req *model.StackListReq) ([]*model.StackListItem, error) {
	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	xCli := compose.NewComposeService(c)
	stacks, err := xCli.List(ctx, api.ListOptions{All: true})
	if err != nil {
		return nil, err
	}
	items, err := dao.Stack.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	arr := lo.Map(stacks, func(stack api.Stack, _ int) *model.StackListItem {
		return &model.StackListItem{

			StackItem: &model.StackItem{
				// Id:          stack.ID,
				StackName:   stack.Name,
				StackStatus: model.StackStatus(stack.Status),
			},
			IsManaged: lo.ContainsBy(items, func(item *model.StackModel) bool {
				return item.StackName == stack.Name
			}),
		}
	})

	return arr, nil
}
