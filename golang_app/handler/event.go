package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func EventHandler(app fiber.Router, ctx context.Context, serviceEvent service.Event, serviceEventType service.EventType) {
	app.Get("/", listEvents(ctx, serviceEvent))
	app.Get("/:eventId", getEvent(ctx, serviceEvent))
	app.Post("/", createEvent(ctx, serviceEvent, serviceEventType))
	app.Put("/:eventId", updateEvent(ctx, serviceEvent, serviceEventType))
	app.Delete("/:eventId", deleteEvent(ctx, serviceEvent))
}

func createEvent(ctx context.Context, serviceEvent service.Event, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var eventInput entity.Event // Utilisation de votre struct Event

		// Parse le corps de la requête JSON
		err := c.BodyParser(&eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		eventTypeId := eventInput.Edges.EventType.ID
		eventType, err := serviceEventType.FindOne(ctx, eventTypeId)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "EventType not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		// Assurez-vous d'utiliser le bon type ici
		newEvent := entity.NewEvent(
			eventInput.Name,
			eventInput.Address,
			eventInput.EventCode,
			eventInput.Date,
			&eventType.EventType,
		)

		err = serviceEvent.Create(ctx, newEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		event, err := service.FindOne(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(event)
	}
}

func updateEvent(ctx context.Context, serviceEvent service.Event, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Récupère l'ID de l'événement à mettre à jour depuis les paramètres de la route
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Cherche l'événement à mettre à jour
		existingEvent, err := serviceEvent.FindOne(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
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

		// Crée une nouvelle instance de EventInput pour récupérer les modifications
		var eventInput entity.Event
		err = c.BodyParser(&eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Si l'utilisateur souhaite modifier le type d'événement
		if eventInput.Edges.EventType != nil {
			eventTypeId := eventInput.Edges.EventType.ID
			eventType, err := serviceEventType.FindOne(ctx, eventTypeId)
			if err != nil {
				if ent.IsNotFound(err) {
					return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
						"status":       "error",
						"error_detail": "EventType not found",
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": err.Error(),
				})
			}
			// Assure-toi d'utiliser le bon type ici
			existingEvent.Edges.EventType = &eventType.EventType // Utilise l'instance ent.EventType
		}

		// Met à jour les autres champs
		existingEvent.Name = eventInput.Name
		existingEvent.Address = eventInput.Address
		existingEvent.EventCode = eventInput.EventCode
		existingEvent.Date = eventInput.Date
		existingEvent.IsPublic = eventInput.IsPublic
		existingEvent.IsFinished = eventInput.IsFinished

		// Enregistre les modifications
		_, err = serviceEvent.Update(ctx, existingEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// id := c.Params("eventId")
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		del_err := service.Delete(ctx, id)
		if del_err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)

	}
}

func listEvents(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		events, err := service.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(events)
	}
}
