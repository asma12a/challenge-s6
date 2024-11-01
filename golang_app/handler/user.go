package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func UserHandler(app fiber.Router, ctx context.Context, service service.User) {
	app.Get("/", middleware.IsAdminMiddleware, listUsers(ctx, service))
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
				"error":  err.Error(),
			})
		}

		newUser, err := entity.NewUser(
			userInput.Email,
			userInput.Name,
			userInput.Password,
		)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		_, err = service.Create(ctx, newUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		var userInput *entity.User

		if err := c.BodyParser(&userInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user.Name = userInput.Name
		user.Email = userInput.Email
		user.Password = userInput.Password
		user.Roles = userInput.Roles

		updatedUser, err := service.Update(ctx, userInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := service.Delete(ctx, id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
				"error":  err.Error(),
			})
		}

		toJ := make([]presenter.User, len(users))

		for i, user := range users {
			toJ[i] = presenter.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Roles: user.Roles,
			}
		}

		return c.JSON(toJ)
	}
}
