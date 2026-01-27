package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type RequestTimerMiddleware struct {
	logger *logrus.Logger
}

func NewRequestTimerMiddleware(logger *logrus.Logger) *RequestTimerMiddleware {
	return &RequestTimerMiddleware{logger: logger}
}

func (m *RequestTimerMiddleware) Track() fiber.Handler {
	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		c.Locals("request_start_time", startTime)
		err := c.Next()
		duration := time.Since(startTime)

		m.logger.WithFields(logrus.Fields{
			"method":      c.Method(),
			"path":        c.Path(),
			"status":      c.Response().StatusCode(),
			"duration_ms": duration.Milliseconds(),
			"duration":    duration.String(),
			"ip":          c.IP(),
			"user_agent":  c.Get("User-Agent"),
		}).Info("Request completed")

		c.Set("X-Response-Time", duration.String())
		return err
	}
}
