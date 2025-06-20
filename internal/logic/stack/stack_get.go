package stack

import (
	"context"
	"xompose/internal/dao"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (s *sStack) StackGet(ctx context.Context, req *model.StackGetReq) (*model.StackListItem, error) {

	item, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound, "服务不存在")
	}

	return &model.StackListItem{
		StackItem: item.ToItem(),
		IsManaged: true,
	}, nil
}
