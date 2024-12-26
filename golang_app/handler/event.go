package handler

import (
	"context"
	"strconv"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/asma12a/challenge-s6/viewer"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func EventHandler(app fiber.Router, ctx context.Context, serviceEvent service.Event, serviceSport service.Sport, serviceTeam service.Team, serviceTeamUser service.TeamUser) {
	// User scoped routes
	app.Get("/user", listUserEvents(ctx, serviceEvent))

	// Team scoped routes
	TeamHandler(app.Group("/:eventId/teams", func(c *fiber.Ctx) error {
		eventID, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}

		_, err = serviceEvent.FindOne(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find event",
				"error":  err.Error(),
			})
		}

		return c.Next()
	}), ctx, serviceEvent, serviceTeam, serviceTeamUser)

	// Global event routes
	app.Get("/search", searchEvent(ctx, serviceEvent))
	app.Get("/recommended", listRecommendedEvents(ctx, serviceEvent))
	app.Get("/:eventId", getEvent(ctx, serviceEvent))
	app.Post("/", createEvent(ctx, serviceEvent, serviceSport))
	app.Put("/:eventId", updateEvent(ctx, serviceEvent, serviceSport))
	app.Delete("/:eventId", middleware.IsEventOrganizer(ctx, serviceEvent), deleteEvent(ctx, serviceEvent))

	// Admin routes
	app.Get("/", middleware.IsAdminMiddleware, listEvents(ctx, serviceEvent))

}

var validate = validator.New()

func createEvent(ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var eventInput entity.Event // Inclut tous les champs de l'entité Event

		err := c.BodyParser(&eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error bodyparser",
				"error":  err.Error(),
			})
		}

		// Valide les champs du JSON
		if err := validate.Struct(eventInput); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
				"status": "error validate",
				"error":  err.Error(),
			})
		}

		// Vérifie si le sport existe à partir de l'ID fourni
		sport, err := serviceSport.FindOne(ctx, eventInput.SportID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find sport",
				"error":  err.Error(),
			})
		}

		newEvent := entity.NewEvent(
			eventInput.Name,
			eventInput.Address,
			eventInput.Latitude,
			eventInput.Longitude,
			eventInput.Date,
			sport.ID,
			eventInput.EventType,
		)

		err = serviceEvent.Create(c.UserContext(), newEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error create event",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)
	}
}

func getEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := ulid.Parse(c.Params("eventId"))

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		event, err := service.FindOneWithDetails(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Mapper les données de event vers presenter.Event
		toJ := presenter.Event{
			ID:         event.ID,
			Name:       event.Name,
			Address:    event.Address,
			Latitude:   event.Latitude,
			Longitude:  event.Longitude,
			EventCode:  event.EventCode,
			Date:       event.Date,
			CreatedAt:  event.CreatedAt,
			CreatedBy:  event.CreatedBy,
			IsPublic:   event.IsPublic,
			IsFinished: event.IsFinished,
			EventType:  event.EventType,
		}

		if condition := event.Edges.Sport; condition != nil {
			toJ.Sport = presenter.Sport{
				ID:       condition.ID,
				Name:     condition.Name,
				ImageURL: condition.ImageURL,
			}
		}

		teamsToJ := make([]presenter.Team, 0, len(event.Edges.Teams))
		for _, eventTeam := range event.Edges.Teams {
			if eventTeam != nil {
				teamToj := presenter.Team{
					ID:         eventTeam.ID,
					Name:       eventTeam.Name,
					MaxPlayers: eventTeam.MaxPlayers,
					Players:    []presenter.Player{},
				}

				for _, teamUser := range eventTeam.Edges.TeamUsers {
					user := teamUser.Edges.User
					teamToj.Players = append(teamToj.Players, presenter.Player{
						ID: teamUser.ID,
						Name: func() string {
							if user != nil {
								return user.Name
							} else {
								return ""
							}
						}(),
						Email:  teamUser.Email,
						Role:   presenter.Role(teamUser.Role),
						Status: presenter.Status(teamUser.Status),
						UserID: func() ulid.ID {
							if user != nil {
								return user.ID
							} else {
								return ""
							}
						}(),
					})
				}
				teamsToJ = append(teamsToJ, teamToj)
			}
		}
		toJ.Teams = teamsToJ

		return c.JSON(toJ)
	}
}

func updateEvent(ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Récupère l'ID de l'événement à mettre à jour depuis les paramètres de la route
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		existingEvent, err := serviceEvent.FindOne(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		var eventInput entity.Event
		err = c.BodyParser(&eventInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		if eventInput.SportID != "" {
			sportId := eventInput.SportID
			sport, err := serviceSport.FindOne(ctx, sportId)
			if err != nil {
				if ent.IsNotFound(err) {
					return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
						"status": "error",
						"error":  "Sport not found",
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error",
					"error":  err.Error(),
				})
			}
			existingEvent.SportID = sport.ID
		}

		existingEvent.Name = eventInput.Name
		existingEvent.Address = eventInput.Address
		existingEvent.EventCode = eventInput.EventCode
		existingEvent.Date = eventInput.Date
		existingEvent.EventType = eventInput.EventType

		_, err = serviceEvent.Update(c.UserContext(), existingEvent)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

func deleteEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// id := c.Params("eventId")
		id, err := ulid.Parse(c.Params("eventId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		err = service.Delete(ctx, id)
		if err != nil {
			if ent.IsNotFound(err) {
				return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
					"status": "error",
					"error":  "Event not found",
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusNoContent)

	}
}

func listEvents(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		events, err := service.List(ctx)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.Event, len(events))

		for i, event := range events {
			toJ[i] = presenter.Event{
				ID:         event.ID,
				Name:       event.Name,
				Address:    event.Address,
				Latitude:   event.Latitude,
				Longitude:  event.Longitude,
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
				CreatedBy:  event.CreatedBy,
				IsPublic:   event.IsPublic,
				IsFinished: event.IsFinished,
				EventType:  event.EventType,
			}
			if condition := event.Edges.Sport; condition != nil {
				toJ[i].Sport = presenter.Sport{
					ID:       condition.ID,
					Name:     condition.Name,
					ImageURL: condition.ImageURL,
				}
			}
		}
		return c.JSON(toJ)
	}
}

func searchEvent(ctx context.Context, service service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {

		search := c.Query("search")
		eventType := c.Query("type")
		sportIDStr := c.Query("sport")

		var sportID *ulid.ID
		if sportIDStr != "" {
			parsedID, err := ulid.Parse(sportIDStr)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"status": "error",
					"error":  "Invalid sportID format",
				})
			}
			sportID = &parsedID
		}

		events, err := service.Search(ctx, search, eventType, sportID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.Event, len(events))

		for i, event := range events {
			toJ[i] = presenter.Event{
				ID:         event.ID,
				Name:       event.Name,
				Address:    event.Address,
				Latitude:   event.Latitude,
				Longitude:  event.Longitude,
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
				CreatedBy:  event.CreatedBy,
				IsPublic:   event.IsPublic,
				IsFinished: event.IsFinished,
				EventType:  event.EventType,
			}
			if condition := event.Edges.Sport; condition != nil {
				toJ[i].Sport = presenter.Sport{
					ID:       condition.ID,
					Name:     condition.Name,
					ImageURL: condition.ImageURL,
				}

			}
		}

		return c.JSON(toJ)
	}
}

func listUserEvents(ctx context.Context, serviceEvent service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		currentUser, err := viewer.UserFromContext(c.UserContext())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		events, err := serviceEvent.ListUserEvents(ctx, currentUser.ID)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.Event, len(events))

		for i, event := range events {
			toJ[i] = presenter.Event{
				ID:         event.ID,
				Name:       event.Name,
				Address:    event.Address,
				Latitude:   event.Latitude,
				Longitude:  event.Longitude,
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
				CreatedBy:  event.CreatedBy,
				IsPublic:   event.IsPublic,
				IsFinished: event.IsFinished,
				EventType:  event.EventType,
			}
			if condition := event.Edges.Sport; condition != nil {
				toJ[i].Sport = presenter.Sport{
					ID:       condition.ID,
					Name:     condition.Name,
					ImageURL: condition.ImageURL,
				}
			}
		}
		return c.JSON(toJ)
	}
}
func listRecommendedEvents(ctx context.Context, serviceEvent service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		defaultLatitude := 46.603354
		defaultLongitude := 1.888334
		latitudeStr := c.Query("latitude", strconv.FormatFloat(defaultLatitude, 'f', -1, 64))
		longitudeStr := c.Query("longitude", strconv.FormatFloat(defaultLongitude, 'f', -1, 64))

		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		if err != nil {
			latitude = defaultLatitude
		}

		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		if err != nil {
			longitude = defaultLongitude
		}

		events, err := serviceEvent.ListRecommendedEvents(ctx, latitude, longitude)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"error": err.Error(),
			})
		}

		toJ := make([]presenter.Event, len(events))

		for i, event := range events {
			toJ[i] = presenter.Event{
				ID:         event.ID,
				Name:       event.Name,
				Address:    event.Address,
				Latitude:   event.Latitude,
				Longitude:  event.Longitude,
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
				CreatedBy:  event.CreatedBy,
				IsPublic:   event.IsPublic,
				IsFinished: event.IsFinished,
				EventType:  event.EventType,
			}
			if condition := event.Edges.Sport; condition != nil {
				toJ[i].Sport = presenter.Sport{
					ID:       condition.ID,
					Name:     condition.Name,
					ImageURL: condition.ImageURL,
				}
			}
		}
		return c.JSON(toJ)
	}
}
