package stack

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"xompose/api/stack/v1"
)

func (c *ControllerV1) StackRestart(ctx context.Context, req *v1.StackRestartReq) (res *v1.StackRestartRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
