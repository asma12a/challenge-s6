package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"slices"

	"github.com/asma12a/challenge-s6/config/mailer"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/gofiber/fiber/v2"
	passwordValidator "github.com/wagslane/go-password-validator"
)

// Fonction pour générer un mot de passe aléatoire

func UserHandler(app fiber.Router, ctx context.Context, service service.User) {
	app.Get("/", middleware.IsAdminMiddleware, listUsers(ctx, service))
	app.Get("/:userId", middleware.IsAdminOrSelfAuthMiddleware, getUser(ctx, service))
	app.Post("/", middleware.IsAdminMiddleware, createUser(ctx, service))
	app.Put("/:userId/password", middleware.IsAdminOrSelfAuthMiddleware, updateUserPassword(ctx, service))
	app.Put("/:userId", middleware.IsAdminMiddleware, updateUserForAdmin(ctx, service))
	app.Put("/:userId/user", middleware.IsSelfAuthMiddleware, updateUser(ctx, service))
	app.Delete("/:userId", middleware.IsAdminMiddleware, deleteUser(ctx, service))
}

// Génère un mot de passe aléatoire de longueur spécifiée
func generateRandomPassword(length int) (string, error) {
	// Crée un tableau de bytes aléatoires
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func createUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userInput *entity.User
		err := c.BodyParser(&userInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Génère un mot de passe aléatoire
		randomPassword, err := generateRandomPassword(12)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de la génération du mot de passe",
			})
		}

		newUser, err := entity.NewUser(
			userInput.Email,
			userInput.Name,
			randomPassword,
		)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		_, err = service.Create(c.UserContext(), newUser)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		body := fmt.Sprintf(`
			<p>Bonjour %s,</p>
			<p>Votre compte a été créé avec succès. Voici vos informations de connexion :</p>
			<p>Email: %s</p>
			<p>Mot de passe: %s</p>
			<p>Veuillez vous connecter et changer votre mot de passe dès que possible pour plus de sécurité.</p>
			<p>Cordialement, l'équipe SquadGo.</p>
		`, newUser.Name, newUser.Email, randomPassword)

		// Envoi de l'email avec le mot de passe généré
		if err := mailer.SendEmail(newUser.Email, "Confirmation d'inscription", body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'envoi de l'email de confirmation",
			})
		}

		// Réponse réussie
		return c.SendStatus(fiber.StatusCreated)
	}
}

func getUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		data := presenter.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}

		return c.Status(fiber.StatusOK).JSON(data)
	}
}

func updateUserForAdmin(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrInvalidID.Error(),
			})
		}

		var userInput *entity.User

		if err := c.BodyParser(&userInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := validate.Struct(userInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		oldEmail := user.Email
		newEmail := userInput.Email
		emailSubject := "Mise à jour de votre compte"
		var emailBody string

		user.Name = userInput.Name
		user.Email = userInput.Email

			if userInput.Password != "" {
				if err := passwordValidator.Validate(userInput.Password, 60); err != nil {
					return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
						"status": "error",
						"error":  entity.ErrPasswordNotStrong.Error(),
					})
				}
	
				hashedPassword, err := user.GeneratePassword(userInput.Password)
				if err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
						"status": "error",
						"error":  err.Error(),
					})
				}
				user.Password = hashedPassword
			}
	
			if slices.Contains(user.Roles, "admin") && userInput.Roles != nil {
				user.Roles = userInput.Roles
			}

		updatedUser, err := service.Update(c.UserContext(), user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if oldEmail != newEmail {
			emailBody = fmt.Sprintf(`
				<p>Bonjour %s,</p>
				<p>Votre adresse email a été modifiée. Vous devez désormais utiliser l'adresse suivante pour vous connecter :</p>
				<p>Nouvelle adresse email : %s</p>
				<p>Si vous n'êtes pas à l'origine de cette modification, veuillez nous contacter immédiatement.</p>
				<p>Cordialement,</p>
				<p>L'équipe SquadGo</p>
			`, user.Name, newEmail)

			if err := mailer.SendEmail(oldEmail, emailSubject, emailBody); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"status": "error",
					"error":  "Erreur lors de l'envoi de l'email à l'ancienne adresse",
				})
			}
		} else {
			emailBody = fmt.Sprintf(`
				<p>Bonjour %s,</p>
				<p>Des changements ont été effectués sur votre compte :</p>
				<p>Nom : %s</p>
				<p>Rôles : %v</p>
				<p>Si vous n'êtes pas à l'origine de ces modifications, veuillez nous contacter immédiatement.</p>
				<p>Cordialement,</p>
				<p>L'équipe SquadGo</p>
			`, user.Name, user.Name, user.Roles)

			if err := mailer.SendEmail(newEmail, emailSubject, emailBody); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
					"status": "error",
					"error":  "Erreur lors de l'envoi de l'email à la nouvelle adresse",
				})
			}
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User updated and email sent",
			"data":    updatedUser,
		})
	}
}

func updateUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrInvalidID.Error(),
			})
		}

		var userInput *entity.User

		if err := c.BodyParser(&userInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := validate.Struct(userInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user.Name = userInput.Name
		user.Email = userInput.Email



		updatedUser, err := service.Update(c.UserContext(), user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User updated",
			"data":    updatedUser,
		})
	}
}



func updateUserPassword(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  entity.ErrInvalidID.Error(),
			})
		}

		var userInput struct {
			Password string `json:"password,omitempty"`
		}

		if err := c.BodyParser(&userInput); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if err := validate.Struct(userInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if userInput.Password != "" {
			if err := passwordValidator.Validate(userInput.Password, 60); err != nil {
				return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
					"status": "error",
					"error":  entity.ErrPasswordNotStrong.Error(),
				})
			}

			hashedPassword, err := user.GeneratePassword(userInput.Password)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}
			user.Password = hashedPassword
		}

		_, err = service.Update(c.UserContext(), user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteUser(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Récupérer l'utilisateur avant de le supprimer pour pouvoir lui envoyer un email
		user, err := service.FindOne(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Suppression de l'utilisateur
		if err := service.Delete(ctx, id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Préparer le contenu de l'email
		emailSubject := "Suppression de votre compte"
		emailBody := fmt.Sprintf(`
			<p>Bonjour %s,</p>
			<p>Nous vous informons que votre compte a été supprimé par un administrateur pour des raisons spécifiques.</p>
			<p>Si vous pensez qu'il s'agit d'une erreur, veuillez nous contacter à l'adresse support@squadgo.com.</p>
			<p>Cordialement,</p>
			<p>L'équipe SquadGo</p>
		`, user.Name)

		// Envoi de l'email de suppression du compte
		if err := mailer.SendEmail(user.Email, emailSubject, emailBody); err != nil {
			// Si l'envoi échoue, on retourne une erreur 500 mais on continue la suppression de l'utilisateur
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Erreur lors de l'envoi de l'email de suppression",
			})
		}

		// Réponse de succès après la suppression de l'utilisateur
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "User deleted and email sent",
		})
	}
}

func listUsers(ctx context.Context, service service.User) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := service.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		toJ := make([]presenter.User, len(users))

		for i, user := range users {
			toJ[i] = presenter.User{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Roles: user.Roles,
			}
		}

		return c.JSON(toJ)
	}
}
