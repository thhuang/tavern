package auth

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/thhuang/kakao/util/rest"

	"github.com/thhuang/go-server/internal/apps/nikki/models"
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

type handler struct {
	validate *validator.Validate
}

func (h *handler) ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func (h *handler) register(c *fiber.Ctx) error {
	p := models.Auth{}
	if err := rest.ParseBody(c, h.validate, &p); err != nil {
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
		return errors.Wrap(err, "logic::BodyParser")
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
