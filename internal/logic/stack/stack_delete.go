package stack

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sStack) StackDelete(ctx context.Context, req *model.StackDeleteReq) error {

	stack, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return err
	}
	if stack == nil {
		return gerror.NewCode(gcode.CodeNotFound, "服务不存在")
	}

	if err = dao.Stack.DeleteByName(ctx, req.StackName); err != nil {
		return gerror.Wrap(err, "删除失败")
	}

	return nil
}
