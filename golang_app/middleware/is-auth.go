package middleware

import (
	"context"
	"fmt"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
)

func IsAuthMiddleware(c *fiber.Ctx) error {
	token, err := CheckToken(c)
	fmt.Println(token, err)
	if err != nil {
		return err
	}

	// Mets à jour le contexte user en utilisant le viewer
	ctx := viewer.NewUserContext(context.Background(), &viewer.User{ID: ulid.ID(token.UserID)})
	c.SetUserContext(ctx)

	return c.Next()
}
