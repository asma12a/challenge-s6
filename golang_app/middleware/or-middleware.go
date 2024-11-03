package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func OrMiddleware(middlewares ...fiber.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var lastErr error
		for _, middleware := range middlewares {
			err := middleware(c)
			if err == nil {
				// Si une middleware réussit, on autorise la requête
				return c.Next()
			}
			lastErr = err
		}
		return lastErr
	}
}
