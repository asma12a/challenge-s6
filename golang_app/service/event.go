package service

import (
	"context"
	"log"
	"github.com/asma12a/challenge-s6/ent/user"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/teamuser"
	"github.com/asma12a/challenge-s6/ent/sport"
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

func (repo *Event) Create(ctx context.Context, event *entity.Event, teamsInput []struct {
	entity.Team
	Players    []struct {
		Email string `json:"email"`
		Role  string `json:"role,omitempty"`
	} `json:"players"`
}) error {

	tx, err := repo.db.Tx(ctx)
	if err != nil {
		log.Println(err, "error creating transaction")
		return err
	}

	newEvent := tx.Event.Create().
		SetName(event.Name).
		SetAddress(event.Address).
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

	teamNames := make(map[string]bool)

	for _, team := range teamsInput {
		if _, ok := teamNames[team.Name]; ok {
			log.Println("error creating team")
			_ = tx.Rollback()
			return entity.ErrCannotBeCreated
		}
		teamNames[team.Name] = true

		createdTeam, err := tx.Team.Create().
			SetName(team.Name).
			SetMaxPlayers(team.MaxPlayers).
			Save(ctx)
		if err != nil {
			log.Println(err, "error creating team")
			_ = tx.Rollback()
			return err
		}
		_, err = tx.EventTeams.Create().
			SetEventID(createdEvent.ID).
			SetTeamID(createdTeam.ID).
			Save(ctx)
		if err != nil {
			log.Println(err, "error creating event team")
			_ = tx.Rollback()
			return err
		}

		for _, player := range team.Players {
			playerEmail := player.Email
			user, err := tx.User.Query().Where(user.EmailEQ(playerEmail)).Only(ctx)
			if err != nil {
				if ent.IsNotFound(err) {
					teamUserCreate := tx.TeamUser.Create().
						SetTeamID(createdTeam.ID).
						SetEmail(player.Email).
						SetStatus("pending")

					// Définir le rôle uniquement s'il est présent dans le payload
					if player.Role != "" {
						teamUserCreate = teamUserCreate.SetRole(teamuser.Role(player.Role))
					}
					_, err = teamUserCreate.Save(ctx)
											
					if err != nil {
						log.Println("error adding player to team with null userId:", err)
						_ = tx.Rollback()
						return err
					}
				}else{
					log.Println("error finding user with email:", err)
					_ = tx.Rollback()
					return err
				}
			}else{
				teamUserCreate := tx.TeamUser.Create().
					SetTeamID(createdTeam.ID).
					SetUserID(user.ID).
					SetStatus("valid")
					if player.Role != "" {
						teamUserCreate = teamUserCreate.SetRole(teamuser.Role(player.Role))
					}
				_, err = teamUserCreate.Save(ctx)
				if err != nil {
					log.Println("error adding player to team with userId:", err)
					_ = tx.Rollback()
					return err
				}
			}

		}
	}
	if err := tx.Commit(); err != nil {
		log.Println("Erreur lors de la validation de la transaction :", err)
		return err
	}

	return nil
}

func (e *Event) FindOne(ctx context.Context, id ulid.ID) (*entity.Event, error) {
	event, err := e.db.Event.Query().Where(event.IDEQ(id)).Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Event{Event: *event}, nil
}

func (repo *Event) Update(ctx context.Context, event *entity.Event) (*entity.Event, error) {

	// Prepare the update query
	e, err := repo.db.Event.
		UpdateOneID(event.ID).
		SetName(event.Name).
		SetAddress(event.Address).
		SetEventCode(event.EventCode).
		SetDate(event.Date).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Event{Event: *e}, nil
}

func (e *Event) Delete(ctx context.Context, id ulid.ID) error {
	err := e.db.Event.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (e *Event) List(ctx context.Context) ([]*ent.Event, error) {
	return e.db.Event.Query().WithSport().All(ctx)
}

func (e *Event) Search(ctx context.Context, search, eventType string, sportID *ulid.ID) ([]*ent.Event, error) {
	query := e.db.Event.Query()
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
