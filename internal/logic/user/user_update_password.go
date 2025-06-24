package user

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/crypto/bcrypt"
)

func (s *sUser) UserChangePassword(ctx context.Context, req *model.ChangePasswordReq) error {
	u, err := dao.User.GetUserById(ctx, req.Id)
	if err != nil {
		return gerror.Wrap(err, "")
	}
	if u == nil {
		return gerror.NewCode(gcode.CodeNotFound, "用户不存在")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.CurrentPassword)); err != nil {
		return gerror.NewCode(gcode.CodeInvalidRequest, "密码不正确")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		glog.Error(ctx, err)
		return gerror.NewCode(gcode.CodeInternalError, "")
	}
	if err = dao.User.UpdatePasswordById(ctx, req.Id, string(hashedPassword)); err != nil {
		glog.Error(ctx, err)
		return gerror.NewCode(gcode.CodeInternalError, "")
	}

	return nil
}
