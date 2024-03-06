package main

import (
	app "github.com/thhuang/go-server/internal/apps/nikki"
	"github.com/thhuang/go-server/utils/ctx"
)

func main() {
	ctx, cancel := ctx.WithCancel(ctx.Backbround())
	defer cancel()
	app.New(ctx).Run(ctx)
}
