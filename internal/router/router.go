package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/social-media-monitoring-x/internal/handler"
)

type Router struct {
	App fiber.Router
}

func NewRouter(app fiber.Router) Router {
	return Router{
		App: app,
	}
}

func (r Router) SetupRouter() {
	r.App.Get("/test", handler.Setup)
}
