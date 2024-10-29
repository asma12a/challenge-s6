package middleware

import (
	"context"
	"os"
	"strings"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsAuthMiddleware(c *fiber.Ctx) error {
	// Récupérer le token depuis l'en-tête Authorization
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}

	// Enlever le préfixe "Bearer " si présent
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token format",
		})
	}

	// Vérifier et parser le token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		// Vérifiez la méthode de signature
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Extraire les informations de l'utilisateur
	user_id, ok := claims["id"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Mets à jour le contexte en utilisant le viewer

	// TODO : CORRIGER POUR ACCEDER AU CONTEXTE DE L'UTILISATEUR
	ctx := viewer.NewUserContext(context.Background(), &viewer.User{ID: ulid.ID(user_id)})
	c.Locals("user", fiber.Map{"id": user_id})
	c.SetUserContext(ctx)

	return c.Next()
}
