package main

import (
	app "github.com/thhuang/go-server/internal/apps/nikki"
	"github.com/thhuang/kakao/util/ctx"
)

func main() {
	ctx, cancel := ctx.WithCancel(ctx.Background())
	defer cancel()
	app.New(ctx).Run(ctx)
}
