package api

import (
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal/service"
)

type Handler struct {
	PingApi     *PingApi
	AuthHandler *AuthHandler
	UserHandler *UserHandler
}

func NewHandler(app *cmd.App, service *service.Service) *Handler {
	return &Handler{
		PingApi:     NewPingApi(),
		AuthHandler: NewAuthHandler(app, service),
		UserHandler: NewUserHandler(app, service),
	}
}
