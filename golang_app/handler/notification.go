package handler

import (
	"context"
	"log"
	"time"

	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/service"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func NotificationHandler(app fiber.Router, ctx context.Context, serviceNotification service.NotificationService, serviceEvent service.Event, rdb *redis.Client) {
	app.Post("/", sendNotification(serviceNotification))
	app.Get("/event/notify", notifyPlayerBeforeEvent(ctx, serviceNotification, serviceEvent, rdb))
	app.Post("/fcm_token/:fcm_token", middleware.IsAuthMiddleware, storeFcmToken(ctx, serviceNotification, rdb))

}

func sendNotification(serviceSNotification service.NotificationService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var notificationInput = struct {
			Token string `json:"token"`
			Title string `json:"title"`
			Body  string `json:"body"`
		}{}

		err := c.BodyParser(&notificationInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = serviceSNotification.SendPushNotification(notificationInput.Token, notificationInput.Title, notificationInput.Body)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func notifyPlayerBeforeEvent(ctx context.Context, serviceSNotification service.NotificationService, serviceEvent service.Event, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		events, err := serviceEvent.GetPlayersBeforeEvent(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		for _, event := range events {
			for _, team := range event.Edges.Teams {
				for _, teamUser := range team.Edges.TeamUsers {
					log.Println("Envoi de la notification à", teamUser.Edges.User.Email, event.Name)
					fcm_token, err := serviceSNotification.GetTokenFromRedis(ctx, rdb, string(teamUser.Edges.User.ID)+"_FCM")
					if err == nil {
						err = serviceSNotification.SendPushNotification(
							fcm_token,
							"Rappel",
							"L'événement "+event.Name+" commence bientôt",
						)
						if err != nil {
							return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
								"status": "error",
								"error":  err.Error(),
							})
						}
						time.Sleep(5 * time.Second)
					}

				}
			}
		}
		return c.SendStatus(fiber.StatusOK)
	}
}

func storeFcmToken(ctx context.Context, serviceNotification service.NotificationService, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser, err := viewer.UserFromContext(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		fcm_token := c.Params("fcm_token")
		if currentUser.ID != "" && fcm_token != "" {
			err = serviceNotification.StoreTokenInRedis(ctx, rdb, string(currentUser.ID)+"_FCM", fcm_token)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"status": "error store token in redis",
					"error":  err.Error(),
				})
			}
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
