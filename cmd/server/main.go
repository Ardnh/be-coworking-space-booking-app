package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	service "github.com/Ardnh/be-coworking-space-booking-app/internal/application/services"
	config "github.com/Ardnh/be-coworking-space-booking-app/internal/config"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/infrastructure/database/postgresql"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/infrastructure/database/redis"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/infrastructure/database/repository"
	handlers "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/handlers"
	middleware "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/middleware"
	"github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/routes"
	logger "github.com/Ardnh/be-coworking-space-booking-app/internal/utils/logger"
	"github.com/go-playground/validator/v10"

	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("🚀 Starting Project Management Backend...")
	log.Println("📝 TODO: Implement application bootstrap")
	log.Println("💡 See cmd/README.md for implementation guidance")

	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory:", err)
	}

	modelPath := filepath.Join(workDir, "internal/config", "casbin_model.conf")
	policyPath := filepath.Join(workDir, "internal/config", "casbin_policy.csv")

	// TODO: Implement application initialization
	// 1. Load configuration
	logger := logger.New()
	cfg := config.LoadConfig()
	validator := validator.New()

	// 2. Initialize database
	db, err := postgresql.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer postgresql.CloseDB(db)

	redisDb := redis.NewRedisDB(cfg)
	defer redisDb.Close()

	// Setup Casbin Enforcer
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		log.Fatal("Failed to create casbin enforcer:", err)
	}

	// 4. Wire up dependencies
	// 5. Start HTTP server
	app := fiber.New()
	requestTimer := middleware.NewRequestTimerMiddleware(logger)

	app.Use(requestTimer.Track())

	// Repository
	userRepository := repository.NewUsersRepository(db, redisDb)
	authRepository := repository.NewAuthRepository(db, redisDb)
	vendorRepository := repository.NewVendorRepository(db, redisDb)
	resourceRepository := repository.NewResourceRespository(db, redisDb)
	resourceTypeRepository := repository.NewResourceType(db, redisDb)
	blockedDateRepository := repository.NewBlockedDateRepository(db, redisDb)

	// Service
	userService := service.NewUserService(userRepository)
	authService := service.NewAuthService(authRepository)
	vendorService := service.NewVendorService(vendorRepository)
	resourceService := service.NewResourceService(resourceRepository)
	resourceTypeService := service.NewResourceTypeService(resourceTypeRepository)
	blockedDateService := service.NewBlockedDateService(blockedDateRepository)

	// Handler
	userHandler := handlers.NewUserHandlers(userService, validator, logger)
	authHandler := handlers.NewAuthHandlers(authService, userService, validator, logger)
	vendorHandler := handlers.NewVendorHandlers(vendorService, validator, logger)
	resourcehandler := handlers.NewResourceHandler(resourceService, validator, logger)
	resourceTypeHandler := handlers.NewResourceTypeHandlers(resourceTypeService, validator, logger)
	blockedDateHandler := handlers.NewBlockedDateHandlers(blockedDateService, validator, logger)

	// Routes
	routes.SetupAPIRoutes(
		app,
		logger,
		enforcer,
		userHandler,
		authHandler,
		vendorHandler,
		resourcehandler,
		resourceTypeHandler,
		blockedDateHandler,
	)

	portListen := fmt.Sprintf(":%s", cfg.App.Port)
	if err := app.Listen(portListen); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
	os.Exit(0)
}
