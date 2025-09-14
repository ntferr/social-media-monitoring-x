package fiber

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/social-media-monitoring-x/internal/config"
	"go.uber.org/dig"
)

func NewServer(cfg *config.AppConfig) *fiber.App {
	return fiber.New(
		fiber.Config{
			AppName:           cfg.Name,
			EnablePrintRoutes: true,
			ErrorHandler:      fiber.DefaultErrorHandler,
			JSONEncoder:       json.Marshal,
			JSONDecoder:       json.Unmarshal,
		},
	).Use(
		cors.New(
			cors.Config{
				AllowOrigins: cors.ConfigDefault.AllowOrigins,
				AllowHeaders: cors.ConfigDefault.AllowHeaders,
				AllowMethods: cors.ConfigDefault.AllowMethods,
			}),
		compress.New(
			compress.Config{
				Level: compress.LevelDefault,
			},
		),
	).(*fiber.App)
}

func Unwrap(d *dig.Container) *fiber.App {
	var result *fiber.App
	d.Invoke(func(f *fiber.App) {
		result = f
	})
	return result
}

func Test(c *fiber.Ctx) error {
	return c.SendString("Healthy")
}
