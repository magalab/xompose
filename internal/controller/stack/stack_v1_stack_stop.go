package stack

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"xompose/api/stack/v1"
)

func (c *ControllerV1) StackStop(ctx context.Context, req *v1.StackStopReq) (res *v1.StackStopRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
