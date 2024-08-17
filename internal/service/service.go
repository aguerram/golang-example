package service

import (
	"mostafa/learn_go/cmd"
)

type Service struct {
	AuthService
}

func NewService(app *cmd.App) *Service {
	return &Service{
		AuthService: NewAuthService(app),
	}
}
