package stack

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"xompose/api/stack/v1"
)

func (c *ControllerV1) StackDeploy(ctx context.Context, req *v1.StackDeployReq) (res *v1.StackDeployRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
