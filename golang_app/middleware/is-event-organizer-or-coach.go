package middleware

import (
	"context"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/service"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/gofiber/fiber/v2"
)

func IsEventOrganizerOrCoach(ctx context.Context, serviceEvent service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := CheckToken(c)
		if err != nil {
			return err
		}

		// check if the token.UserID is the same as the event (from eventId in the URL) createdBy
		eventID, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}

		event, err := serviceEvent.FindOne(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Internal server error",
			})
		}

		teamUsers, err := serviceEvent.FindAllTeamUsers(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  "Internal server error",
			})
		}

		isCoach := false
		for _, teamUser := range teamUsers {
			if teamUser.Edges.User != nil && teamUser.Edges.User.ID == ulid.ID(token.UserID) && (teamUser.Role == "coach" || teamUser.Role == "org") {
				isCoach = true
				break
			}
		}

		isOrganizer := event.CreatedBy == ulid.ID(token.UserID)
		if !isCoach && !isOrganizer {
			return c.Status(fiber.StatusForbidden).JSON(&fiber.Map{
				"status": "error",
				"error":  "Access denied",
			})
		}

		// Mets à jour le contexte user en utilisant le viewer
		ctx := viewer.NewUserContext(context.Background(), &viewer.User{ID: ulid.ID(token.UserID)})
		c.SetUserContext(ctx)

		return c.Next()
	}
}
