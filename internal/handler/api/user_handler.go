package api

import (
	"github.com/gofiber/fiber/v2"
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal/service"
)

type UserHandler struct {
	app     *cmd.App
	service *service.Service
}

func NewUserHandler(app *cmd.App, service *service.Service) *UserHandler {
	return &UserHandler{
		app,
		service,
	}
}

func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	userId := c.Locals("userId")

	return c.JSON(fiber.Map{
		"userId": userId,
	})
}
