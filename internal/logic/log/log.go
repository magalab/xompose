package log

import (
	"xompose/internal/service"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Client struct {
	Name        string
	Request     *ghttp.Request
	messageChan chan string
}

type sLog struct {
	clients *gmap.StrAnyMap
}

func init() {
	service.RegisterLog(New())
}

func New() *sLog {
	return &sLog{
		clients: gmap.NewStrAnyMap(true),
	}
}
