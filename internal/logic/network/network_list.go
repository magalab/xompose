package user

import (
	"context"
	"xompose/internal/model"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/samber/lo"
)

func (s *sNetwork) NetworkList(ctx context.Context) ([]*model.NetworkItem, error) {
	c, _ := client.NewClientWithOpts(client.FromEnv)
	items, err := c.NetworkList(ctx, network.ListOptions{})
	if err != nil {
		return nil, err
	}
	networks := lo.Map(items, func(item network.Summary, _ int) *model.NetworkItem {
		return &model.NetworkItem{
			Name:   item.Name,
			Scope:  item.Scope,
			Driver: model.Driver(item.Driver),
		}
	})

	return networks, nil
}
