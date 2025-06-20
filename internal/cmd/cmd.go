package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"xompose/internal/cmd/middleware"
	"xompose/internal/controller/login"
	"xompose/internal/controller/stack"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// s.Logger().SetHandlers(middleware.JsonHandler)
			s.Use(middleware.MiddlewareRequestTimeCost, middleware.MiddlewareCORS, middleware.MiddlewareHandlerSession)
			s.SetSessionMaxAge(time.Hour * 24 * 7)
			s.SetSessionCookieMaxAge(time.Hour * 24 * 7)
			v1Group := s.Group("/api/xompose/v1")
			v1Group.
				Middleware(ghttp.MiddlewareHandlerResponse).
				Bind(
					login.NewV1(),
					stack.NewV1(),
				)
			s.Run()
			return nil
		},
	}
)
