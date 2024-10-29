package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func EventHandler(app fiber.Router, ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) {
	app.Get("/", middleware.IsAuthMiddleware, listEvents(ctx, serviceEvent))
	app.Get("/:eventId", getEvent(ctx, serviceEvent))
	app.Post("/", middleware.IsAuthMiddleware, createEvent(ctx, serviceEvent, serviceSport))
	app.Put("/:eventId", updateEvent(ctx, serviceEvent, serviceSport))
	app.Delete("/:eventId", deleteEvent(ctx, serviceEvent))
}

var validate = validator.New()

func createEvent(ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var eventInput entity.Event

		// Parse le body de la requête
		err := c.BodyParser(&eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrCannotParseJSON.Error(),
			})
		}

		// Valide les champs du JSON
		if err := validate.Struct(eventInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Vérifie si le sport existe à partir de l'ID fourni
		sport, err := serviceSport.FindOne(ctx, eventInput.SportID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrEntityNotFound("Sport").Error(),
			})
		}

		newEvent := entity.NewEvent(
			eventInput.Name,
			eventInput.Address,
			eventInput.EventCode,
			eventInput.Date,
			eventInput.EventType,
			sport.ID,
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
					"status": "error",
					"error":  "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
			EventType:  event.EventType,
			Sport: presenter.Sport{
				ID:       event.Edges.Sport.ID,
				Name:     event.Edges.Sport.Name,
				ImageURL: event.Edges.Sport.ImageURL,
			},
		}

		return c.JSON(toJ)
	}
}

func updateEvent(ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Récupère l'ID de l'événement à mettre à jour depuis les paramètres de la route
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingEvent, err := serviceEvent.FindOne(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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

		if eventInput.SportID != "" {
			sportId := eventInput.SportID
			sport, err := serviceSport.FindOne(ctx, sportId)
			if err != nil {
				if ent.IsNotFound(err) {
					return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
						"status": "error",
						"error":  "Sport not found",
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}
			existingEvent.SportID = sport.ID
		}

		existingEvent.Name = eventInput.Name
		existingEvent.Address = eventInput.Address
		existingEvent.EventCode = eventInput.EventCode
		existingEvent.Date = eventInput.Date
		existingEvent.EventType = eventInput.EventType

		_, err = serviceEvent.Update(ctx, existingEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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

		err = service.Delete(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
				EventType:  event.EventType,
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
