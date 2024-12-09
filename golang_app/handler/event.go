package handler

import (
	"context"
	"log"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/asma12a/challenge-s6/middleware"
	"github.com/asma12a/challenge-s6/presenter"
	"github.com/asma12a/challenge-s6/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func EventHandler(app fiber.Router, ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) {
	app.Get("/", middleware.IsAdminMiddleware, listEvents(ctx, serviceEvent))
	app.Get("/search", searchEvent(ctx, serviceEvent))
	app.Get("/:eventId", getEvent(ctx, serviceEvent))
	app.Post("/", createEvent(ctx, serviceEvent, serviceSport))
	app.Post("/:eventId/team", addTeam(ctx, serviceEvent))
	//app.Post("/:eventId/team/:teamId/player", addPlayer(ctx, serviceEvent))
	app.Put("/:eventId", updateEvent(ctx, serviceEvent, serviceSport))
	app.Delete("/:eventId", deleteEvent(ctx, serviceEvent))

	// Handle event teams routes
}



var validate = validator.New()

func createEvent(ctx context.Context, serviceEvent service.Event, serviceSport service.Sport) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var eventInput struct {
			entity.Event // Inclut tous les champs de l'entité Event
			Teams        []struct {
				entity.Team
				Players []struct {
					Email string `json:"email"`
					Role  string `json:"role,omitempty"`
				} `json:"players"`
			} `json:"teams"`
		}

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

		log.Println(eventInput.SportID, "HEEEEE")

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
			eventInput.Date,
			sport.ID,
			eventInput.EventType,
		)

		err = serviceEvent.Create(c.UserContext(), newEvent, eventInput.Teams)
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
			EventCode:  event.EventCode,
			Date:       event.Date,
			CreatedAt:  event.CreatedAt,
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

		
		for _, eventTeam := range event.Edges.EventTeams {
			if eventTeam.Edges.Team != nil {
				team := eventTeam.Edges.Team
				teamToj := presenter.Team{
					ID:   team.ID,
					Name: team.Name,
					MaxPlayers: team.MaxPlayers,
				}

				for _, teamUser := range team.Edges.TeamUsers {
					user := teamUser.Edges.User
					if user != nil {
						teamToj.Players = append(teamToj.Players, presenter.User{
							ID:    user.ID,
							Name:  user.Name,
							Email: user.Email,
						})
					}
				}
				toJ.Teams = append(toJ.Teams, teamToj)
			}
		}



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
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
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
				EventCode:  event.EventCode,
				Date:       event.Date,
				CreatedAt:  event.CreatedAt,
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

func addTeam(ctx context.Context, serviceEvent service.Event) fiber.Handler {
	return func(c *fiber.Ctx) error {
		eventIDStr := c.Params("eventId")
		eventID, err := ulid.Parse(eventIDStr)


		var teamsInput []entity.Team
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error",
				"error":  "Invalid event ID format",
			})
		}

		log.Println("team input", teamsInput)

		

		err = c.BodyParser(&teamsInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error bodyparser",
				"error":  err.Error(),
			})
		}

		event, err := serviceEvent.FindOne(ctx, eventID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
				"status": "error find event",
				"error":  err.Error(),
			})
		}

	

		err = serviceEvent.AddTeam(ctx, event.ID, teamsInput)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status": "error add team",
				"error":  err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusOK)
	}
	
}