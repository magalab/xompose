package middleware

import (
	"xompose/internal/consts"
	"xompose/internal/model"
	"xompose/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareHandlerSession(r *ghttp.Request) {
	ctx := r.GetCtx()
	u, err := service.Login().LoginWithSession(ctx)
	if err != nil || u == nil {
		// -1 为特殊标记位
		u = &model.UserItem{Id: -1}
		r.SetCtxVar(consts.CtxUserKey, u)
	} else {
		r.SetCtxVar(consts.CtxUserKey, u)
		r.SetCtxVar(consts.CtxUserIDKey, u.Id)
	}
	r.Middleware.Next()
}
