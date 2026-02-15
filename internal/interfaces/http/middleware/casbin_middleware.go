package middleware

import (
	http "github.com/Ardnh/be-coworking-space-booking-app/internal/interfaces/http/responses"
	"github.com/casbin/casbin/v3"
	"github.com/gofiber/fiber/v2"
)

type CasbinMiddleware struct {
	enforcer *casbin.Enforcer
}

func NewCasbinMiddleware(enforcer *casbin.Enforcer) *CasbinMiddleware {
	return &CasbinMiddleware{enforcer: enforcer}
}

func (m *CasbinMiddleware) Authorize() fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Ambil user type dari context (sudah di-set oleh auth middleware)
		userType := c.Locals("user_type")
		if userType == nil {
			return http.NewErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized", nil)
		}

		subject := userType.(string)
		object := c.Path()
		action := c.Method()

		// Enforce policy
		allowed, err := m.enforcer.Enforce(subject, object, action)
		if err != nil {
			return http.NewErrorResponse(c, fiber.StatusInternalServerError, "Failed to check permission", err)
		}

		if !allowed {
			return http.NewErrorResponse(c, fiber.StatusForbidden, "Forbidden", "You don't have permission to access this resource")
		}

		return c.Next()
	}
}
