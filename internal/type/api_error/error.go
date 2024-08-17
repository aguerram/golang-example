package api_error

import (
	"github.com/gofiber/fiber/v2"
	"mostafa/learn_go/internal/resource"
)

func CannotParseRequest() error {
	return &fiber.Error{
		Code:    fiber.StatusBadRequest,
		Message: resource.MessageCannotParseRequest,
	}
}

func Unauthorized() error {
	return &fiber.Error{
		Code:    fiber.StatusUnauthorized,
		Message: resource.MessageUnauthorized,
	}
}
