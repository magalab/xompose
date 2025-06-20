package login

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

// UserLogin 用户登录(用户名+密码)
func (s *sLogin) UserLogin(ctx context.Context, req *model.UsernameLoginReq) error {
	u, err := dao.User.GetUserByName(ctx, req.Username)
	if err != nil {
		return err
	}
	if u == nil {
		return gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "密码错误")
	}
	if _, err = s.doLogin(ctx, &model.DoLoginReq{Id: u.Id}); err != nil {
		return gerror.Wrap(err, "Login failed")
	}
	// if _, err := dao.LoginLog.Ctx(ctx).Insert(&do.LoginLog{
	// 	UserId:    u.UserId,
	// 	LoginType: consts.PlatformWeb,
	// }); err != nil {
	// 	return gerror.Wrap(err, "failed to add login log")
	// }

	return nil
}
