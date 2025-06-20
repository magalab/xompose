package stack

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "xompose/api/stack/v1"
)

func (c *ControllerV1) StackDelete(ctx context.Context, req *v1.StackDeleteReq) (res *v1.StackDeleteRes, err error) {

	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
