package main

import (
	_ "goframe-shop/internal/packed"

	_ "goframe-shop/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
