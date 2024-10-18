package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func TeamHandler(app fiber.Router, ctx context.Context, serviceTeam service.Team) {
	app.Get("/:teamId", getTeam(ctx, serviceTeam))
	app.Post("/", createTeam(ctx, serviceTeam))
	app.Put("/:teamId", updateTeam(ctx, serviceTeam))
	app.Delete("/:teamId", deleteTeam(ctx, serviceTeam))
}


func createTeam(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var teamInput entity.Team

		err := c.BodyParser(&teamInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		newTeam := entity.NewTeam(
			teamInput.Name,
			teamInput.MaxPlayers,
		)

		err = serviceTeam.Create(ctx, newTeam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getTeam(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("teamId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		team, err := serviceTeam.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		return c.JSON(&presenter.Team{
			ID:         team.ID,
			Name:       team.Name,
			MaxPlayers: team.MaxPlayers,
		})
	}
}

func updateTeam(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var teamInput entity.Team

		err := c.BodyParser(&teamInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		id, err := ulid.Parse(c.Params("teamId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		teamInput.ID = id
		team, err := serviceTeam.Update(ctx, &teamInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(&presenter.Team{
			ID:         team.ID,
			Name:       team.Name,
			MaxPlayers: team.MaxPlayers,
		})
	}
}

func deleteTeam(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("teamId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = serviceTeam.Delete(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

