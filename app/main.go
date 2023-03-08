package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	authController "github.com/thhuang/go-server/app/controllers/auth"
)

var (
	limiterConfig = limiter.Config{Max: 10, Expiration: 30 * time.Second}
)

func main() {
	app := fiber.New()

	// Default middlewares
	app.Use(  logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(limiter.New(limiterConfig))

	// Handlers
	authController.New(app.Group("/auth"))

	log.Fatal(app.Listen(":3000"))
}
