package routes

import (
	"github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/handlers"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/middleware"
	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupAPIRoutes(
	app *fiber.App,
	log *logrus.Logger,
	enforcer *casbin.Enforcer,
	userHandler *handlers.UserHandlers,
	authHandler *handlers.AuthHandlers,
	vendorHandler *handlers.VendorHandlers,
) {

	// Middleware
	casbinMiddleware := middleware.NewCasbinMiddleware(enforcer)
	authMiddleware := middleware.NewAuthMiddleware()

	// API v1 group
	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"message": "service is healty",
		})
	})

	// Public routes
	public := api.Group("/")
	{
		// Auth routes
		public.Post("/login", authHandler.Login)
		public.Post("/register", authHandler.Register)
	}

	// Protected routes (require authentication)
	protected := api.Group("/", authMiddleware.Authenticate())
	{
		// User Routes
		users := protected.Group("/users")
		users.Get("/", casbinMiddleware.Authorize(), userHandler.GetAllUser)
		users.Post("/user", casbinMiddleware.Authorize(), userHandler.CreateUser)
		users.Put("/user/:id", casbinMiddleware.Authorize(), userHandler.UpdateUser)
		users.Delete("/user/:id", casbinMiddleware.Authorize(), userHandler.DeleteUser)

		// Vendor Routes
		vendor := protected.Group("/vendors")
		vendor.Get("/", casbinMiddleware.Authorize(), vendorHandler.GetAllVendors)
		vendor.Get("/:vendorId", casbinMiddleware.Authorize(), vendorHandler.GetVendorByID)
		vendor.Get("/:vendorId/resources", casbinMiddleware.Authorize(), vendorHandler.GetVendorsResourcesByVendorID)
		vendor.Post("/", casbinMiddleware.Authorize(), vendorHandler.CreateVendor)
		vendor.Put("/:id", casbinMiddleware.Authorize(), vendorHandler.UpdateVendor)
		vendor.Delete("/:id", casbinMiddleware.Authorize(), vendorHandler.DeleteVendor)
	}

}
