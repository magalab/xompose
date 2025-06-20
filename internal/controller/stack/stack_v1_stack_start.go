package stack

import (
	"context"

	v1 "xompose/api/stack/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StackStart(ctx context.Context, req *v1.StackStartReq) (res *v1.StackStartRes, err error) {
	if err := service.Stack().StackStart(ctx, req.StackName); err != nil {
		return nil, err
	}

	return
}
