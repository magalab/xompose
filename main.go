package main

import (
	_ "xompose/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "xompose/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"xompose/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
