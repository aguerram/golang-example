package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal/resource"
	"mostafa/learn_go/internal/service"
	"mostafa/learn_go/internal/type/api_error"
	"mostafa/learn_go/internal/type/request"
	"mostafa/learn_go/internal/util"
)

type AuthHandler struct {
	service *service.Service
	app     *cmd.App
}

func NewAuthHandler(app *cmd.App, service *service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
		app:     app,
	}
}

func (a *AuthHandler) PostLogin(c *fiber.Ctx) error {
	// get post values from request and marshal it to request.LoginRequest
	var loginRequest request.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		return api_error.CannotParseRequest()
	}
	if err := util.ValidateRequest(loginRequest, a.app.Validator); err != nil {
		return err
	}
	login, err := a.service.AuthService.Login(loginRequest)
	if err != nil {
		log.Errorf("error while logging in: %v", err)
		return fiber.NewError(fiber.StatusUnauthorized, resource.MessageInvalidCredentials)
	}
	return c.JSON(login)
}
