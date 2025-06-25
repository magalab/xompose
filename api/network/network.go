// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package network

import (
	"context"

	"xompose/api/network/v1"
)

type INetworkV1 interface {
	NetworkList(ctx context.Context, req *v1.NetworkListReq) (res *v1.NetworkListRes, err error)
}
