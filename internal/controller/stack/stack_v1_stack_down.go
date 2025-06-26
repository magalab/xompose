package stack

import (
	"context"

	v1 "xompose/api/stack/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StackDown(ctx context.Context, req *v1.StackDownReq) (res *v1.StackDownRes, err error) {
	if err := service.Stack().StackDown(ctx, req.StackGetReq); err != nil {
		return nil, err
	}

	return
}
