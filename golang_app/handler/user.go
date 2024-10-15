package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(app fiber.Router, ctx context.Context, service service.User) {
	app.Get("/", listUsers(ctx, service))
	app.Get("/:userId", getUser(ctx, service))
	app.Post("/", createUser(ctx, service))
	app.Put("/:userId", updateUser(ctx, service))
	app.Delete("/:userId", deleteUser(ctx, service))

}

func createUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userInput *entity.User
		err := c.BodyParser(&userInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		newUser, err := entity.NewUser(
			userInput.Email,
			userInput.Name,
			userInput.Password,
			userInput.Role,
		)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		createdUser, err := service.Create(ctx, newUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   createdUser,
			"error":  nil,
		})
	}
}

func getUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("userId")

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User found",
			"data":    user,
		})
	}
}

func updateUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("userId")

		var userInput *entity.User
		err := c.BodyParser(&userInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		user.Name = userInput.Name
		user.Email = userInput.Email
		user.Password = userInput.Password
		user.Role = userInput.Role

		updatedUser, err := service.Update(ctx, userInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User updated",
			"data":    updatedUser,
		})
	}
}

func deleteUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("userId")

		err := service.Delete(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User deleted",
		})
	}
}

func listUsers(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(users)
	}
}
