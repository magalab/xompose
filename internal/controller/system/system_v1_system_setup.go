package system

import (
	"context"

	v1 "xompose/api/system/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) SystemSetup(ctx context.Context, req *v1.SystemSetupReq) (res *v1.SystemSetupRes, err error) {
	if err := service.System().SystemSetup(ctx, req.SystemSetupReq); err != nil {
		return nil, err
	}

	return
}
