# ROUTES Directory

## 📖 Filosofi
**"Routes are the map of your API"**. Filosofi routes adalah mendefinisikan semua endpoint API di satu tempat yang terorganisir, sehingga mudah untuk melihat struktur API secara keseluruhan dan mengelola versioning.

## 🎯 Definisi
Directory yang berisi definisi semua HTTP routes/endpoints aplikasi. Routes menghubungkan URL paths dengan handler functions dan middlewares.

## 💡 Tanggung Jawab
- Define URL paths and HTTP methods
- Map routes to handlers
- Apply middlewares to routes
- Group related routes
- API versioning
- Route documentation

## 📝 Contoh Struktur
```
routes/
├── routes.go         # Main routes setup
└── README.md         # Dokumentasi ini
```

## 🔧 Contoh Kode

### routes.go
```go
package routes

import (
    "myproject/handlers"
    "myproject/middlewares"
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(
    app *fiber.App,
    userHandler *handlers.UserHandler,
    productHandler *handlers.ProductHandler,
    authHandler *handlers.AuthHandler,
) {
    // Health check
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status":  "ok",
            "message": "Server is running",
        })
    })

    // API v1 routes
    v1 := app.Group("/api/v1")

    // Public routes
    auth := v1.Group("/auth")
    auth.Post("/register", authHandler.Register)
    auth.Post("/login", authHandler.Login)
    auth.Post("/forgot-password", authHandler.ForgotPassword)
    auth.Post("/reset-password", authHandler.ResetPassword)

    // Protected routes
    protected := v1.Group("", middlewares.AuthMiddleware())

    // User routes
    users := protected.Group("/users")
    users.Get("/", userHandler.GetAll)           // GET /api/v1/users
    users.Get("/:id", userHandler.GetByID)       // GET /api/v1/users/:id
    users.Put("/:id", userHandler.Update)        // PUT /api/v1/users/:id
    users.Delete("/:id", userHandler.Delete)     // DELETE /api/v1/users/:id

    // Product routes
    products := protected.Group("/products")
    products.Get("/", productHandler.GetAll)
    products.Get("/:id", productHandler.GetByID)
    products.Post("/", productHandler.Create)
    products.Put("/:id", productHandler.Update)
    products.Delete("/:id", productHandler.Delete)

    // Profile routes
    profile := protected.Group("/profile")
    profile.Get("/", userHandler.GetProfile)
    profile.Put("/", userHandler.UpdateProfile)
    profile.Put("/password", userHandler.ChangePassword)

    // Admin only routes
    admin := v1.Group("/admin", middlewares.AuthMiddleware(), middlewares.AdminOnly())
    admin.Get("/users", userHandler.GetAll)
    admin.Delete("/users/:id", userHandler.Delete)
    admin.Get("/analytics", userHandler.GetAnalytics)
}
```

## 🎯 Route Patterns & Best Practices

### RESTful Conventions
```
GET    /api/v1/users          → Get all users
GET    /api/v1/users/:id      → Get single user
POST   /api/v1/users          → Create user
PUT    /api/v1/users/:id      → Update user (full)
PATCH  /api/v1/users/:id      → Update user (partial)
DELETE /api/v1/users/:id      → Delete user
```

## ⚠️ Best Practices
- ✅ Use RESTful conventions when possible
- ✅ Version your API (/api/v1, /api/v2)
- ✅ Group related routes
- ✅ Use meaningful URL paths
- ✅ Apply middlewares efficiently
- ✅ Document your routes
- ❌ Don't use verbs in URLs
- ❌ Don't nest too deep

## 🎓 Tips Fiber
- `app.Get()`, `app.Post()`, `app.Put()`, `app.Delete()` untuk HTTP methods
- `app.Group()` untuk group routes
- Middleware applied dengan chaining
- Fiber lebih cepat dan memory efficient
