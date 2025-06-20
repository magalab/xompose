package stack

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"xompose/api/stack/v1"
)

func (c *ControllerV1) StackUpdate(ctx context.Context, req *v1.StackUpdateReq) (res *v1.StackUpdateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
