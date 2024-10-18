package handler

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"time"

	"github.com/asma12a/challenge-s6/config/mailer"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func AuthHandler(app fiber.Router, ctx context.Context, serviceUser service.User, rdb *redis.Client) {
	app.Post("/signup", signUp(ctx, serviceUser, rdb))
}

func signUp(ctx context.Context, service service.User, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userInput *entity.User
		err := c.BodyParser(&userInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		newUser, err := entity.NewUser(
			userInput.Email,
			userInput.Name,
			userInput.Password,
		)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err,
			})
		}

		createdUser, err := service.Create(ctx, newUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		token := ulid.MustNew("")

		err = rdb.Set(ctx, "token:"+string(createdUser.ID), string(token), 30*time.Minute).Err()
		if err != nil {
			log.Println("Erreur lors du stockage de la clé :", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors du stockage du token",
			})
		}

		t, err := template.ParseFiles("template/signup_confirmation.html")
		if err != nil {
			log.Println("Erreur lors du chargement du template :", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors du chargement du template",
			})
		}

		data := struct {
			Name    string
			Message string
			Link    string
		}{
			Name:    createdUser.Name,
			Message: "Votre inscription a été réussie !",
			Link:    "http://localhost:3001/verify/" + string(token),
		}

		var body string
		buf := new(bytes.Buffer)
		err = t.Execute(buf, data)
		if err != nil {
			log.Println("Erreur lors de l'exécution du template :", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'exécution du template",
			})
		}
		body = buf.String()

		if err := mailer.SendEmail(createdUser.Email, "Confirmation d'inscription", body); err != nil {
			log.Println("Erreur lors de l'envoi de l'email :", err)
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'envoi de l'email de confirmation",
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}
