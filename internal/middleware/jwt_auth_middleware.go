package middleware

import (
	"github.com/gofiber/fiber/v2"
	"mostafa/learn_go/cmd"
	"mostafa/learn_go/internal/type/api_error"
	"mostafa/learn_go/internal/util"
	"strings"
)

func NewJwtAuthMiddleware(app *cmd.App) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(fiber.HeaderAuthorization)
		if authorizationHeader == "" {
			return api_error.Unauthorized()
		}
		authValue := strings.Split(authorizationHeader, " ")
		if len(authValue) != 2 || authValue[0] != "Bearer" {
			return api_error.Unauthorized()
		}
		jwtToken := authValue[1]
		userId, err := util.ValidateUserJWT(jwtToken, app.Env.JwtSecret)
		if err != nil {
			return api_error.Unauthorized()
		}
		ctx.Locals("userId", userId)
		return ctx.Next()
	}
}
