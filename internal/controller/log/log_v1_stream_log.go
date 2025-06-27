package log

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "xompose/api/log/v1"
	"xompose/internal/service"
)

func (c *ControllerV1) StreamLog(ctx context.Context, req *v1.StreamLogReq) (res *v1.StreamLogRes, err error) {
	service.Log().LogStream(g.RequestFromCtx(ctx))

	return
}
