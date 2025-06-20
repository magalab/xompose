package v1

import (
	"xompose/api"
	"xompose/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// StackAddReq 新建服务
type StackAddReq struct {
	g.Meta `path:"/stack" tags:"Stack" method:"post" sm:"新建服务"`
	api.Auth
	*model.StackAddReq
}
type StackAddRes struct{}

// StackGetReq 获取服务详情
type StackGetReq struct {
	g.Meta `path:"/stack" tags:"Stack" method:"get" sm:"获取服务详情"`
	api.Auth
	*model.StackGetReq
}
type StackGetRes struct {
	*model.StackListItem
}

// StackListReq 获取服务列表
type StackListReq struct {
	g.Meta `path:"/stacks" tags:"Stack" method:"get" sm:"获取服务列表"`
	api.Auth
	*model.StackListReq
}
type StackListRes struct {
	Items []*model.StackListItem `json:"items" dc:"列表"`
}

type StackUpdateReq struct {
	g.Meta `path:"/stack" tags:"Stack" method:"put" sm:"TBD"`
	// 数据
}
type StackUpdateRes struct {
	// 数据
}

// StackDeleteReq 删除服务
type StackDeleteReq struct {
	g.Meta `path:"/stack" tags:"Stack" method:"delete" sm:"删除(软)服务"`
	api.Auth
	*model.StackDeleteReq
}
type StackDeleteRes struct{}

// StackDeployReq 部署服务
type StackDeployReq struct {
	g.Meta    `path:"/stack/deploy" tags:"Stack" method:"post" sm:"部署服务"`
	StackName string `json:"stackName"`
}

type StackDeployRes struct{}

// StackStartReq 启动服务
type StackStartReq struct {
	g.Meta    `path:"/stack/start" tags:"Stack" method:"post" sm:"启动服务"`
	StackName string `json:"stackName"`
}

type StackStartRes struct{}

// StackStopReq 停止服务
type StackStopReq struct {
	g.Meta    `path:"/stack/stop" tags:"Stack" method:"post" sm:"停止服务"`
	StackName string `json:"stackName"`
}

type StackStopRes struct{}

// StackRestartReq 重启服务
type StackRestartReq struct {
	g.Meta    `path:"/stack/restart" tags:"Stack" method:"post" sm:"重启服务"`
	StackName string `json:"stackName"`
}

type StackRestartRes struct{}
