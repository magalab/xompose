package middleware

import (
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
)

func MiddlewareRequestTimeCost(r *ghttp.Request) {
	start := time.Now()
	r.Middleware.Next()

	elapsed := time.Since(start).Milliseconds()
	r.SetCtxVar("elapsed_time", elapsed)
}
