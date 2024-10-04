package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(app fiber.Router, ctx context.Context, service service.User) {
	app.Get("/", listUsers(ctx, service))
}

func listUsers(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.ListUsers(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(users)
	}
}
