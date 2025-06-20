package login

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (s *sLogin) LoginWithSession(ctx context.Context) (*model.UserItem, error) {
	r := g.RequestFromCtx(ctx)
	userVar, err := r.Session.Get("user_id", "")
	if err != nil {
		return nil, gerror.WrapCodef(gcode.CodeNotAuthorized, err, "session get user failed")
	}
	if userVar.IsEmpty() {
		return nil, gerror.WrapCodef(gcode.CodeNotAuthorized, err, "user not login")
	}
	u, err := dao.User.GetUserById(ctx, userVar.Int())
	if err != nil {
		return nil, gerror.Wrap(err, "login failed")
	}
	if u == nil {
		if err = r.Session.RemoveAll(); err != nil {
			return nil, gerror.Wrap(err, "session remove failed")
		}

		return nil, gerror.Wrap(err, "user not found")
	}
	// 更新 session
	if err := r.Session.Set("time", gtime.Timestamp()); err != nil {
		return nil, gerror.Wrap(err, "session set time failed")
	}
	if err := r.Session.Set("user_id", u.Id); err != nil {
		return nil, gerror.Wrap(err, "session set user_id failed")
	}
	r.Cookie.SetSessionId(r.Session.MustId())

	return u.ToItem(), nil

}

func (s *sLogin) doLogin(ctx context.Context, req *model.DoLoginReq) (*model.UserItem, error) {
	r := g.RequestFromCtx(ctx)
	userVar, err := r.Session.Get("user_id", "")
	if err != nil {
		return nil, gerror.Wrap(err, "session get user failed")
	}
	// 登录中, 更新 session
	if !userVar.IsEmpty() {
		return s.LoginWithSession(ctx)
	}
	u, err := dao.User.GetUserById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "login failed")
	}
	// 更新 session
	if err := r.Session.Set("time", gtime.Timestamp()); err != nil {
		return nil, gerror.Wrap(err, "session set time failed")
	}
	if err := r.Session.Set("user_id", u.Id); err != nil {
		return nil, gerror.Wrap(err, "session set user_id failed")
	}
	r.Cookie.SetSessionId(r.Session.MustId())

	return u.ToItem(), nil
}
