# MIDDLEWARES Directory

## 📖 Filosofi
**"Middleware is the gatekeeper of your application"**. Filosofinya adalah memproses requests sebelum sampai ke handler dan responses sebelum dikirim ke client. Middleware memungkinkan separation of concerns untuk cross-cutting concerns seperti authentication, logging, CORS, dll.

## 🎯 Definisi
Directory yang berisi middleware functions yang dieksekusi sebelum atau sesudah handler. Middleware adalah fungsi yang memiliki akses ke context dan next function dalam request-response cycle.

## 💡 Tanggung Jawab
- Authentication & Authorization
- Request logging
- CORS handling
- Rate limiting
- Request validation
- Error handling
- Response compression
- Security headers

## 📝 Contoh Struktur
```
middlewares/
├── auth.go           # Authentication middleware
├── logger.go         # Request logging
├── cors.go           # CORS configuration
├── rate_limiter.go   # Rate limiting
└── README.md         # Dokumentasi ini
```

## 🔧 Contoh Kode

### auth.go
```go
package middlewares

import (
    "strings"
    "myproject/utils"
    "github.com/gofiber/fiber/v2"
)

// AuthMiddleware - Verify JWT token
func AuthMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // 1. Get token from header
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":  "error",
                "message": "Authorization header required",
            })
        }

        // 2. Extract token (format: "Bearer <token>")
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":  "error",
                "message": "Invalid authorization format",
            })
        }

        token := parts[1]

        // 3. Validate token
        claims, err := utils.ValidateJWT(token)
        if err != nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":  "error",
                "message": "Invalid or expired token",
            })
        }

        // 4. Set user info in context (locals)
        c.Locals("user_id", claims.UserID)
        c.Locals("email", claims.Email)
        c.Locals("role", claims.Role)

        // 5. Continue to next handler
        return c.Next()
    }
}

// AdminOnly - Check if user is admin
func AdminOnly() fiber.Handler {
    return func(c *fiber.Ctx) error {
        role := c.Locals("role")
        if role == nil {
            return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "status":  "error",
                "message": "Unauthorized",
            })
        }

        if role != "admin" {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
                "status":  "error",
                "message": "Admin access required",
            })
        }

        return c.Next()
    }
}

// OptionalAuth - Auth optional, tidak block request
func OptionalAuth() fiber.Handler {
    return func(c *fiber.Ctx) error {
        authHeader := c.Get("Authorization")
        if authHeader != "" {
            parts := strings.Split(authHeader, " ")
            if len(parts) == 2 && parts[0] == "Bearer" {
                claims, err := utils.ValidateJWT(parts[1])
                if err == nil {
                    c.Locals("user_id", claims.UserID)
                    c.Locals("email", claims.Email)
                    c.Locals("role", claims.Role)
                }
            }
        }
        return c.Next()
    }
}
```

### logger.go
```go
package middlewares

import (
    "fmt"
    "time"
    "github.com/gofiber/fiber/v2"
)

// CustomLogger - Custom request logger
func CustomLogger() fiber.Handler {
    return func(c *fiber.Ctx) error {
        start := time.Now()

        // Process request
        err := c.Next()

        // Calculate duration
        duration := time.Since(start)

        // Log details
        fmt.Printf(
            "[%s] %s %s | Status: %d | Duration: %v | IP: %s\n",
            c.Method(),
            c.Path(),
            c.Protocol(),
            c.Response().StatusCode(),
            duration,
            c.IP(),
        )

        return err
    }
}
```

### cors.go
```go
package middlewares

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS - Enable CORS with default config
func CORS() fiber.Handler {
    return cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
    })
}

// CORSWithConfig - Configurable CORS
func CORSWithConfig(allowedOrigins []string) fiber.Handler {
    return cors.New(cors.Config{
        AllowOrigins: strings.Join(allowedOrigins, ","),
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, PUT, DELETE, PATCH, OPTIONS",
        AllowCredentials: true,
    })
}
```

### rate_limiter.go
```go
package middlewares

import (
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/limiter"
)

// RateLimit - Limit requests per IP
func RateLimit(max int, expiration time.Duration) fiber.Handler {
    return limiter.New(limiter.Config{
        Max:        max,
        Expiration: expiration,
        KeyGenerator: func(c *fiber.Ctx) string {
            return c.IP()
        },
        LimitReached: func(c *fiber.Ctx) error {
            return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
                "status":  "error",
                "message": "Rate limit exceeded. Please try again later.",
            })
        },
    })
}

// Example usage: RateLimit(100, time.Minute) // 100 requests per minute
```

## 🎯 Middleware Execution Order
```go
app := fiber.New()

// Global middlewares (executed for all routes)
app.Use(recover.New())
app.Use(CustomLogger())
app.Use(CORS())

// Route-specific middleware
api := app.Group("/api/v1")

// Public routes
api.Post("/login", authHandler.Login)
api.Post("/register", authHandler.Register)

// Protected routes (with auth middleware)
protected := api.Group("/users", AuthMiddleware())
protected.Get("/", userHandler.GetAll)
protected.Get("/:id", userHandler.GetByID)

// Admin only routes
admin := api.Group("/admin", AuthMiddleware(), AdminOnly())
admin.Get("/users", userHandler.GetAll)
admin.Delete("/users/:id", userHandler.Delete)
```

## ⚠️ Best Practices
- ✅ Order matters - place auth before authorization
- ✅ Always call `c.Next()` to continue chain
- ✅ Use `c.Locals()` to pass data between middlewares
- ✅ Keep middleware focused on single responsibility
- ✅ Use Fiber's built-in middlewares when available
- ✅ Return error to stop execution
- ❌ Don't put business logic in middleware

## 🎓 Tips Fiber
- `c.Locals("key", value)` untuk set data
- `c.Locals("key")` untuk get data
- `return c.Next()` untuk continue
- Return error untuk stop execution
- Fiber punya banyak built-in middleware (logger, cors, limiter, compress, dll)
