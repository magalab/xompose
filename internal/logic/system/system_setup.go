package system

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

func (s *sSystem) SystemSetup(ctx context.Context, req *model.SystemSetupReq) error {
	count, err := dao.User.Count(ctx)
	if err != nil {
		return gerror.Wrap(err, "")
	}
	if count > 0 {
		return gerror.NewCode(gcode.CodeInvalidRequest, "系统已初始化")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return gerror.Wrap(err, "")
	}
	if _, err = dao.User.Ctx(ctx).Insert(&do.User{
		Username: req.Username,
		Password: hashedPassword,
	}); err != nil {
		return gerror.Wrap(err, "")
	}

	return nil
}
