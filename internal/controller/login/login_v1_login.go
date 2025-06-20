package login

import (
	"context"

	v1 "xompose/api/login/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	if err = service.Login().UserLogin(ctx, req.UsernameLoginReq); err != nil {
		return nil, err
	}

	return
}
