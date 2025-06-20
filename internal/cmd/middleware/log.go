package middleware

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// HandlerOutputJson is the structure outputting logging content as single json.
type HandlerOutputJson struct {
	Time       string                 `json:""`           // Formatted time string, like "2016-01-09 12:00:00".
	TraceId    string                 `json:",omitempty"` // Trace id, only available if tracing is enabled.
	CtxStr     map[string]interface{} `json:",omitempty"` // The retrieved context value string from context, only available if Config.CtxKeys configured.
	Level      string                 `json:""`           // Formatted level string, like "DEBU", "ERRO", etc. Eg: ERRO
	CallerFunc string                 `json:",omitempty"` // The source function name that calls logging, only available if F_CALLER_FN set.
	CallerPath string                 `json:",omitempty"` // The source file path and its line number that calls logging, only available if F_FILE_SHORT or F_FILE_LONG set.
	Prefix     string                 `json:",omitempty"` // Custom prefix string for logging content.
	Content    string                 `json:""`           // Content is the main logging content, containing error stack string produced by logger.
	Stack      string                 `json:",omitempty"` // Stack string produced by logger, only available if Config.StStatus configured.
	Uri        string                 `json:",omitempty"`
	Status     int                    `json:",omitempty"`
	ClientIp   string                 `json:",omitempty"`
	RemoteIp   string                 `json:",omitempty"`
	Host       string                 `json:",omitempty"`
	Params     *gjson.Json            `json:",omitempty"`
	URL        string                 `json:",omitempty"`
}

func JsonHandler(ctx context.Context, in *glog.HandlerInput) {
	output := HandlerOutputJson{
		Time:       in.TimeFormat,
		TraceId:    in.TraceId,
		Level:      in.LevelFormat,
		CallerFunc: in.CallerFunc,
		CallerPath: in.CallerPath,
		Prefix:     in.Prefix,
		Content:    in.Content,
		Stack:      in.Stack,
	}
	r := g.RequestFromCtx(ctx)
	if r != nil {
		rj, _ := r.GetJson()
		if rj != nil {
			output.Params = rj
		}
		output.URL = r.RequestURI
		output.CtxStr = getCtxMap(ctx)
		output.Uri = r.Request.URL.Path
		output.Status = r.Response.Status
		output.Host = r.Host
		output.ClientIp = r.GetClientIp()
		output.RemoteIp = r.GetRemoteIp()
	}
	jsonBytes, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}
	in.Buffer.Write(jsonBytes)
	in.Buffer.Write([]byte("\n"))
	in.Next(ctx)
}

func getCtxMap(ctx context.Context) map[string]interface{} {
	if ctx == nil {
		return nil
	}
	ret := map[string]interface{}{}
	for _, ctxKey := range []string{"elapsed_time", "user_id"} {
		var ctxValue interface{}
		if ctxValue = ctx.Value(ctxKey); ctxValue == nil {
			ctxValue = ctx.Value(gctx.StrKey(gconv.String(ctxKey)))
		}
		ret[ctxKey] = ctxValue
	}

	return ret
}
