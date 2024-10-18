package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func EventTeamsHandler(app fiber.Router, ctx context.Context, serviceEventTeams service.EventTeamsService, serviceEvent service.Event, serviceTeam service.Team) {
	app.Get("/", listEventTeams(ctx, serviceEventTeams))
	app.Get("/:eventTeamsId", getEventTeams(ctx, serviceEventTeams))
	app.Post("/", createEventTeams(ctx, serviceEventTeams, serviceEvent, serviceTeam))
	app.Put("/:eventTeamsId", updateEventTeams(ctx, serviceEventTeams, serviceEvent, serviceTeam))
	app.Delete("/:eventTeamsId", deleteEventTeams(ctx, serviceEventTeams))
}

func createEventTeams(ctx context.Context, serviceEventTeams service.EventTeamsService, serviceEvent service.Event, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var eventTeamsInput entity.EventTeams

		// Parse le corps de la requête
		err := c.BodyParser(&eventTeamsInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Vérification de l'existence de l'Event
		event, err := serviceEvent.FindOne(ctx, eventTeamsInput.EventID)
		if err != nil {
			if err == entity.ErrNotFound {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		// Vérification de l'existence de l'équipe (Team)
		team, err := serviceTeam.FindOne(ctx, eventTeamsInput.TeamID)
		if err != nil {
			if err == entity.ErrNotFound {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "Team not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		// Crée l'enregistrement EventTeams
		newEventTeams := entity.NewEventTeams(event.ID, team.ID)
		err = serviceEventTeams.Create(ctx, newEventTeams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getEventTeams(ctx context.Context, serviceEventTeams service.EventTeamsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventTeamsId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		eventTeams, err := serviceEventTeams.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "EventTeams not found",
			})
		}

		toJ := presenter.EventTeams{
			ID:      eventTeams.ID,
			EventID: eventTeams.EventID,
			TeamID:  eventTeams.TeamID,
		}

		return c.JSON(toJ)
	}
}

func updateEventTeams(ctx context.Context, serviceEventTeams service.EventTeamsService, serviceEvent service.Event, serviceTeam service.Team) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventTeamsId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingEventTeams, err := serviceEventTeams.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "EventTeams not found",
			})
		}

		var eventTeamsInput entity.EventTeams
		err = c.BodyParser(&eventTeamsInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Vérification de l'existence de l'Event
		event, err := serviceEvent.FindOne(ctx, eventTeamsInput.EventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "Event not found",
			})
		}

		// Vérification de l'existence de l'équipe (Team)
		team, err := serviceTeam.FindOne(ctx, eventTeamsInput.TeamID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "Team not found",
			})
		}

		existingEventTeams.EventID = event.ID
		existingEventTeams.TeamID = team.ID

		_, err = serviceEventTeams.Update(ctx, existingEventTeams)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteEventTeams(ctx context.Context, serviceEventTeams service.EventTeamsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventTeamsId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = serviceEventTeams.Delete(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "EventTeams not found",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

func listEventTeams(ctx context.Context, service service.EventTeamsService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventTeams, err := service.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.EventTeams, len(eventTeams))

		for i, eventTeam := range eventTeams {
			toJ[i] = presenter.EventTeams{
				ID:      eventTeam.ID,
				EventID: eventTeam.EventID,
				TeamID:  eventTeam.TeamID,
			}
		}
		return c.JSON(toJ)
	}
}
