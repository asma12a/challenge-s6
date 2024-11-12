package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/asma12a/challenge-s6/config"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func IsAdminOrSelfAuthMiddleware(c *fiber.Ctx) error {
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
		return []byte(config.Env.JWTSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	// Extraire les informations de l'utilisateur
	user_id, ok_id := claims["id"].(string)
	roles, ok_roles := claims["roles"].(string)
	expiration, ok_exp := claims["exp"].(float64)
	if !ok_id || !ok_exp || !ok_roles {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Vérifier si le token a expiré
	if int64(expiration) < time.Now().Unix() {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token has expired",
		})
	}

	// Récupérer l'userId depuis les paramètres
	paramUserID := c.Params("userId")
	if paramUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No userId parameter provided",
		})
	}

	// Vérifier si l'utilisateur est admin ou si l'userId des paramètres est le même que celui du token
	if !strings.Contains(roles, "admin") && paramUserID != user_id {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}

	// Mets à jour le contexte en utilisant le viewer
	ctx := viewer.NewUserContext(context.Background(), &viewer.User{ID: ulid.ID(user_id)})
	c.Locals("user", fiber.Map{"id": user_id})
	c.SetUserContext(ctx)

	return c.Next()
}
