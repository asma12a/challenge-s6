package handler

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/asma12a/challenge-s6/api/presenter"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/usecase/event"
	"github.com/gofiber/fiber/v2"
)

func NewEventHandler(app fiber.Router, ctx context.Context, service event.UseCase) {
	app.Post("/", createEvent(ctx, service))
	app.Get("/", listEvents(ctx, service))
	app.Get("/:eventId", getEvent(ctx, service))
	app.Post("/:eventId", updateEvent(ctx, service))
	app.Delete("/:eventId", deleteEvent(ctx, service))
}

func createEvent(ctx context.Context, service event.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var event entity.Event

		// Utilisation d'un décodeur avec l'option DisallowUnknownFields
		decoder := json.NewDecoder(bytes.NewReader(c.Body()))
		decoder.DisallowUnknownFields() // Empêche les clés inconnues

		err := decoder.Decode(&event)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid request body or unknown fields: " + err.Error(),
			})
		}

		// Appel du service pour créer l'événement
		createdEvent, err := service.CreateEvent(ctx, event.Name, event.Address, event.EventCode, event.Date, event.IsPublic, event.IsFinished)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Event{
			ID:         createdEvent.ID,
			Name:       createdEvent.Name,
			Address:    createdEvent.Address,
			EventCode:  createdEvent.EventCode,
			Date:       createdEvent.Date,
			IsPublic:   createdEvent.IsPublic,
			IsFinished: createdEvent.IsFinished,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func getEvent(ctx context.Context, service event.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := entity.StringToID(c.Params("eventId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		event, err := service.GetEvent(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Event{
			ID:         event.ID,
			Name:       event.Name,
			Address:    event.Address,
			EventCode:  event.EventCode,
			Date:       event.Date,
			IsPublic:   event.IsPublic,
			IsFinished: event.IsFinished,
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Event Found",
			"data":    toJ,
		})
	}
}

func updateEvent(ctx context.Context, service event.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("eventId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		var event *entity.Event

		err = c.BodyParser(&event)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		event.ID = id

		event, err = service.UpdateEvent(ctx, event)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := presenter.Event{
			ID:         event.ID,
			Name:       event.Name,
			Address:    event.Address,
			EventCode:  event.EventCode,
			Date:       event.Date,
			IsPublic:   event.IsPublic,
			IsFinished: event.IsFinished,
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   toJ,
			"error":  nil,
		})
	}
}

func deleteEvent(ctx context.Context, service event.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := entity.StringToID(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "Bad Id Format",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		err = service.DeleteEvent(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error deleting Event",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "Event deleted successfully",
			"error":  nil,
		})
	}
}

func listEvents(ctx context.Context, service event.UseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		events, err := service.ListEvents(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		toJ := make([]presenter.Event, len(events))

		for i, event := range events {
			toJ[i] = presenter.Event{
				ID:         event.ID,
				Name:       event.Name,
				Address:    event.Address,
				EventCode:  event.EventCode,
				Date:       event.Date,
				IsPublic:   event.IsPublic,
				IsFinished: event.IsFinished,
			}
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Events Found",
			"data":    toJ,
		})
	}
}
