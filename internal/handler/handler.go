package handler

import "github.com/gofiber/fiber/v2"

// Setup
func Setup(ctx *fiber.Ctx) error {
	return ctx.JSON(
		fiber.Map{
			"Application": "is healthy",
		},
	)
}
