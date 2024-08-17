package api

import "github.com/gofiber/fiber/v2"

type PingApi struct {
}

func NewPingApi() *PingApi {
	return &PingApi{}
}

func (a *PingApi) GetPing(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}
