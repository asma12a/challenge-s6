package handler

import (
	"context"
	"time"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
)

func MessageHandler(app fiber.Router, ctx context.Context, serviceMessage service.MessageService, serviceEvent service.Event, serviceUser service.User) {
	app.Get("/", listMessages(ctx, serviceMessage))
	app.Get("/:messageId", getMessage(ctx, serviceMessage))
	app.Post("/", createMessage(ctx, serviceMessage, serviceEvent, serviceUser))
	app.Put("/:messageId", updateMessage(ctx, serviceMessage, serviceEvent, serviceUser))
	app.Delete("/:messageId", deleteMessage(ctx, serviceMessage))
}

// createMessage permet de créer un nouveau message
func createMessage(ctx context.Context, serviceMessage service.MessageService, serviceEvent service.Event, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var messageInput entity.Message

		err := c.BodyParser(&messageInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Vérification de l'existence de l'événement (Event)
		event, err := serviceEvent.FindOne(ctx, messageInput.EventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "Event not found",
			})
		}

		// Vérification de l'existence de l'utilisateur (User)
		user, err := serviceUser.FindOne(ctx, messageInput.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "User not found",
			})
		}

		newMessage := entity.NewMessage(
			event.ID,
			user.ID,
			messageInput.Content,
			time.Now(), // Timestamp pour la création du message
		)

		err = serviceMessage.Create(ctx, newMessage)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

// getMessage permet de récupérer un message par son ID
func getMessage(ctx context.Context, serviceMessage service.MessageService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("messageId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		message, err := serviceMessage.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "Message not found",
			})
		}

		// Mapper les données de message vers presenter.Message
		toJ := presenter.Message{
			ID:        message.ID,
			EventID:   message.EventID,
			UserID:    message.UserID,
			Content:   message.Content,
			CreatedAt: message.CreatedAt,
		}

		return c.JSON(toJ)
	}
}

// updateMessage permet de mettre à jour un message existant
func updateMessage(ctx context.Context, serviceMessage service.MessageService, serviceEvent service.Event, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("messageId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingMessage, err := serviceMessage.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "Message not found",
			})
		}

		var messageInput entity.Message
		err = c.BodyParser(&messageInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Si le champ event_id est fourni dans le body, on vérifie l'événement
		if messageInput.EventID != "" {
			event, err := serviceEvent.FindOne(ctx, messageInput.EventID)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "Event not found",
				})
			}
			existingMessage.EventID = event.ID
		}

		// Si le champ user_id est fourni dans le body, on vérifie l'utilisateur
		if messageInput.UserID != "" {
			user, err := serviceUser.FindOne(ctx, messageInput.UserID)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "User not found",
				})
			}
			existingMessage.UserID = user.ID
		}

		// Mettre à jour le contenu s'il est fourni
		if messageInput.Content != "" {
			existingMessage.Content = messageInput.Content
		}

		// Mettre à jour le message dans la base de données
		_, err = serviceMessage.Update(ctx, existingMessage)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

// deleteMessage permet de supprimer un message par son ID
func deleteMessage(ctx context.Context, serviceMessage service.MessageService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("messageId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = serviceMessage.Delete(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": "Message not found",
			})
		}

		return c.SendStatus(fiber.StatusNoContent)
	}
}

// listMessages permet de lister tous les messages
func listMessages(ctx context.Context, serviceMessage service.MessageService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		messages, err := serviceMessage.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.Message, len(messages))

		for i, message := range messages {
			toJ[i] = presenter.Message{
				ID:        message.ID,
				EventID:   message.EventID,
				UserID:    message.UserID,
				Content:   message.Content,
				CreatedAt: message.CreatedAt,
			}
		}

		return c.JSON(toJ)
	}
}
