package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func ActionLogHandler(app fiber.Router, ctx context.Context, serviceActionLog service.ActionLogService, serviceUser service.User) {
	app.Get("/", listActionLogs(ctx, serviceActionLog))
	app.Post("/", createActionLog(ctx, serviceActionLog, serviceUser))
}

// createActionLog permet de créer un nouveau ActionLog
func createActionLog(ctx context.Context, serviceActionLog service.ActionLogService, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var actionLogInput entity.ActionLog

		err := c.BodyParser(&actionLogInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Vérification de l'existence de l'utilisateur (User)
		user, err := serviceUser.FindOne(ctx, *actionLogInput.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  "User not found",
			})
		}

		// Création d'un nouvel ActionLog
		newActionLog := entity.NewActionLog(
			&user.ID,
			actionLogInput.Action,
			actionLogInput.Description,
		)

		err = serviceActionLog.Create(c.UserContext(), newActionLog)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

// listActionLogs permet de lister tous les ActionLogs
func listActionLogs(ctx context.Context, serviceActionLog service.ActionLogService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		actionLogs, err := serviceActionLog.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(actionLogs)
	}
}
