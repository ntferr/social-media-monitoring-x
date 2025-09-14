package handler

import "github.com/gofiber/fiber/v2"

type App struct {
	*fiber.App
}

// Setup
func (app App) Setup() {
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("application is healthy")
	})
}
