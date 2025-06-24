package v1

import (
	"xompose/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type SystemSetupReq struct {
	g.Meta `path:"/system/setup" tags:"System" method:"post" sm:"创建管理员用户"`
	*model.SystemSetupReq
}
type SystemSetupRes struct{}
