package middleware

import (
	"strings"
	"time"

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
		// Si le token est invalide, re	nvoyer une erreur
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	// Extraire les informations de l'utilisateur
	user_id, ok_id := claims["id"].(string)
	roles, ok_roles := claims["roles"].(string)
	expiration, ok_exp := claims["exp"].(float64)
	if !ok_id || !ok_exp || !ok_roles {
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")

	}

	if int64(expiration) < time.Now().Unix() {
		return TokenCheck{}, fiber.NewError(fiber.StatusUnauthorized, "Invalid has expired")

	}

	return TokenCheck{UserID: user_id, Roles: roles}, nil
}
