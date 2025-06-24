package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "xompose/api/user/v1"
	"xompose/internal/consts"
	"xompose/internal/model"
	"xompose/internal/service"
)

func (c *ControllerV1) ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {
	u := ctx.Value(consts.CtxUserKey).(*model.UserItem)
	req.Id = u.Id
	if err := service.User().UserChangePassword(ctx, req.ChangePasswordReq); err != nil {
		return nil, gerror.Wrap(err, "")
	}

	return
}
