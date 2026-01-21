package main

import (
	"fmt"
	"log"
	"os"

	config "github.com/Ardnh/be-coworking-space-booking-app/internal/config"
	middleware "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/middleware"
	logger "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("🚀 Starting Project Management Backend...")
	log.Println("📝 TODO: Implement application bootstrap")
	log.Println("💡 See cmd/README.md for implementation guidance")

	// TODO: Implement application initialization
	// 1. Load configuration
	logger := logger.New()
	cfg := config.LoadConfig()

	// 2. Initialize database
	// 3. Setup Redis
	// 4. Wire up dependencies
	// 5. Start HTTP server
	app := fiber.New()

	app.Use(middleware.Logger(logger))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "service is healty",
		})
	})

	portListen := fmt.Sprintf(":%s", cfg.App.Port)
	if err := app.Listen(portListen); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
	os.Exit(0)
}
