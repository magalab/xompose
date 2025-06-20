package stack

import (
	"context"

	v1 "xompose/api/stack/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StackGet(ctx context.Context, req *v1.StackGetReq) (res *v1.StackGetRes, err error) {
	item, err := service.Stack().StackGet(ctx, req.StackGetReq)
	if err != nil {
		return nil, err
	}

	return &v1.StackGetRes{StackListItem: item}, nil
}
