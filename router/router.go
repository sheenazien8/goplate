package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sheenazien8/goplate/middleware"
	"github.com/sheenazien8/goplate/pkg/controllers"
	"github.com/sheenazien8/goplate/pkg/inertia"
)

func SetupRouter(app *fiber.App) {

	app.Use(cors.New())

	// Serve static files from public directory
	app.Static("/", "./public")

	// Add Inertia middleware
	app.Use(inertia.InertiaMiddleware())

	// Inertia.js routes
	var homeController = controllers.HomeControllerInstance
	app.Get("/", homeController.Index)
	app.Get("/about", homeController.About)

	// API routes (keep existing functionality)
	api := app.Group("/api")
	api.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	var logController = controllers.LogControllerInstance
	app.Get("/logs", middleware.BasicAuth(), logController.ShowLogsPage)
}
