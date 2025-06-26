// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"xompose/internal/model"
)

type (
	IStack interface {
		StackDelete(ctx context.Context, req *model.StackDeleteReq) error
		StackDown(ctx context.Context, req *model.StackGetReq) error
		StackGet(ctx context.Context, req *model.StackGetReq) (*model.StackListItem, error)
		StackList(ctx context.Context, req *model.StackListReq) ([]*model.StackListItem, error)
		StackAdd(ctx context.Context, req *model.StackAddReq) error
		StackStart(ctx context.Context, stackName string) error
		StackStatus(ctx context.Context, req *model.StackGetReq) ([]*model.StackStatusItem, error)
		StackStop(ctx context.Context, stackName string) error
	}
)

var (
	localStack IStack
)

func Stack() IStack {
	if localStack == nil {
		panic("implement not found for interface IStack, forgot register?")
	}
	return localStack
}

func RegisterStack(i IStack) {
	localStack = i
}
