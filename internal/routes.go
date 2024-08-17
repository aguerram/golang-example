package internal

import (
	"github.com/gofiber/fiber/v2"
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal/handler/api"
	"mostafa/learn_go/internal/middleware"
)

func RegisterRoutes(api fiber.Router, handler *api.Handler, app *cmd.App) {
	api.Get("/ping", handler.PingApi.GetPing)

	//auth endpoints
	api.Post("/auth/login", handler.AuthHandler.PostLogin)

	//protected routes
	privateApi := api.Group("/")

	privateApi.Use(middleware.NewJwtAuthMiddleware(app))

	privateApi.Get("/user/me", handler.UserHandler.GetMe)
}
