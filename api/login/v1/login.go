package v1

import (
	"xompose/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// LoginReq 用户名密码登录
type LoginReq struct {
	g.Meta `path:"/login/username" tags:"Login" method:"post" summary:"账号密码登录"`
	*model.UsernameLoginReq
}

type LoginRes struct{}

// RegisterReq 注册请求
type RegisterReq struct {
	g.Meta `path:"/register" tags:"Register" method:"post" summary:"注册"`
	*model.RegisterReq
}

type RegisterRes struct{}
