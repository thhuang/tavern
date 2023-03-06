package trytry

import "github.com/gofiber/fiber/v2"


func main() {
	app := fiber.New()

	app.Get(path:"/", func(c *fiber.Ctx) error{
		return c.SendString(body:"Hello World!")
	})
    app.Listern(addr: ":8000")
}