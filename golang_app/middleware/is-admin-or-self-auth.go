package middleware

import (
	"context"
	"strings"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
)

func IsAdminOrSelfAuthMiddleware(c *fiber.Ctx) error {
	token, err := CheckToken(c)
	if err != nil {
		return err
	}

	// Récupérer l'userId depuis les paramètres
	paramUserID := c.Params("userId")
	if paramUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No userId parameter provided",
		})
	}

	// Vérifier si l'utilisateur est admin ou si l'userId des paramètres est le même que celui du token
	if !strings.Contains(token.Roles, "admin") && paramUserID != token.UserID {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	// Mets à jour le contexte user en utilisant le viewer
	ctx := viewer.NewUserContext(context.Background(), &viewer.User{ID: ulid.ID(token.UserID)})
	c.SetUserContext(ctx)

	return c.Next()
}
