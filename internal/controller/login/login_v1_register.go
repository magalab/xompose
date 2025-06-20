package login

import (
	"context"

	v1 "xompose/api/login/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	_, err = service.Login().Register(ctx, req.RegisterReq)
	if err != nil {
		return nil, err
	}

	return
}
