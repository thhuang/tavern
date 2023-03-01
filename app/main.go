package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	authApi "github.com/thhuang/go-server/app/handlers/auth"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	authApi.New(app.Group("/auth"))

	log.Fatal(app.Listen(":3000"))
}
