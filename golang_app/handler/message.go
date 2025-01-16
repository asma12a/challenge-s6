package handler

import (
	"context"
	"time"

	"github.com/asma12a/challenge-s6/ent"
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
	app.Get("/event/:eventId", listMessagesByEvent(ctx, serviceMessage))
}

// createMessage permet de créer un nouveau message
func createMessage(ctx context.Context, serviceMessage service.MessageService, serviceEvent service.Event, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var messageInput entity.Message

		// Parser le corps de la requête
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
				"status": "error",
				"error":  "Event not found",
			})
		}

		user, err := serviceUser.FindOne(ctx, messageInput.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  "User not found",
			})
		}

		if user.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "User name is missing",
			})
		}

		newMessage := &entity.Message{
			EventID: event.ID,
			UserID:  user.ID,
			Message: ent.Message{
				UpdatedAt: time.Now(),
				UserName:  user.Name,
				Content:   messageInput.Content,
				CreatedAt: time.Now(),
			},
		}

		defer func() {
			if r := recover(); r != nil {
				c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"status": "error",
					"error":  "An internal server error occurred",
				})
			}
		}()

		err = serviceMessage.Create(c.UserContext(), newMessage)
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
				"status": "error",
				"error":  "Message not found",
			})
		}

		// Mapper les données de message vers presenter.Message
		toJ := presenter.Message{
			ID:      message.ID,
			EventID: message.EventID,
			User: presenter.User{
				ID:    message.Edges.User.ID,
				Name:  message.Edges.User.Name,
				Email: message.Edges.User.Email,
				Roles: message.Edges.User.Roles,
			},
			UserName:  message.UserName,
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
				"status": "error",
				"error":  "Message not found",
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
					"status": "error",
					"error":  "Event not found",
				})
			}
			existingMessage.EventID = event.ID
		}

		// Si le champ user_id est fourni dans le body, on vérifie l'utilisateur
		if messageInput.UserID != "" {
			user, err := serviceUser.FindOne(ctx, messageInput.UserID)
			if err != nil {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "User not found",
				})
			}
			existingMessage.UserID = user.ID
		}

		// Mettre à jour le contenu s'il est fourni
		if messageInput.Content != "" {
			existingMessage.Content = messageInput.Content
		}

		// Mettre à jour le message dans la base de données
		_, err = serviceMessage.Update(c.UserContext(), existingMessage)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
				"status": "error",
				"error":  "Message not found",
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

		return c.JSON(messages)
	}
}

// listMessagesByEvent permet de lister les messages associés à un événement
func listMessagesByEvent(ctx context.Context, serviceMessage service.MessageService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventID, err := ulid.Parse(c.Params("eventID"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid Event ID",
			})
		}

		messages, err := serviceMessage.ListByEvent(c.UserContext(), eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  "No messages found for this event",
			})
		}

		c.Set("Cache-Control", "public, max-age=3600")
		return c.JSON(messages)
	}
}
