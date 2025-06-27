// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package log

import (
	"context"

	"xompose/api/log/v1"
)

type ILogV1 interface {
	StreamLog(ctx context.Context, req *v1.StreamLogReq) (res *v1.StreamLogRes, err error)
}
