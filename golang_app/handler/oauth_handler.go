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
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
)

var jwtSecret = []byte("secret")

type JWTClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}

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

// Fonction pour créer un mot de passe temporaire sécurisé
func generateSecurePassword() (string, error) {
	password := make([]byte, 16)
	_, err := rand.Read(password)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(password), nil
}

func (h *OAuthHandler) OAuthLoginHandler(c *fiber.Ctx) error {
	state := generateState()
	c.Cookie(&fiber.Cookie{
		Name:     "state",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
	})

	url := oauth.GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
	log.Printf("url %v", url)

	return c.JSON(fiber.Map{
		"auth_url": url,
	})
}

func (h *OAuthHandler) OAuthCallbackHandler(c *fiber.Ctx) error {
	log.Printf("OAuthCallbackHandler début")

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Code not found"})
	}

	token, err := oauth.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Erreur d'échange de token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token exchange failed"})
	}

	userInfo, err := fetchGoogleUserInfo(token)
	if err != nil {
		log.Printf("Erreur de récupération des informations utilisateur: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}

	user, err := h.handleUserInDatabase(c, userInfo)
	if err != nil {
		log.Printf("Erreur lors de la gestion de l'utilisateur: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error handling user"})
	}

	jwtToken, err := generateJWT(user.Email, user.Name)
	if err != nil {
		log.Printf("Erreur lors de la génération du JWT: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating JWT"})
	}

	return c.JSON(fiber.Map{
		"jwt": jwtToken,
		"user_info": fiber.Map{
			"email": user.Email,
			"name":  user.Name,
		},
	})
}

func fetchGoogleUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := oauth.GoogleOAuthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (h *OAuthHandler) handleUserInDatabase(ctx *fiber.Ctx, userInfo map[string]interface{}) (*entity.User, error) {
	email, _ := userInfo["email"].(string)
	name, _ := userInfo["name"].(string)

	existingUser, err := h.UserService.FindByEmail(ctx.Context(), email)
	if err != nil && err != entity.ErrNotFound {
		return nil, err
	}

	if existingUser == nil {
		tempPassword, err := generateSecurePassword()
		if err != nil {
			log.Printf("Erreur lors de la génération du mot de passe: %v", err)
			return nil, err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tempPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Erreur lors du hachage du mot de passe: %v", err)
			return nil, err
		}

		newUser := &entity.User{
			User: ent.User{
				Email:    email,
				Name:     name,
				Roles:    []string{"user"},
				Password: string(hashedPassword),
			},
		}

		return h.UserService.Create(ctx.Context(), newUser)
	}

	return existingUser, nil
}

func generateJWT(email, name string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &JWTClaims{
		Email: email,
		Name:  name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "SQUAD-GO",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
