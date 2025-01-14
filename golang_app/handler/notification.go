package handler

import (
	"context"

	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func NotificationHandler(app fiber.Router, ctx context.Context, serviceSNotification service.NotificationService) {
	app.Post("/", sendNotification(serviceSNotification))

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


