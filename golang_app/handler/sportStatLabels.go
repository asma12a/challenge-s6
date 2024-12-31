package handler

import (
	"context"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func SportStatLabelsHandler(app fiber.Router, ctx context.Context, serviceSportStatLables service.SportStatLabels,serviceSport service.Sport, serviceEvent service.Event,serviceUser service.User) {
	app.Post("/:eventId/addUserStat", addUserStat(ctx, serviceSportStatLables, serviceEvent,serviceUser))
	app.Post("/", createSportStatLables(ctx, serviceSportStatLables,serviceSport))
	app.Get("/:sportId", listSportStatLabelsBySport(ctx, serviceSportStatLables))
	app.Get("/", listSportStatLabels(ctx, serviceSportStatLables))

} 



func createSportStatLables(ctx context.Context, serviceSportStatLables service.SportStatLabels,serviceSport service.Sport) fiber.Handler {


	return func(c *fiber.Ctx) error {
		var sportStatLabelInput entity.SportStatLabels	

		err := c.BodyParser(&sportStatLabelInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := validate.Struct(sportStatLabelInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error validate",
				"error": err.Error(),
			})
		}

		sport, err := serviceSport.FindOne(ctx, sportStatLabelInput.SportID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find sport",
				"error":  err.Error(),
			})
		}

		newSportStatLabel := entity.NewSportStatLabels(
			sportStatLabelInput.Label, 
			sportStatLabelInput.Unit, 
			sportStatLabelInput.IsMain, 
			sport.ID,
		)

		err = serviceSportStatLables.Create(c.UserContext(), newSportStatLabel)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error create stat label",
				"error":  err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusCreated)
	}
}

func listSportStatLabels(ctx context.Context, serviceSportStatLables service.SportStatLabels) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sportStatLabels, err := serviceSportStatLables.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error find all",
				"error":  err.Error(),
			})
		}
		toJ := make([]presenter.SportStatLabels, len(sportStatLabels))
		for i, sportStatLabel := range sportStatLabels {
			toJ[i] = presenter.SportStatLabels{
				ID:     sportStatLabel.ID,
				Label:  sportStatLabel.Label,
				Unit:   sportStatLabel.Unit,
				IsMain: sportStatLabel.IsMain,
			}
			if condition := sportStatLabel.Edges.Sport; condition != nil {
				toJ[i].Sport = presenter.Sport{
					ID:   condition.ID,
					Name: condition.Name,
				}
			}
		}
		return c.JSON(toJ)
	}
}

func listSportStatLabelsBySport(ctx context.Context, serviceSportStatLables service.SportStatLabels) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sportIDStr := c.Params("sportId")
		sportID, err := ulid.Parse(sportIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid sport ID format",
			})
		}

		sportStatLabels, err := serviceSportStatLables.FindBySportID(ctx, sportID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error find all",
				"error":  err.Error(),
			})
		}
		toJ := make([]presenter.SportStatLabels, len(sportStatLabels))
		for i, sportStatLabel := range sportStatLabels {
			toJ[i] = presenter.SportStatLabels{
				ID:     sportStatLabel.ID,
				Label:  sportStatLabel.Label,
				Unit:   sportStatLabel.Unit,
				IsMain: sportStatLabel.IsMain,
			}
			if condition := sportStatLabel.Edges.Sport; condition != nil {
				toJ[i].Sport = presenter.Sport{
					ID:   condition.ID,
					Name: condition.Name,
				}
			}
		}
		return c.JSON(toJ)
	}
}


func addUserStat(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceEvent service.Event,serviceUser service.User ) fiber.Handler {
	return func(c *fiber.Ctx) error {

		eventIDStr := c.Params("eventId")
		eventID, err := ulid.Parse(eventIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}
		var userStatInput struct {
			UserID   ulid.ID `json:"user_id" validate:"required"`
			Stats    []struct {
				StatID    ulid.ID `json:"stat_id" validate:"required"`
				StatValue int    `json:"stat_value" validate:"required"`
			} `json:"stats" validate:"required,dive"`
		}

		err = c.BodyParser(&userStatInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := validate.Struct(userStatInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error validate",
				"error": err.Error(),
			})
		}

		event, err := serviceEvent.FindOne(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find event",
				"error":  err.Error(),
			})
		}

		user, err := serviceUser.FindOne(ctx, userStatInput.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find user",
				"error":  err.Error(),
			})
		}

		err = serviceSportStatLables.AddUserStat(ctx, event.ID, user.ID, userStatInput.Stats)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error add user stat",
				"error":  err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusOK)

	}
	
}




