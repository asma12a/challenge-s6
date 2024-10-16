package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func SportHandler(app fiber.Router, ctx context.Context, serviceSport service.Sport) {
	app.Get("/", listSports(ctx, serviceSport))
	app.Get("/:sportId", getSport(ctx, serviceSport))
	app.Post("/", createSport(ctx, serviceSport))
	app.Put("/:sportId", updateSport(ctx, serviceSport))
	app.Delete("/:sportId", deleteSport(ctx, serviceSport))
}

func createSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var sportInput entity.Sport

		err := c.BodyParser(&sportInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		newSport := entity.NewSport(
			sportInput.Name,
			sportInput.ImageURL,
		)

		err = serviceSport.Create(ctx, newSport)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("sportId")

		sport, err := serviceSport.FindOne(ctx, id)
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

		toJ := presenter.Sport{
			ID:   sport.ID,
			Name: sport.Name,
		}

		return c.JSON(toJ)
	}
}

func updateSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("sportId")

		existingSport, err := serviceSport.FindOne(ctx, id)
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

		var sportInput entity.Sport
		err = c.BodyParser(&sportInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingSport.Name = sportInput.Name
		existingSport.ImageURL = sportInput.ImageURL

		_, err = serviceSport.Update(ctx, existingSport)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("sportId")

		err := serviceSport.Delete(ctx, id)
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

func listSports(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sports, err := serviceSport.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.EventType, len(sports))

		for i, sport := range sports {
			toJ[i] = presenter.EventType{
				ID:   sport.ID,
				Name: sport.Name,
			}
		}
		return c.JSON(toJ)
	}
}
