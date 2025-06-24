// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"xompose/api/system/v1"
)

type ISystemV1 interface {
	SystemSetup(ctx context.Context, req *v1.SystemSetupReq) (res *v1.SystemSetupRes, err error)
}
