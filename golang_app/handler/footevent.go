package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func FootEventHandler(app fiber.Router, ctx context.Context, service service.FootEvent) {
	app.Get("/", listFootEvents(ctx, service))
	app.Get("/:footEventId", getFootEvent(ctx, service))
	app.Post("/", createFootEvent(ctx, service))
	app.Put("/:footEventId", updateFootEvent(ctx, service))
	app.Delete("/:footEventId", deleteFootEvent(ctx, service))
}

func createFootEvent(ctx context.Context, service service.FootEvent) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var footEventInput *entity.FootEvent
		if err := c.BodyParser(&footEventInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		newFootEvent := entity.NewFootEvent(
			footEventInput.EventID,
			footEventInput.TeamAID,
			footEventInput.TeamBID,
		)

		createdFootEvent, err := service.Create(ctx, newFootEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   createdFootEvent,
		})
	}
}

func getFootEvent(ctx context.Context, service service.FootEvent) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("footEventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		footEvent, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   footEvent,
		})
	}
}

func updateFootEvent(ctx context.Context, service service.FootEvent) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("footEventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		var footEventInput *entity.FootEvent
		if err := c.BodyParser(&footEventInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		// Assurez-vous que l'ID de footEventInput est défini
		footEventInput.ID = id

		updatedFootEvent, err := service.Update(ctx, footEventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status": "success",
			"data":   updatedFootEvent,
		})
	}
}

func deleteFootEvent(ctx context.Context, service service.FootEvent) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("footEventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		if err := service.Delete(ctx, id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		return c.JSON(&fiber.Map{
			"status":  "success",
			"message": "FootEvent deleted",
		})
	}
}

func listFootEvents(ctx context.Context, service service.FootEvent) fiber.Handler {

	return func(c *fiber.Ctx) error {
		footEvents, err := service.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(footEvents)
	}
}
