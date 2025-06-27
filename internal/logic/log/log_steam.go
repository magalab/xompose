package log

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func (s *sLog) LogStream(r *ghttp.Request) {
	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")

	// stackName := r.Get("stackName").String()
	// serviceName := r.Get("serviceName", "").String()
	// client := &Client{
	// 	Name:        fmt.Sprintf("%s%s", stackName, serviceName),
	// 	Request:     r,
	// 	messageChan: make(chan string, 100),
	// }

	r.Response.Writefln("id: %d\nevent: connect\n\n", time.Now().UnixNano())
	r.Response.Flush()

	c, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithAPIVersionNegotiation())

	// logs, _ := c.ContainerLogs(context.Background(), client.Name, container.LogsOptions{
	logs, err := c.ContainerLogs(context.Background(), "ee87290edb36", container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	})
	if err != nil {
		glog.Infof(context.Background(), "get log stream error: %v", err)
		panic(err)
	}

	defer logs.Close()

	logCh := make(chan []byte)
	errCh := make(chan error, 1)

	go func() {
		defer close(logCh)
		var buffer bytes.Buffer
		buf := make([]byte, 4096)
		for {
			n, err := logs.Read(buf)
			if n > 0 {
				buffer.Write(buf[:n])
				for {
					if buffer.Len() < 8 {
						break
					}
					header := buffer.Bytes()[:8]
					dataLen := int(binary.BigEndian.Uint32(header[4:8]))
					if buffer.Len() < 8+dataLen {
						break
					}
					frame := make([]byte, 8+dataLen)
					buffer.Read(frame)
					logCh <- frame
				}
			}
			if err != nil {
				if err != io.EOF {
					errCh <- err
				}
				break
			}

		}
	}()

	for {
		select {
		case data := <-logCh:
			var combinedBuf bytes.Buffer
			_, _ = stdcopy.StdCopy(&combinedBuf, &combinedBuf, bytes.NewReader(data))
			g.Dump(combinedBuf.String())
			r.Response.Writefln("id: %d\nevent: %s\ndata: %s\n\n", time.Now().UnixNano(), "message", combinedBuf)
			r.Response.Flush()
		case <-r.Context().Done():
			glog.Infof(context.Background(), "log stream closed")
			return
		case err := <-errCh:
			glog.Infof(context.Background(), "log stream error: %v", err)
			close(errCh)
			return
		}
	}

}
