package middleware

import (
	"strings"
	"time"

	"github.com/Ardnh/be-coworking-space-booking-app/internal/config"
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (am *AuthMiddleware) Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: No token provided", nil)
		}

		// Expect header format: "Bearer <token>"
		tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if tokenString == "" {
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token format", nil)
		}

		// Load JWT secret key from config
		cfg := config.LoadConfig()
		secretKey := []byte(cfg.App.JWTSecret)
		if len(secretKey) == 0 {
			return http.NewErrorResponse(c, fiber.StatusInternalServerError, "JWT secret not configured", nil)
		}

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			// Ensure token uses correct signing method
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
			}
			return secretKey, nil
		})

		if err != nil {
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Invalid or expired token", err.Error())
		}

		// Validate claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Optional: check expiry manually
			if exp, ok := claims["exp"].(float64); ok && time.Now().Unix() > int64(exp) {
				return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Token has expired", nil)
			}

			// Store user info in context
			c.Locals("user_id", claims["user_id"])
			c.Locals("user_type", claims["user_type"])
			return c.Next()
		}

		return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Invalid token claims", nil)
	}
}
