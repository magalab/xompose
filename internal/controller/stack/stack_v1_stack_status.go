package stack

import (
	"context"

	v1 "xompose/api/stack/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StackStatus(ctx context.Context, req *v1.StackStatusReq) (res *v1.StackStatusRes, err error) {
	items, err := service.Stack().StackStatus(ctx, req.StackGetReq)
	if err != nil {
		return nil, err
	}

	return &v1.StackStatusRes{Items: items}, nil
}
