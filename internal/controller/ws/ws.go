package ws

import (
	"context"
	"io"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gorilla/websocket"
)

var wsUpGrader = websocket.Upgrader{
	// CheckOrigin allows any origin in development
	// In production, implement proper origin checking for security
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	// Error handler for upgrade failures
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		// Implement error handling logic here
	},
}

var cli *client.Client

func init() {
	var err error
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
}

func TerminalWs(r *ghttp.Request) {
	ws, err := wsUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		r.Response.Write(err.Error())
		return
	}
	defer ws.Close()

	containerName := r.Get("container").String()
	g.Dump(containerName)
	var ctx = r.Context()

	hr, err := exec(containerName, "/root")
	if err != nil {
		glog.Error(ctx, err)
		return
	}
	// 关闭I/O流
	defer hr.Close()
	// 退出进程
	defer func() {
		hr.Conn.Write([]byte("exit\r"))
	}()

	go func() {
		wsWriterCopy(hr.Conn, ws)
	}()
	wsReaderCopy(ws, hr.Conn)
	// Log connection closure
	glog.Info(ctx, "websocket connection closed")
}

func exec(containerName string, workdir string) (hr types.HijackedResponse, err error) {
	// 执行/bin/bash命令
	ir, err := cli.ContainerExecCreate(context.TODO(), containerName, container.ExecOptions{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		WorkingDir:   workdir,
		Cmd:          []string{"/bin/bash"},
		Tty:          true,
	})
	if err != nil {
		return
	}

	// 附加到上面创建的/bin/bash进程中
	hr, err = cli.ContainerExecAttach(context.TODO(), ir.ID, container.ExecStartOptions{Detach: false, Tty: true})
	if err != nil {
		return
	}
	return
}

func wsWriterCopy(reader io.Reader, writer *websocket.Conn) {
	buf := make([]byte, 8192)
	for {
		nr, err := reader.Read(buf)
		if nr > 0 {
			err := writer.WriteMessage(websocket.BinaryMessage, buf[0:nr])
			if err != nil {
				return
			}
		}
		if err != nil {
			return
		}
	}
}

func wsReaderCopy(reader *websocket.Conn, writer io.Writer) {
	for {
		messageType, p, err := reader.ReadMessage()
		g.Dump(messageType, string(p), err)
		if err != nil {
			return
		}
		if messageType == websocket.TextMessage {
			writer.Write(p)
		}
	}
}
