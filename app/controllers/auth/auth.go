package auth

import (
	"github.com/gofiber/fiber/v2"
)

func New(
	r fiber.Router,
) {
	h := handler{}

	r.Get("/ping", h.ping)
}

type handler struct{}

func (h *handler) ping(c *fiber.Ctx) error {
	return c.JSON(map[string]interface{}{
		"message": "pong",
	})
}
