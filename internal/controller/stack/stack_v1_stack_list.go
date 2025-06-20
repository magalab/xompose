package stack

import (
	"context"

	v1 "xompose/api/stack/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StackList(ctx context.Context, req *v1.StackListReq) (res *v1.StackListRes, err error) {
	stacks, err := service.Stack().StackList(ctx, req.StackListReq)
	if err != nil {
		return nil, err
	}

	return &v1.StackListRes{Items: stacks}, nil
}
