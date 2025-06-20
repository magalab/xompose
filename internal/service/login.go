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
	ILogin interface {
		LoginWithSession(ctx context.Context) (*model.UserItem, error)
		// UserLogin 用户登录(用户名+密码)
		UserLogin(ctx context.Context, req *model.UsernameLoginReq) error
		Register(ctx context.Context, req *model.RegisterReq) (*model.UserItem, error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
