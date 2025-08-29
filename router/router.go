package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sheenazien8/galaplate/middleware"
	"github.com/sheenazien8/galaplate/pkg/controllers"
)

func SetupRouter(app *fiber.App) {

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	var logController = controllers.LogControllerInstance
	app.Get("/logs", middleware.BasicAuth(), logController.ShowLogsPage)
}
