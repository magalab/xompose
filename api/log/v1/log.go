package v1

import "github.com/gogf/gf/v2/frame/g"

type StreamLogReq struct {
	g.Meta      `path:"/log/stream" method:"get" tags:"Log" sm:"流式日志"`
	StackName   string `json:"stackName" v:"required#请输入stack名称" dc:"stack名称"`
	ServiceName string `json:"serviceName" dc:"service名称"`
}

type StreamLogRes struct{}
