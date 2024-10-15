package handler

import (
	"context"
	"log"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func EventHandler(app fiber.Router, ctx context.Context, service service.Event) {
	app.Get("/", listEvents(ctx, service))
	app.Get("/:eventId", getEvent(ctx, service))
	app.Post("/", createEvent(ctx, service))
	app.Put("/:eventId", updateEvent(ctx, service))
	app.Delete("/:eventId", deleteEvent(ctx, service))
}

func createEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var eventInput *entity.Event
		err := c.BodyParser(&eventInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		newEvent := entity.NewEvent(
			eventInput.Name,
			eventInput.Address,
			eventInput.EventCode,
			eventInput.Date,
			eventInput.IsPublic,
			eventInput.IsFinished,
		)

		createdEvent, err := service.Create(ctx, newEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   createdEvent,
			"error":  nil,
		})
	}
}

func getEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := ulid.Parse(c.Params("eventId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User Found",
			"data":    user,
		})
	}
}

func updateEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			log.Printf("Erreur lors de la conversion de l'ID: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		var eventInput *entity.Event
		err = c.BodyParser(&eventInput)
		if err != nil {
			log.Printf("Erreur lors du parsing du corps de la requête: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		eventInput.ID = id

		updatedEvent, err := service.Update(ctx, eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   updatedEvent,
			"error":  nil,
		})
	}
}

func deleteEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		err = service.Delete(ctx, id)
		if err != nil {
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
