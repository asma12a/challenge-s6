package middleware

import (
	"strings"

	"github.com/asma12a/challenge-s6/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type TokenCheck struct {
	UserID string
	Roles  string
}

func CheckToken(c *fiber.Ctx) (TokenCheck, error) {
	// Récupérer le token depuis l'en-tête Authorization
	token := c.Get("Authorization")
	if token == "" {
		// Token manquant, renvoyer une erreur avec un code 401
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "No token provided")
	}

	// Enlever le préfixe "Bearer " si présent
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	} else {
		// Format invalide du token, renvoyer une erreur
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
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
		// Si le token est invalide, renvoyer une erreur
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Extraire les informations de l'utilisateur
	user_id, ok_id := claims["id"].(string)
	if !ok_id {
		// Si les informations du token sont manquantes
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
	}

	return TokenCheck{UserID: user_id}, nil
}
