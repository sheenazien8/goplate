package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sheenazien8/goplate/pkg/inertia"
	inertiaCore "github.com/sheenazien8/inertia-go"
)

type HomeController struct{}

var HomeControllerInstance = &HomeController{}

// Index renders the home page using Inertia.js
func (h *HomeController) Index(c *fiber.Ctx) error {
	props := inertiaCore.Props{
		"message": "Hello from GoPlate with Inertia.js!",
		"user":    map[string]any{"name": "John Doe", "email": "john@example.com"},
	}

	return inertia.FiberManager.Render(c, "Home/Index", props)
}

// About renders the about page using Inertia.js
func (h *HomeController) About(c *fiber.Ctx) error {
	props := inertiaCore.Props{
		"title":       "About GoPlate",
		"description": "GoPlate is a Go-based REST API boilerplate with Inertia.js support.",
	}

	return inertia.FiberManager.Render(c, "Home/About", props)
}
