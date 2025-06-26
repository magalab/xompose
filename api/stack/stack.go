// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package stack

import (
	"context"

	"xompose/api/stack/v1"
)

type IStackV1 interface {
	StackAdd(ctx context.Context, req *v1.StackAddReq) (res *v1.StackAddRes, err error)
	StackGet(ctx context.Context, req *v1.StackGetReq) (res *v1.StackGetRes, err error)
	StackList(ctx context.Context, req *v1.StackListReq) (res *v1.StackListRes, err error)
	StackUpdate(ctx context.Context, req *v1.StackUpdateReq) (res *v1.StackUpdateRes, err error)
	StackDelete(ctx context.Context, req *v1.StackDeleteReq) (res *v1.StackDeleteRes, err error)
	StackDeploy(ctx context.Context, req *v1.StackDeployReq) (res *v1.StackDeployRes, err error)
	StackStart(ctx context.Context, req *v1.StackStartReq) (res *v1.StackStartRes, err error)
	StackStop(ctx context.Context, req *v1.StackStopReq) (res *v1.StackStopRes, err error)
	StackDown(ctx context.Context, req *v1.StackDownReq) (res *v1.StackDownRes, err error)
	StackRestart(ctx context.Context, req *v1.StackRestartReq) (res *v1.StackRestartRes, err error)
	StackStatus(ctx context.Context, req *v1.StackStatusReq) (res *v1.StackStatusRes, err error)
}
