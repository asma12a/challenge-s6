package handler

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"
	"strings"
	"time"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/config/mailer"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	passwordValidator "github.com/wagslane/go-password-validator"
)

// Définition de MyCustomClaims pour les claims personnalisés du JWT
type MyCustomClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func AuthHandler(app fiber.Router, ctx context.Context, serviceUser service.User, serviceTeamUser service.TeamUser, rdb *redis.Client) {
	app.Post("/signup", signUp(ctx, serviceUser, serviceTeamUser, rdb))
	app.Post("/login", login(ctx, serviceUser))
	app.Get("/me", middleware.IsAuthMiddleware, me(ctx, serviceUser))
	app.Get("/verify/:token", verify(ctx, serviceUser, rdb))
}

type SignUpRequestInput struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func signUp(ctx context.Context, serviceUser service.User, serviceTeamUser service.TeamUser, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userInput SignUpRequestInput
		err := c.BodyParser(&userInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrCannotParseJSON.Error(),
			})
		}

		if err := validate.Struct(userInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := passwordValidator.Validate(userInput.Password, 60); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrPasswordNotStrong.Error(),
			})
		}

		newUser, err := entity.NewUser(
			userInput.Email,
			userInput.Name,
			userInput.Password,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		createdUser, err := serviceUser.Create(c.UserContext(), newUser)
		if err != nil {
			return c.Status(fiber.StatusConflict).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
		go func() {
			err = serviceTeamUser.UpdateTeamUserWithUser(c.UserContext(), *createdUser)
			if err != nil {
				log.Println("Error updating team user:", err)
			}
		}()

		token := ulid.MustNew("")

		redisKey := fmt.Sprintf("token:%s", token)

		err = rdb.Set(ctx, redisKey, string(createdUser.ID), 30*time.Minute).Err()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		t, err := template.ParseFiles("template/signup_confirmation.html")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors du chargement du template",
			})
		}

		fmt.Println("Serveur: ", config.Env.ServerURL)

		link := fmt.Sprintf("%s/api/auth/verify/%s", config.Env.ServerURL, token)

		data := struct {
			Name    string
			Message string
			Link    string
		}{
			Link: link,
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

func verify(ctx context.Context, serviceUser service.User, rdb *redis.Client) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Récupérer le token de l'URL (en paramètre)
		token := c.Params("token")
		if token == "" {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Token manquant",
			})
		}

		// Chercher le token dans Redis
		storedToken, err := rdb.Get(ctx, "token:"+token).Result()
		if err != nil {
			if err == redis.Nil {
				t, err := template.ParseFiles("template/token_invalid.html")
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
						"status": "error",
						"error":  "Erreur lors du chargement du template",
					})
				}
				data := struct {
					Message string
				}{
					Message: "Le lien de vérification que vous avez utilisé est invalide ou a expiré. Veuillez demander un nouveau lien de vérification.",
				}

				var buf bytes.Buffer
				err = t.Execute(&buf, data)
				if err != nil {
					return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
						"status": "error",
						"error":  "Erreur lors de l'exécution du template.",
					})
				}
				c.Set("Content-Type", "text/html")
				return c.SendString(buf.String())
			}
			// Autre erreur Redis
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de la récupération du token",
			})
		}

		// Si le token est valide, récupérer l'utilisateur associé
		userID, err := ulid.Parse(storedToken)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Problème lors de la récupération de l'ID de l'utilisateur",
			})
		}

		// Récupérer l'utilisateur par ID
		updatedUser, err := serviceUser.FindOne(c.UserContext(), userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  "Utilisateur non trouvé",
			})
		}

		updatedUser.IsActive = true
		updatedUser.UpdatedAt = time.Now()

		_, err = serviceUser.Update(c.UserContext(), updatedUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de la mise à jour de l'utilisateur",
			})
		}

		// Supprimer le token de Redis après validation
		err = rdb.Del(ctx, "token:"+token).Err()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de la suppression du token",
			})
		}

		// Charger et remplir le template HTML
		t, err := template.ParseFiles("template/verification_success.html")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors du chargement du template",
			})
		}

		// Données à passer au template
		data := struct {
			Name    string
			Message string
		}{
			Name:    updatedUser.Name,
			Message: "Votre compte a été activé avec succès !",
		}

		// Exécution du template avec les données et envoi de la réponse HTML.
		var buf bytes.Buffer
		err = t.Execute(&buf, data)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'exécution du template.",
			})
		}

		// Définir le type de réponse comme HTML et envoyer le contenu.
		c.Set("Content-Type", "text/html")
		return c.SendString(buf.String())
	}
}

type LoginRequestInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func login(ctx context.Context, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {

		// Parse le corps de la requête
		var loginInput LoginRequestInput

		err := c.BodyParser(&loginInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrCannotParseJSON.Error(),
			})
		}

		// Valide les champs du JSON
		if err := validate.Struct(loginInput); err != nil {
			log.Printf("Validation error: %v", err)

			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Récupération des informations d'email et de mot de passe
		email := loginInput.Email
		password := loginInput.Password

		// Rechercher l'utilisateur par email
		user, err := serviceUser.FindByEmail(ctx, email)
		if err != nil {
			log.Printf("User not found: %v", err)

			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrEntityNotFound("User").Error(),
			})
		}

		// Validation du mot de passe
		err = entity.ValidatePassword(user, password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrInvalidPassword.Error(),
			})
		}

		if !user.IsActive {
			return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrUserNotActive.Error(),
			})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"roles": strings.Join(user.Roles, ","),
			"iss":   "squadgo",
			"iat":   time.Now().Unix(),
			"exp":   time.Now().Add(30 * 24 * time.Hour).Unix(),
			"nbf":   time.Now().Unix(),
		})

		s, err := token.SignedString([]byte(config.Env.JWTSecret))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
			"user": presenter.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Roles: user.Roles,
			},
			"token": s,
		})
	}
}

func me(ctx context.Context, serviceUser service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {

		currentUser, err := viewer.UserFromContext(c.UserContext())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := serviceUser.FindOne(ctx, currentUser.ID)

		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrEntityNotFound("User").Error(),
			})
		}

		data := presenter.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Roles: user.Roles,
		}

		return c.Status(fiber.StatusOK).JSON(data)
	}
}
