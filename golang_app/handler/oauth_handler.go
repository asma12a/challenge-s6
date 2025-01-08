package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/oauth"
	"github.com/asma12a/challenge-s6/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
)

// Clé secrète pour signer les JWT (à stocker de manière sécurisée)
var jwtSecret = []byte("secret")

// Structure pour la génération du token JWT
type JWTClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

// Fonction pour générer un token aléatoire pour `state`
func generateState() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Printf("Erreur lors de la génération du state: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

type OAuthHandler struct {
	UserService *service.User
}

func NewOAuthHandler(userService *service.User) *OAuthHandler {
	return &OAuthHandler{UserService: userService}
}

func (h *OAuthHandler) OAuthLoginHandler(c *fiber.Ctx) error {
	state := generateState()
	c.Cookie(&fiber.Cookie{
		Name:     "state",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
		Secure:   true, // Assure-toi que ton serveur est en HTTPS
	})
	// Génère l'URL de redirection vers Google
	url := oauth.GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func (h *OAuthHandler) OAuthCallbackHandler(c *fiber.Ctx) error {
	// Récupérer le code de l'URL de redirection
	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code not found"})
	}

	// Échanger le code contre un jeton d'accès
	token, err := oauth.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Erreur d'échange de token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token exchange failed"})
	}

	// Obtenir les informations utilisateur
	client := oauth.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Erreur de requête utilisateur: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		log.Printf("Erreur de décodage des données utilisateur: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to decode user info"})
	}

	email, _ := userInfo["email"].(string)
	name, _ := userInfo["name"].(string)

	// Vérifier si l'utilisateur existe dans la base de données
	existingUser, err := h.UserService.FindByEmail(c.Context(), email)
	if err != nil && err != entity.ErrNotFound {
		log.Printf("Erreur lors de la vérification de l'utilisateur: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error checking user"})
	}

	if existingUser == nil {
		newUser := &entity.User{
			User: ent.User{
				Email:    email,
				Name:     name,
				Roles:    []string{"user"},
				Password: "defaultpass", // Ou générer un mot de passe temporaire
			},
		}

		createdUser, err := h.UserService.Create(c.Context(), newUser)
		if err != nil {
			log.Printf("Erreur lors de la création de l'utilisateur: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user"})
		}
		log.Printf("Nouvel utilisateur créé: %v", createdUser)
	} else {
		log.Printf("Utilisateur existant connecté: %v", existingUser)
	}

	// Générer un JWT pour l'utilisateur
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &JWTClaims{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "your-app-name", // À adapter selon ton application
		},
	}

	// Créer un token JWT signé
	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenJWT.SignedString(jwtSecret)
	if err != nil {
		log.Printf("Erreur lors de la création du JWT: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating JWT"})
	}

	// Retourner le JWT au frontend
	// Renvoie le token et les infos utilisateur
	return c.JSON(fiber.Map{
		"access_token": token.AccessToken,
		"user_info": fiber.Map{
			"email": email,
			"name":  userInfo["family_name"],
		},
	})
}
