package api

import (
	"context"
	"xompose/internal/consts"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gvalid"
)

type Auth struct {
	AuthCheck bool `v:"auth-login" json:"-" form:"-"`
}

func authCheck(ctx context.Context, in gvalid.RuleFuncInput) error {
	u := ctx.Value(consts.CtxUserKey).(*model.UserItem)
	if u == nil || !u.IsLogin() {
		return gerror.NewCode(gcode.CodeNotAuthorized, "未登录")
	}

	return nil
}

func init() {
	gvalid.RegisterRule("auth-login", authCheck)
}
