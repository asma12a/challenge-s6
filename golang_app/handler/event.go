package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func EventHandler(app fiber.Router, ctx context.Context, serviceEvent service.Event, serviceEventType service.EventType, serviceSport service.Sport) {
	app.Get("/", listEvents(ctx, serviceEvent))
	app.Get("/:eventId", getEvent(ctx, serviceEvent))
	app.Post("/", createEvent(ctx, serviceEvent, serviceEventType, serviceSport))
	app.Put("/:eventId", updateEvent(ctx, serviceEvent, serviceEventType, serviceSport))
	app.Delete("/:eventId", deleteEvent(ctx, serviceEvent))
}

func createEvent(ctx context.Context, serviceEvent service.Event, serviceEventType service.EventType, serviceSport service.Sport) fiber.Handler {
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
		sportId := eventInput.Edges.Sport.ID
		sport, err := serviceSport.FindOne(ctx, sportId)
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
			&sport.Sport,
		)

		err = serviceEvent.Create(ctx, newEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("eventId")

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

func updateEvent(ctx context.Context, serviceEvent service.Event, serviceEventType service.EventType, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("eventId")

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

		var eventInput entity.Event
		err = c.BodyParser(&eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

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
			existingEvent.Edges.EventType = &eventType.EventType
		}

		// Vérifie le sport
		if eventInput.Edges.Sport != nil {
			sportId := eventInput.Edges.Sport.ID
			sport, err := serviceSport.FindOne(ctx, sportId)
			if err != nil {
				if ent.IsNotFound(err) {
					return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
						"status":       "error",
						"error_detail": "Sport not found",
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": err.Error(),
				})
			}
			existingEvent.Edges.Sport = &sport.Sport
		}

		existingEvent.Name = eventInput.Name
		existingEvent.Address = eventInput.Address
		existingEvent.EventCode = eventInput.EventCode
		existingEvent.Date = eventInput.Date

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
		id := c.Params("eventId")

		err := service.Delete(ctx, id)
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
