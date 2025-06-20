package middleware

import "github.com/gogf/gf/v2/net/ghttp"

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"*.xr-oasis.com"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
