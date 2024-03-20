package app

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	authController "github.com/thhuang/go-server/internal/apps/nikki/controllers/auth"
	"github.com/thhuang/kakao/util/ctx"
)

var (
	limiterConfig = limiter.Config{Max: 10, Expiration: 30 * time.Second}
)

type App struct {
	srv *fiber.App
}

func New(ctx ctx.CTX) *App {
	srv := fiber.New()

	// Default middlewares
	srv.Use(logger.New())
	srv.Use(recover.New())
	srv.Use(requestid.New())
	srv.Use(limiter.New(limiterConfig))

	// Handlers
	authController.New(srv.Group("/auth"))

	return &App{
		srv: srv,
	}
}

func (a *App) Run(ctx ctx.CTX) {
	if err := a.srv.Listen(":3000"); err != nil {
		ctx.WithField("err", err).Fatal("srv.Listen failed")
	}
}
