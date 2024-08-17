package cmd

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
	"mostafa/learn_go/internal/model"
	"mostafa/learn_go/internal/type/global"
)

type App struct {
	Env       *AppEnv
	Validator *global.XValidator
	DB        *gorm.DB
}

func NewApp() (*App, error) {
	env, err := loadEnv()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to load environment %v", err))
	}
	if env.AppEnvironment == EnvDevelopment {
		log.Info("environment is development %v", env)
	}
	db, err := setupPostgresqlDB(*env)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("unable to connect to database %v", err))
	}
	if err = db.AutoMigrate(model.Models...); err != nil {
		return nil, errors.New(fmt.Sprintf("unable to migrate models %v", err))
	}

	return &App{
		Env:       env,
		DB:        db,
		Validator: global.NewValidator(),
	}, nil
}
