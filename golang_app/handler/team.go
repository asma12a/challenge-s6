package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
)

func TeamHandler(app fiber.Router, ctx context.Context, serviceTeam service.Team) {
	// User interaction with teams
	app.Post("/:teamId/join", joinTeam(ctx, serviceTeam))

	// "/:eventId/teams" scoped
	app.Get("/", listEventTeams(ctx, serviceTeam))
	app.Get("/:teamId", getTeam(ctx, serviceTeam))
	app.Post("/", addTeam(ctx, serviceTeam))
	app.Put("/:teamId", updateTeam(ctx, serviceTeam))
	app.Delete("/:teamId", deleteTeam(ctx, serviceTeam))
}

func listEventTeams(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventID, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}

		teams, err := serviceTeam.FindAll(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		toJ := make(([]presenter.Team), len(teams))
		for i, team := range teams {
			toJ[i] = presenter.Team{
				ID:         team.ID,
				Name:       team.Name,
				MaxPlayers: team.MaxPlayers,
			}

			for _, player := range team.Edges.TeamUsers {
				user := player.Edges.User
				toJ[i].Players = append(toJ[i].Players, presenter.Player{
					ID:     player.ID,
					Name:   user.Name,
					Email:  user.Email,
					Role:   presenter.Role(player.Role),
					Status: presenter.Status(player.Status),
				})
			}
		}

		return c.JSON(toJ)
	}
}

func addTeam(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventID, err := ulid.Parse(c.Params("eventId"))

		var teamsInput entity.Team
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}

		err = c.BodyParser(&teamsInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error bodyparser",
				"error":  err.Error(),
			})
		}

		err = serviceTeam.AddTeam(ctx, eventID, teamsInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error add team",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
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
		team, err := serviceTeam.Update(c.UserContext(), &teamInput)
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

func joinTeam(ctx context.Context, serviceTeam service.Team) fiber.Handler {
	// user join a event, byt joining a team in the event, must check if user isn't already in a team of the event
	return func(c *fiber.Ctx) error {
		currentUser, err := viewer.UserFromContext(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		eventIDStr := c.Params("eventId")
		teamIDStr := c.Params("teamId")

		eventID, err := ulid.Parse(eventIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}

		teamID, err := ulid.Parse(teamIDStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid team ID format",
			})
		}

		err = serviceTeam.JoinTeam(ctx, eventID, teamID, currentUser.ID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
