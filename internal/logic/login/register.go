package login

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"golang.org/x/crypto/bcrypt"
)

func (s *sLogin) Register(ctx context.Context, req *model.RegisterReq) (*model.UserItem, error) {
	u, err := dao.User.GetUserByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if u != nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "用户名已存在")
	}
	// 密码需要通过 bcrypt 加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	id, err := dao.User.Ctx(ctx).InsertAndGetId(&do.User{
		Username: req.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return &model.UserItem{
		Id:       int(id),
		Username: req.Username,
	}, nil
}
