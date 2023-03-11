package auth

import (
	"github.com/gofiber/fiber/v2"
	authEntity "github.com/thhuang/go-server/app/entities/auth"
	apiUtils "github.com/thhuang/go-server/utils/api"
	errorModel "github.com/thhuang/go-server/utils/error"
)

func New(
	r fiber.Router,
) {
	h := handler{}

	r.Get("/ping", h.ping)
	r.Post("/register", h.register)
	r.Post("/login", h.login)
}

type handler struct{}

func (h *handler) ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func (h *handler) register(c *fiber.Ctx) error {
	p := authEntity.UserAuth{}
	if err := apiUtils.ParseBody(c, &p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorModel.ErrorResponse{
			Code:    errorModel.ErrorCodeUnknown,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusInternalServerError).JSON(errorModel.ErrorResponse{
		Code: errorModel.ErrorCodeNotImplemented,
	})
}

func (h *handler) login(c *fiber.Ctx) error {
	payload := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	// return c.JSON(payload)
	if payload.Name == "123" {

		return c.JSON(map[string]interface{}{
			"message": "login works",
		})
	} else {

		return c.JSON(map[string]interface{}{
			"message": "login fail"})
	}

}
