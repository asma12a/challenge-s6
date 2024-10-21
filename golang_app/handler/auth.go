package handler

import (
	"bytes"
	"context"
	"html/template"
	"os"
	"time"

	"github.com/asma12a/challenge-s6/config/mailer"
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

// Définition de MyCustomClaims pour les claims personnalisés du JWT
type MyCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func AuthHandler(app fiber.Router, ctx context.Context, serviceUser service.User, rdb *redis.Client) {
	app.Post("/signup", signUp(ctx, serviceUser, rdb))
	app.Post("/login", login(ctx, serviceUser))

}

func signUp(ctx context.Context, serviceUser service.User, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userInput *entity.User
		err := c.BodyParser(&userInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
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
				"error":  err.Error(),
			})
		}

		createdUser, err := serviceUser.Create(ctx, newUser)
		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		token := ulid.MustNew("")

		err = rdb.Set(ctx, "token:"+string(createdUser.ID), string(token), 30*time.Minute).Err()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors du stockage du token",
			})
		}

		t, err := template.ParseFiles("template/signup_confirmation.html")
		if err != nil {
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
			Link: "http://localhost:3001/verify/" + string(token),
		}

		var body string
		buf := new(bytes.Buffer)
		err = t.Execute(buf, data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'exécution du template",
			})
		}
		body = buf.String()

		if err := mailer.SendEmail(createdUser.Email, "Confirmation d'inscription", body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'envoi de l'email de confirmation",
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func login(ctx context.Context, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Parse le corps de la requête
		var loginInput presenter.Login
		err := c.BodyParser(&loginInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "error",
				"message": "Invalid input",
				"error":   err.Error(),
			})
		}

		// Récupération des informations d'email et de mot de passe
		email := loginInput.Email
		password := loginInput.Password

		// Rechercher l'utilisateur par email
		user, err := serviceUser.FindByEmail(ctx, email)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status":       "error",
					"error_detail": "User not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":       "error",
				"error_detail": err,
				"error":        err.Error(),
			})
		}

		// Validation du mot de passe
		err = entity.ValidatePassword(user, password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status":  "error",
				"message": "Invalid password",
			})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": user.Name,
			"iss":  "squadgo",
			"iat":  time.Now().Unix(),
			"exp":  time.Now().Add(30 * 24 * time.Hour).Unix(),
			"nbf":  time.Now().Unix(),
		})

		s, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"token":  s,
		})
	}
}
