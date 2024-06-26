package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ilkerkorkut/go-examples/microservice/observability/internal/metric"
)

func MetricsMiddleware() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) (err error) {
		metric.TestCounter.Add(1)
		return c.Next()
	}
}
