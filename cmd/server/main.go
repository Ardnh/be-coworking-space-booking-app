package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Initialize Fiber routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start Fiber server
	app.Listen(":3000")
}
