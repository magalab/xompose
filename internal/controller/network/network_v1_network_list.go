package network

import (
	"context"

	v1 "xompose/api/network/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) NetworkList(ctx context.Context, req *v1.NetworkListReq) (res *v1.NetworkListRes, err error) {
	items, err := service.Network().NetworkList(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.NetworkListRes{Items: items}, nil
}
