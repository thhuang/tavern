package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(map[string]interface{}{
			"message": "pong",
		})
	})

	log.Fatal(app.Listen(":3000"))
}
