package service

import (
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal/model"
	"mostafa/learn_go/internal/type/request"
	"mostafa/learn_go/internal/type/response"
	"mostafa/learn_go/internal/util"
)

type AuthService interface {
	Login(r request.LoginRequest) (response.LoginResponse, error)
}

type AuthServiceImpl struct {
	app *cmd.App
}

func NewAuthService(app *cmd.App) *AuthServiceImpl {
	return &AuthServiceImpl{
		app: app,
	}
}

func (s *AuthServiceImpl) Login(r request.LoginRequest) (response.LoginResponse, error) {
	var user *model.User
	tx := s.app.DB.Model(&model.User{}).Where("username = ?", r.Username).First(&user)

	if tx.Error != nil {
		return response.LoginResponse{}, tx.Error
	}
	if user == nil {
		return response.LoginResponse{}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		log.Debugf("Password of username %s is incorretc", r.Username)
		return response.LoginResponse{}, err
	}

	jwt, err := util.GenerateUserJWT(user, s.app.Env.JwtSecret)
	if err != nil {
		return response.LoginResponse{}, err
	}
	return response.LoginResponse{
		Token: jwt,
	}, nil
}
