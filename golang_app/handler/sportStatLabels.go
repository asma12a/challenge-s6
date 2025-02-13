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
	"github.com/redis/go-redis/v9"
)

func SportStatLabelsHandler(app fiber.Router, ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceSport service.Sport, serviceEvent service.Event, serviceUser service.User, serviceNotification service.NotificationService, rdb *redis.Client) {
	app.Post("/:eventId/addUserStat", middleware.IsEventOrganizerOrCoach(ctx, serviceEvent), addUserStat(ctx, serviceSportStatLables, serviceEvent, serviceUser, serviceNotification, rdb))
	app.Post("/", middleware.IsAdminMiddleware, createSportStatLables(ctx, serviceSportStatLables, serviceSport))
	app.Get("/:eventId/:userId/stats", middleware.IsAuthMiddleware, getUserStatsByEvent(ctx, serviceSportStatLables, serviceEvent, serviceUser))
	app.Get("/:eventId/stats", middleware.IsAuthMiddleware, getAllTeamUserMainStatsByEvent(ctx, serviceSportStatLables, serviceEvent,))
	app.Get("/:sportId/mainStat", middleware.IsAuthMiddleware, findMainStatLabelBySportID(ctx, serviceSportStatLables))
	app.Get("/:sportId/:userId/performance", middleware.IsAuthMiddleware, getUserPerformanceBySport(ctx, serviceSportStatLables, serviceUser))
	app.Get("/:sportId/labels", middleware.IsAuthMiddleware, listSportStatLabelsBySport(ctx, serviceSportStatLables))
	app.Get("/:sportStatLabelId", middleware.IsAdminMiddleware, getSportStatLabel(ctx, serviceSportStatLables))
	app.Get("/", middleware.IsAdminMiddleware, listSportStatLabels(ctx, serviceSportStatLables))
	app.Delete("/:sportStatLabelId", middleware.IsAdminMiddleware, deleteSportStatLabel(ctx, serviceSportStatLables))
	app.Put("/:eventId/updateUserStats", middleware.IsEventOrganizerOrCoach(ctx, serviceEvent), updateUserStats(ctx, serviceSportStatLables, serviceNotification, rdb))
	app.Put("/:sportStatLabelId", middleware.IsAdminMiddleware, updateSportStatLabel(ctx, serviceSportStatLables, serviceSport))

}

func createSportStatLables(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceSport service.Sport) fiber.Handler {

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
				"error":  err.Error(),
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
				toJ[i].Sport = &presenter.Sport{
					ID:   condition.ID,
					Name: condition.Name,
				}
			}
		}
		return c.JSON(toJ)
	}
}

func getSportStatLabel(ctx context.Context, serviceSportStatLables service.SportStatLabels) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sportStatLabelIDStr := c.Params("sportStatLabelId")
		sportStatLabelID, err := ulid.Parse(sportStatLabelIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		sportStatLabel, err := serviceSportStatLables.FindOne(ctx, sportStatLabelID)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "SportStatLabel not found",
				})
			}

			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error find one",
				"error":  err.Error(),
			})
		}

		toJ := presenter.SportStatLabels{
			ID:     sportStatLabel.ID,
			Label:  sportStatLabel.Label,
			Unit:   sportStatLabel.Unit,
			IsMain: sportStatLabel.IsMain,
		}
		if condition := sportStatLabel.Edges.Sport; condition != nil {
			toJ.Sport = &presenter.Sport{
				ID:   condition.ID,
				Name: condition.Name,
			}
		}
		return c.JSON(toJ)
	}
}

func updateSportStatLabel(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sportStatLabelIDStr := c.Params("sportStatLabelId")
		sportStatLabelID, err := ulid.Parse(sportStatLabelIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingSportStatLabel, err := serviceSportStatLables.FindOne(ctx, sportStatLabelID)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "SportStatLabel not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		var sportStatLabelInput entity.SportStatLabels
		err = c.BodyParser(&sportStatLabelInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := validate.Struct(sportStatLabelInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error validate",
				"error":  err.Error(),
			})
		}

		if sportStatLabelInput.SportID != "" {
			sportId := sportStatLabelInput.SportID
			sport, err := serviceSport.FindOne(ctx, sportId)
			if err != nil {
				if ent.IsNotFound(err) {
					return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
						"status": "error",
						"error":  "SportStatLabel not found",
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}
			existingSportStatLabel.SportID = sport.ID
		}

		existingSportStatLabel.Label = sportStatLabelInput.Label
		existingSportStatLabel.Unit = sportStatLabelInput.Unit
		existingSportStatLabel.IsMain = sportStatLabelInput.IsMain

		err = serviceSportStatLables.Update(c.UserContext(), existingSportStatLabel)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error update",
				"error":  err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteSportStatLabel(ctx context.Context, serviceSportStatLables service.SportStatLabels) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sportStatLabelIDStr := c.Params("sportStatLabelId")
		sportStatLabelID, err := ulid.Parse(sportStatLabelIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = serviceSportStatLables.Delete(ctx, sportStatLabelID)
		if ent.IsNotFound(err) {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  "SportStatLabel not found",
			})
		}
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error delete",
				"error":  err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
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
				toJ[i].Sport = &presenter.Sport{
					ID:   condition.ID,
					Name: condition.Name,
				}
			}
		}
		return c.JSON(toJ)
	}
}

func addUserStat(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceEvent service.Event, serviceUser service.User, serviceNotification service.NotificationService, rdb *redis.Client) fiber.Handler {
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
			UserID ulid.ID `json:"user_id" validate:"required"`
			Stats  []struct {
				StatID    ulid.ID `json:"stat_id" validate:"required"`
				StatValue int     `json:"stat_value" validate:"gte=0"`
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
				"error":  err.Error(),
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

		fcmToken, err := serviceNotification.GetTokenFromRedis(ctx, rdb, string(user.ID)+"_FCM")
		if err == nil {
			err = serviceNotification.SendPushNotification(
				fcmToken,
				"Performance",
				"Nouvelles performances ajoutées",
			)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error send notification",
					"error":  err.Error(),
				})
			}
		}
		return c.SendStatus(fiber.StatusOK)

	}

}

func updateUserStats(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceNotification service.NotificationService, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userStatInput struct {
			Stats []struct {
				UserStatID ulid.ID `json:"user_stat_id" validate:"required"`
				StatValue  int     `json:"stat_value" validate:"gte=0"`
			} `json:"stats" validate:"required,dive"`
		}

		err := c.BodyParser(&userStatInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := validate.Struct(userStatInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error validate",
				"error":  err.Error(),
			})
		}

		err = serviceSportStatLables.UpdateUserStat(ctx, userStatInput.Stats)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error update user stat",
				"error":  err.Error(),
			})
		}

		user, err := serviceSportStatLables.GetUserByUserStatID(ctx, userStatInput.Stats[0].UserStatID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error get user",
				"error":  err.Error(),
			})
		}

		fcmToken, err := serviceNotification.GetTokenFromRedis(ctx, rdb, string(user.ID)+"_FCM")
		if err == nil {
			err = serviceNotification.SendPushNotification(
				fcmToken,
				"Performance",
				"Vos performances ont été mises à jour",
			)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error send notification",
					"error":  err.Error(),
				})
			}
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func getUserStatsByEvent(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceEvent service.Event, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {

		eventIDStr := c.Params("eventId")
		eventID, err := ulid.Parse(eventIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		userIDStr := c.Params("userId")
		userID, err := ulid.Parse(userIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		event, err := serviceEvent.FindOne(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find event",
				"error":  err.Error(),
			})
		}

		user, err := serviceUser.FindOne(ctx, userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find user",
				"error":  err.Error(),
			})
		}

		userStats, err := serviceSportStatLables.GetUserStatsByEventID(ctx, user.ID, event.ID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error get user stats",
				"error":  err.Error(),
			})
		}

		toJ := make([]presenter.UserStats, len(userStats))
		for i, userStat := range userStats {
			toJ[i] = presenter.UserStats{
				ID: userStat.ID,
				StatLabel: &presenter.SportStatLabels{
					ID:     userStat.Edges.Stat.ID,
					Label:  userStat.Edges.Stat.Label,
					Unit:   userStat.Edges.Stat.Unit,
					IsMain: userStat.Edges.Stat.IsMain,
				},
				Value: userStat.StatValue,
			}
		}
		return c.JSON(toJ)
	}
}

func getAllTeamUserMainStatsByEvent(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceEvent service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventIDStr := c.Params("eventId")
		eventID, err := ulid.Parse(eventIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error parse event ID",
				"error":  err.Error(),
			})
		}


		event, err := serviceEvent.FindOne(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find event",
				"error":  err.Error(),
			})
		}

		userStats, err := serviceSportStatLables.GetAllTeamUserMainStatsByEventID(ctx, event.ID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error get user stats",
				"error":  err.Error(),
			})
		}

		toJ := make([]presenter.UserStats, len(userStats))
		for i, userStat := range userStats {
			toJ[i] = presenter.UserStats{
				User: &presenter.User{
					ID: userStat.Edges.User.ID,
				},
				StatLabel: &presenter.SportStatLabels{
					Label:  userStat.Edges.Stat.Label,
				},
				Value: userStat.StatValue,
			}
		}		
		return c.JSON(toJ)
	}
}

func findMainStatLabelBySportID(ctx context.Context, serviceSportStatLables service.SportStatLabels) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sportIDStr := c.Params("sportId")
		sportID, err := ulid.Parse(sportIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		sportStatLabels, err := serviceSportStatLables.FindMainStatLabelBySportID(ctx, sportID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error find all",
				"error":  err.Error(),
			})
		}
		toJ := make([]presenter.SportStatLabels, len(sportStatLabels))
		for i, sportStatLabel := range sportStatLabels {
			toJ[i] = presenter.SportStatLabels{
				Label:  sportStatLabel.Label,
			}
		}
		return c.JSON(toJ)
	}
}

func getUserPerformanceBySport(ctx context.Context, serviceSportStatLables service.SportStatLabels, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {

		sportIDStr := c.Params("sportId")
		sportID, err := ulid.Parse(sportIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		userIDStr := c.Params("userId")
		userID, err := ulid.Parse(userIDStr)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := serviceUser.FindOne(ctx, userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find user",
				"error":  err.Error(),
			})
		}

		userStats, err := serviceSportStatLables.GetUserStatsBySportId(ctx, user.ID, sportID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error get user stats",
				"error":  err.Error(),
			})
		}

		aggregatedStats := make(map[string]int) // key: label, value: sum of values
		eventIDs := make(map[ulid.ID]bool)      // to track unique events

		for _, userStat := range userStats {
			if userStat.Edges.Stat != nil {
				// Aggregate stats by label
				label := userStat.Edges.Stat.Label
				aggregatedStats[label] += userStat.StatValue
			}

			if userStat.Edges.Event != nil {
				// Track unique events (using a map to ensure uniqueness)
				eventID := userStat.Edges.Event.ID // Use the actual `event_id` field

				eventIDs[eventID] = true
			}
		}

		toj := presenter.UserPerformance{
			NbEvents: len(eventIDs),
		}

		for label, value := range aggregatedStats {
			toj.Stats = append(toj.Stats, presenter.UserStats{
				StatLabel: &presenter.SportStatLabels{
					Label: label,
				},
				Value: value,
			})
		}

		c.Set("Cache-Control", "public, max-age=3600")
		return c.JSON(toj)
	}
}
