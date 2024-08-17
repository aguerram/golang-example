package middleware

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"mostafa/learn_go/internal/resource"
	"runtime/debug"
)

func NewApiErrorHandler() fiber.Handler {
	return handleError
}

func handleError(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("panic error %v\n%s", r, debug.Stack())
			err := c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": resource.MessageInternalServerError,
			})
			if err != nil {
				panic(err)
			}
		}
	}()
	err := c.Next()
	if err != nil {
		var fiberError *fiber.Error
		if errors.As(err, &fiberError) {
			return c.Status(fiberError.Code).JSON(fiberError)
		}
		log.Error(resource.MessageInternalServerError, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": resource.MessageInternalServerError,
		})
	}

	return err
}
