package transport

import (
	"deps_per_service/internal/handler"
	"github.com/gofiber/fiber/v2"
)

type FiberApp struct {
	app *fiber.App
}

func NewFiberApp(handlers *handler.Handlers) *FiberApp {
	app := fiber.New()
	app.Get("/users", handlers.UserHandler.GetUsers)
	return &FiberApp{
		app: app,
	}
}

func (f *FiberApp) Serve() error {
	return f.app.Listen(":3000")
}
