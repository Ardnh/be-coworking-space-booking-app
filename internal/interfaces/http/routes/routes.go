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
	resourceHandler *handlers.ResourceHandlers,
	resourceTypeHandler *handlers.ResourceTypeHandlers,
	blockedDateHandler *handlers.BlockedDateHandlers,
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
		users := protected.Group("/users", casbinMiddleware.Authorize())
		users.Get("/", userHandler.GetAllUser)
		users.Get("/user", userHandler.GetUserByToken)
		users.Post("/user", userHandler.CreateUser)
		users.Put("/user/:id", userHandler.UpdateUser)
		users.Delete("/user/:id", userHandler.DeleteUser)

		// Vendor Routes
		vendor := protected.Group("/vendors", casbinMiddleware.Authorize())
		vendor.Get("/", vendorHandler.GetAllVendors)
		vendor.Get("/:vendorId", vendorHandler.GetVendorByID)
		vendor.Get("/:vendorId/resources", vendorHandler.GetVendorsResourcesByVendorID)
		vendor.Post("/", vendorHandler.CreateVendor)
		vendor.Put("/:vendorId", vendorHandler.UpdateVendor)
		vendor.Delete("/:vendorId", vendorHandler.DeleteVendor)

		// Resource Routes
		resources := protected.Group("/resources", casbinMiddleware.Authorize())
		resources.Get("/", resourceHandler.GetResource)
		resources.Get("/:resourceId", resourceHandler.GetResourceById)
		resources.Get("/:resourceId/reviews", resourceHandler.GetReviewsByResourceId)
		resources.Post("/", resourceHandler.CreateResource)
		resources.Put("/:resourceId", resourceHandler.UpdateResource)
		resources.Delete("/:resourceId", resourceHandler.DeleteResource)

		// Resource Type Routes
		resourceTypes := protected.Group("/resource-types", casbinMiddleware.Authorize())
		resourceTypes.Get("/", resourceTypeHandler.GetAllResourceType)
		resourceTypes.Get("/:resourceTypeId", resourceTypeHandler.GetResourceTypeById)
		resourceTypes.Post("/", resourceTypeHandler.CreateResourceType)
		resourceTypes.Put("/:resourceTypeId", resourceTypeHandler.UpdateResourceType)
		resourceTypes.Delete("/:resourceTypeId", resourceTypeHandler.DeleteResourceType)

		// Blocked Date Routes
		blockedDate := protected.Group("/blocked-date", casbinMiddleware.Authorize())
		blockedDate.Get("/", blockedDateHandler.GetBlockedDate)
		blockedDate.Get("/:blockedDateId", blockedDateHandler.GetBlockedDateById)
		blockedDate.Post("/", blockedDateHandler.CreateBlockedDate)
		blockedDate.Put("/:blockedDateId", blockedDateHandler.UpdateBlockedDate)
		blockedDate.Delete("/:blockedDateId", blockedDateHandler.DeleteBlockedDate)
	}

}
