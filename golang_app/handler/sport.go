package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func SportHandler(app fiber.Router, ctx context.Context, serviceSport service.Sport) {
	app.Get("/", listSports(ctx, serviceSport))
	app.Get("/:sportId", middleware.IsAdminMiddleware, getSport(ctx, serviceSport))
	app.Post("/", middleware.IsAdminMiddleware, createSport(ctx, serviceSport))
	app.Put("/:sportId", middleware.IsAdminMiddleware, updateSport(ctx, serviceSport))
	app.Delete("/:sportId", middleware.IsAdminMiddleware, deleteSport(ctx, serviceSport))
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

		err = serviceSport.Create(c.UserContext(), newSport)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("sportId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		sport, err := serviceSport.FindOne(ctx, id)
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

		toJ := presenter.Sport{
			ID:       sport.ID,
			Name:     sport.Name,
			Type:     presenter.SportType(sport.Type),
			ImageURL: sport.ImageURL,
		}

		return c.JSON(toJ)
	}
}

func updateSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("sportId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingSport, err := serviceSport.FindOne(ctx, id)
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

		_, err = serviceSport.Update(c.UserContext(), existingSport)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteSport(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("sportId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = serviceSport.Delete(ctx, id)
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

func listSports(ctx context.Context, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sports, err := serviceSport.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.Sport, len(sports))

		for i, sport := range sports {
			toJ[i] = presenter.Sport{
				ID:       sport.ID,
				Name:     sport.Name,
				Type:     presenter.SportType(sport.Type),
				ImageURL: sport.ImageURL,
			}
		}
		return c.JSON(toJ)
	}
}
