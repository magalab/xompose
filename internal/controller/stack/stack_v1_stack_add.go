package stack

import (
	"context"

	v1 "xompose/api/stack/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StackAdd(ctx context.Context, req *v1.StackAddReq) (res *v1.StackAddRes, err error) {
	if err := service.Stack().StackAdd(ctx, req.StackAddReq); err != nil {
		return nil, err
	}

	return
}
