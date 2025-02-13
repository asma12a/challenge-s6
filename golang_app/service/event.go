package service

import (
	"context"
	"log"
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/sport"
	"github.com/asma12a/challenge-s6/ent/team"
	"github.com/asma12a/challenge-s6/ent/teamuser"
	"github.com/asma12a/challenge-s6/ent/user"
	"github.com/asma12a/challenge-s6/entity"
)

type Event struct {
	db *ent.Client
}

func NewEventService(client *ent.Client) *Event {
	return &Event{
		db: client,
	}
}

// @Summary Create an Event
// @Description Create a new event with its teams and players
// @Tags events
// @Accept  json
// @Produce  json
// @Param event body entity.Event true "Event to be created"
// @Param teams body []struct {
//     Name       string               `json:"name"`
//     MaxPlayers int                  `json:"max_players"`
//     Players    []struct {
//         Email string `json:"email"`
//         Role  string `json:"role,omitempty"`
//     } `json:"players"`
// } true "List of teams with players"
// @Success 201 {object} entity.Event "Event created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /events [post]

func (repo *Event) Create(ctx context.Context, event *entity.Event) error {

	tx, err := repo.db.Tx(ctx)
	if err != nil {
		log.Println(err, "error creating transaction")
		return err
	}

	newEvent := tx.Event.Create().
		SetName(event.Name).
		SetAddress(event.Address).
		SetLatitude(event.Latitude).
		SetLongitude(event.Longitude).
		SetEventCode(event.EventCode).
		SetDate(event.Date).
		SetSportID(event.SportID)

	if event.EventType != nil {
		newEvent.SetEventType(*event.EventType)
	}

	createdEvent, err := newEvent.Save(ctx)
	if err != nil {
		log.Println(err, "error creating event")
		_ = tx.Rollback()
		return err
	}

	_, err = tx.Team.Create().
		SetName("Equipe").
		SetEventID(createdEvent.ID).
		Save(ctx)
	if err != nil {
		log.Println(err, "error creating team")
		_ = tx.Rollback
	}

	if err := tx.Commit(); err != nil {
		log.Println("Erreur lors de la validation de la transaction :", err)
		return err
	}

	return nil
}

func (e *Event) FindOneWithDetails(ctx context.Context, id ulid.ID) (*entity.Event, error) {
	event, err := e.db.Event.Query().Where(event.IDEQ(id)).
		WithTeams(func(tq *ent.TeamQuery) {
			tq.Order(ent.Asc(team.FieldCreatedAt)).
				WithTeamUsers(func(tuq *ent.TeamUserQuery) {
					tuq.WithUser()
				})
		}).
		WithSport().
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Event{Event: *event}, nil
}

// @Summary Get an Event by ID
// @Description Get an event by its ID with details of teams and players
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Success 200 {object} entity.Event "Event details"
// @Failure 400 {object} map[string]interface{} "Bad Request"  // Remplacer fiber.Map par map[string]interface{}
// @Failure 404 {object} map[string]interface{} "Event Not Found"  // Remplacer fiber.Map par map[string]interface{}
// @Router /events/{id} [get]
func (e *Event) FindOne(ctx context.Context, id ulid.ID) (*entity.Event, error) {
	event, err := e.db.Event.Query().Where(event.IDEQ(id)).WithSport().Only(ctx)
	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Event{Event: *event}, nil
}

func (e *Event) FindEventByCode(ctx context.Context, code string) (*entity.Event, error) {
	event, err := e.db.Event.Query().
		Where(event.EventCode(code)).
		Where(event.DateGTE(time.Now())).
		WithSport().Only(ctx)
	if ent.IsNotFound(err) {
		return nil, err
	}
	return &entity.Event{Event: *event}, nil
}

// @Summary Update an Event
// @Description Update an existing event by ID
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Param event body entity.Event true "Updated event data"
// @Success 200 {object} entity.Event "Updated event"
// @Failure 400 {object} map[string]interface{} "Bad Request"  // Remplacer fiber.Map par map[string]interface{}
// @Failure 404 {object} map[string]interface{} "Event Not Found"  // Remplacer fiber.Map par map[string]interface{}
// @Router /events/{id} [put]
func (repo *Event) Update(ctx context.Context, event *entity.Event) (*entity.Event, error) {

	// Prepare the update query
	updatedEvent := repo.db.Event.
		UpdateOneID(event.ID).
		SetName(event.Name).
		SetAddress(event.Address).
		SetLatitude(event.Latitude).
		SetLongitude(event.Longitude).
		SetIsPublic(event.IsPublic).
		SetSportID(event.SportID).
		SetDate(event.Date)

	if event.EventType != nil {
		updatedEvent.SetEventType(*event.EventType)
	}

	e, err := updatedEvent.Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Event{Event: *e}, nil
}

// @Summary Delete an Event
// @Description Delete an event by ID
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path string true "Event ID"
// @Success 200 {object} map[string]interface{} "Event deleted"  // Remplacer fiber.Map par map[string]interface{}
// @Failure 404 {object} map[string]interface{} "Event Not Found"  // Remplacer fiber.Map par map[string]interface{}
// @Router /events/{id} [delete]
func (e *Event) Delete(ctx context.Context, id ulid.ID) error {
	err := e.db.Event.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// @Summary Get all Events
// @Description Get a list of all events
// @Tags events
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Event "List of events"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"  // Remplacer fiber.Map par map[string]interface{}
// @Router /events [get]

func (e *Event) List(ctx context.Context) ([]*ent.Event, error) {
	return e.db.Event.Query().WithSport().All(ctx)
}

// @Summary Search Events
// @Description Search for public events based on criteria such as name, address, type, and sport ID
// @Tags events
// @Accept  json
// @Produce  json
// @Param search query string false "Search term"
// @Param eventType query string false "Event type"
// @Param sportID query string false "Sport ID"
// @Success 200 {array} entity.Event "List of events matching search"
// @Failure 400 {object} map[string]interface{} "Bad Request"  // Remplacer fiber.Map par map[string]interface{}
// @Router /events/search [get]
func (e *Event) Search(ctx context.Context, search, eventType string, sportID *ulid.ID, userId ulid.ID) ([]*ent.Event, error) {
	userEvents, err := e.ListUserEvents(ctx, userId)

	if err != nil {
		return nil, err
	}

	var userEventIDs []ulid.ID
	for _, ue := range userEvents {
		userEventIDs = append(userEventIDs, ue.ID)
	}

	query := e.db.Event.Query().
		Where(event.IsPublicEQ(true)).
		Where(event.IDNotIn(userEventIDs...)).
		Where(event.DateGTE(time.Now()))

	if search != "" {
		query.Where(
			event.Or(
				event.NameContainsFold(search),
				event.AddressContainsFold(search),
			),
		)
	}
	if eventType != "" {
		query.Where(event.EventTypeEQ(event.EventType(eventType)))
	}

	if sportID != nil {
		query = query.Where(event.HasSportWith(sport.IDEQ(*sportID)))
	}

	return query.WithSport().All(ctx)
}

// ListUserEvents: Get all events for a user by getting all teams that the user is part of and then getting all events for those teams
func (e *Event) ListUserEvents(ctx context.Context, userID ulid.ID) ([]*ent.Event, error) {
	// teamuser has a user edge and a team edge
	teams, err := e.db.TeamUser.Query().Where(teamuser.HasUserWith(user.IDEQ(userID))).WithTeam().All(ctx)

	if err != nil {
		return nil, err
	}

	var teamIDs []ulid.ID
	for _, team := range teams {
		teamIDs = append(teamIDs, team.Edges.Team.ID)
	}

	events, err := e.db.Event.Query().
		Where(
			event.Or(
				event.HasTeamsWith(team.IDIn(teamIDs...)),
				event.CreatedByEQ(userID),
			),
		).
		Order(ent.Asc(event.FieldDate)).
		WithSport().All(ctx)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (e *Event) ListRecommendedEvents(ctx context.Context, lat, long float64, userID ulid.ID) ([]*ent.Event, error) {
	userEvents, err := e.ListUserEvents(ctx, userID)
	if err != nil {
		return nil, err
	}

	var userEventIDs []ulid.ID
	for _, ue := range userEvents {
		userEventIDs = append(userEventIDs, ue.ID)
	}

	events, err := e.db.Event.Query().
		Where(event.IsPublicEQ(true)).
		Where(event.IDNotIn(userEventIDs...)).
		Where(event.DateGTE(time.Now())).
		WithSport().All(ctx)
	if err != nil {
		return nil, err
	}

	events = entity.SortEventsByDistance(events, lat, long)

	if len(events) > 5 {
		events = events[:5]
	}

	return events, nil
}

func (repo *Event) FindAllTeamUsers(ctx context.Context, eventID ulid.ID) ([]*ent.TeamUser, error) {
	return repo.db.TeamUser.Query().Where(teamuser.HasTeamWith(team.HasEventWith(event.IDEQ(eventID)))).WithUser().All(ctx)
}

func (repo *Event) IsUserInEvent(ctx context.Context, eventID, userID ulid.ID) (bool, error) {
	teamUsers, err := repo.db.TeamUser.Query().
		Where(
			teamuser.HasTeamWith(team.HasEventWith(event.IDEQ(eventID))),
			teamuser.HasUserWith(user.IDEQ(userID)),
		).
		All(ctx)
	if err != nil {
		return false, err
	}
	return len(teamUsers) > 0, nil
}

func (n *Event) GetPlayersBeforeEvent(ctx context.Context) ([]*ent.Event, error) {
	now := time.Now().Truncate(time.Hour)
	exactlyIn24Hours := now.Add(24 * time.Hour) // Dans 24 heures, toujours à la minute près
	

	events, err := n.db.Event.Query().
		Where(
			event.DateGTE(exactlyIn24Hours),           
			event.DateLT(exactlyIn24Hours.Add(time.Hour)), 
		).
		WithTeams(func(tq *ent.TeamQuery) {
			tq.WithTeamUsers(func(tuq *ent.TeamUserQuery) {
				tuq.WithUser()
			})
		}).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return events, nil

}
