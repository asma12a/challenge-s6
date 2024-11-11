package handler

import (
	"context"
	"slices"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
	passwordValidator "github.com/wagslane/go-password-validator"
)

func UserHandler(app fiber.Router, ctx context.Context, service service.User) {
	app.Get("/", middleware.IsAdminMiddleware, listUsers(ctx, service))
	app.Get("/:userId", middleware.IsAdminOrSelfAuthMiddleware, getUser(ctx, service))
	app.Post("/", middleware.IsAdminMiddleware, createUser(ctx, service))
	app.Put("/:userId", middleware.IsAdminOrSelfAuthMiddleware, updateUser(ctx, service))
	app.Delete("/:userId", middleware.IsAdminMiddleware, deleteUser(ctx, service))
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

		data := presenter.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		return c.Status(fiber.StatusOK).JSON(data)
	}
}

func updateUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrInvalidID.Error(),
			})
		}

		var userInput *entity.User

		if err := c.BodyParser(&userInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := validate.Struct(userInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
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

		if userInput.Password != "" {
			if err := passwordValidator.Validate(userInput.Password, 60); err != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
					"status": "error",
					"error":  entity.ErrPasswordNotStrong.Error(),
				})
			}

			hashedPassword, err := user.GeneratePassword(userInput.Password)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}
			user.Password = hashedPassword
		}

		if slices.Contains(user.Roles, "admin") && userInput.Roles != nil {
			user.Roles = userInput.Roles
		}

		updatedUser, err := service.Update(ctx, user)
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
