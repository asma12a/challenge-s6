package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func EventTypeHandler(app fiber.Router, ctx context.Context, serviceEventType service.EventType) {
	app.Get("/", listEventTypes(ctx, serviceEventType))
	app.Get("/:eventTypeId", getEventType(ctx, serviceEventType))
	app.Post("/", createEventType(ctx, serviceEventType))
	app.Put("/:eventTypeId", updateEventType(ctx, serviceEventType))
	app.Delete("/:eventTypeId", deleteEventType(ctx, serviceEventType))
}

func createEventType(ctx context.Context, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var eventTypeInput entity.EventType

		err := c.BodyParser(&eventTypeInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		newEventType := entity.NewEventType(
			eventTypeInput.Name,
		)

		err = serviceEventType.Create(ctx, newEventType)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getEventType(ctx context.Context, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("eventTypeId")

		eventType, err := serviceEventType.FindOne(ctx, id)
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

		return c.JSON(eventType)
	}
}

func updateEventType(ctx context.Context, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("eventTypeId")

		existingEventType, err := serviceEventType.FindOne(ctx, id)
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

		var eventTypeInput entity.EventType
		err = c.BodyParser(&eventTypeInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingEventType.Name = eventTypeInput.Name

		_, err = serviceEventType.Update(ctx, existingEventType)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteEventType(ctx context.Context, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("eventTypeId")

		err := serviceEventType.Delete(ctx, id)
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

func listEventTypes(ctx context.Context, serviceEventType service.EventType) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventTypes, err := serviceEventType.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(eventTypes)
	}
}
