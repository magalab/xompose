package stack

import (
	"context"
	"path/filepath"
	"xompose/internal/dao"
	"xompose/internal/model"
	"xompose/internal/model/do"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func (s *sStack) StackAdd(ctx context.Context, req *model.StackAddReq) error {
	stack, err := dao.Stack.GetByName(ctx, req.StackName)
	if err != nil {
		return err
	}
	if stack != nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "服务不能有相同名称")
	}
	stackPath := g.Cfg().MustGet(ctx, "server.stackPath").String()
	if _, err = dao.Stack.Ctx(ctx).InsertAndGetId(&do.Stack{
		StackName:   req.StackName,
		YamlContent: req.YamlContent,
		StackStatus: model.StackCreated,
		YamlPath:    filepath.Join(stackPath, req.StackName, "compose.yml"),
	}); err != nil {
		return err
	}

	return nil
}
