package inertia

import (
	"github.com/gofiber/fiber/v2"
)

// InertiaMiddleware handles Inertia.js requests in Fiber
func InertiaMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if this is an Inertia request
		if c.Get("X-Inertia") == "" {
			return c.Next()
		}

		// Check version mismatch for GET requests
		if c.Method() == "GET" && c.Get("X-Inertia-Version") != "1.0" {
			// Force a full page reload by sending 409 Conflict with Location header
			c.Set("X-Inertia-Location", c.OriginalURL())
			return c.SendStatus(fiber.StatusConflict)
		}

		return c.Next()
	}
}
