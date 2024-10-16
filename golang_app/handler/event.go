package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
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
		var eventInput entity.Event

		err := c.BodyParser(&eventInput)

		if err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		eventTypeId := eventInput.EventTypeID
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

		sportId := eventInput.SportID
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

		newEvent := entity.NewEvent(
			eventInput.Name,
			eventInput.Address,
			eventInput.EventCode,
			eventInput.Date,
			eventType.ID,
			sport.ID,
		)

		err = serviceEvent.Create(ctx, newEvent, eventType.ID, sport.ID)
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

		// Mapper les données de event vers presenter.Event
		toJ := presenter.Event{
			ID:         event.ID,
			Name:       event.Name,
			Address:    event.Address,
			EventCode:  event.EventCode,
			Date:       event.Date,
			CreatedAt:  event.CreatedAt,
			IsPublic:   event.IsPublic,
			IsFinished: event.IsFinished,
			EventType: presenter.EventType{
				ID:   event.Edges.EventType.ID,
				Name: event.Edges.EventType.Name,
			},
			Sport: presenter.Sport{
				ID:       event.Edges.Sport.ID,
				Name:     event.Edges.Sport.Name,
				ImageURL: event.Edges.Sport.ImageURL,
			},
		}

		return c.JSON(toJ)
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

		if eventInput.EventTypeID != "" {
			eventTypeId := eventInput.EventTypeID
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
			existingEvent.EventTypeID = eventType.ID
		}

		if eventInput.SportID != "" {
			sportId := eventInput.SportID
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
			existingEvent.SportID = sport.ID
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

		toJ := make([]presenter.Event, len(events))

		for i, event := range events {
			toJ[i] = presenter.Event{
				ID:         event.ID,
				Name:       event.Name,
				Address:    event.Address,
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
				IsPublic:   event.IsPublic,
				IsFinished: event.IsFinished,
				EventType: presenter.EventType{
					ID:   event.Edges.EventType.ID,
					Name: event.Edges.EventType.Name,
				},
				Sport: presenter.Sport{
					ID:       event.Edges.Sport.ID,
					Name:     event.Edges.Sport.Name,
					ImageURL: event.Edges.Sport.ImageURL, // Optional
				},
			}
		}

		return c.JSON(toJ)
	}
}
