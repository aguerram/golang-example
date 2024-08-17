package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal"
	"mostafa/learn_go/internal/handler/api"
	"mostafa/learn_go/internal/middleware"
	"mostafa/learn_go/internal/service"
	"os/signal"
	"syscall"
)

func main() {
	app, err := cmd.NewApp()
	if err != nil {
		log.Fatalf("Unable to bootstrap the app %v\n", err)
	}
	httpServer := fiber.New(fiber.Config{
		AppName: fmt.Sprintf("%s:%s", app.Env.ApplicationName, app.Env.ApplicationVersion),
	})
	//initiate service
	appService := service.NewService(app)

	// Register routes
	initApiRoutes(httpServer, app, appService)

	// Start the server
	startHttpServer(httpServer, app)
}

func initApiRoutes(server *fiber.App, app *cmd.App, appService *service.Service) {
	apiRouter := server.Group("/api")
	//register middlewares

	apiRouter.Use(logger.New())
	apiRouter.Use(middleware.NewApiErrorHandler())

	// ------------------
	handler := api.NewHandler(app, appService)

	internal.RegisterRoutes(apiRouter, handler, app)
}

func startHttpServer(server *fiber.App, app *cmd.App) {
	//in development, we want to kill running process first
	//if app.Env.AppEnvironment == cmd.EnvDevelopment {
	//	if err := util.KillProcessOnPort(app.Env.Port); err != nil {
	//		log.Fatalf("Unable to kill process on port %d %v\n", app.Env.Port, err)
	//	}
	//}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := server.Listen(fmt.Sprintf(":%d", app.Env.Port)); err != nil {
			log.Fatalf("Unable to start server %v\n", err)
		}
	}()
	<-ctx.Done()
	if err := server.ShutdownWithContext(ctx); err != nil {
		log.Fatalf("Unable to shutdown server %v\n", err)
	}
	log.Info("Server shutdown successfully")
}
