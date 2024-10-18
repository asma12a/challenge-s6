package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/eventteams"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
)

type EventTeamsService struct {
	db *ent.Client
}

// NewEventTeamsService crée un nouveau service pour l'entité EventTeams
func NewEventTeamsService(client *ent.Client) *EventTeamsService {
	return &EventTeamsService{
		db: client,
	}
}

// Create permet de créer une nouvelle entrée dans EventTeams
func (repo *EventTeamsService) Create(ctx context.Context, eventTeams *entity.EventTeams) error {
	_, err := repo.db.EventTeams.Create().
		SetEventID(eventTeams.EventID).
		SetTeamID(eventTeams.TeamID).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

// FindOne permet de récupérer une entrée EventTeams par ID
func (repo *EventTeamsService) FindOne(ctx context.Context, id ulid.ID) (*entity.EventTeams, error) {
	eventTeams, err := repo.db.EventTeams.Query().
		Where(eventteams.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.EventTeams{EventTeams: *eventTeams}, nil
}

// Update permet de mettre à jour une entrée EventTeams
func (repo *EventTeamsService) Update(ctx context.Context, eventTeams *entity.EventTeams) (*entity.EventTeams, error) {
	e, err := repo.db.EventTeams.
		UpdateOneID(eventTeams.ID).
		SetEventID(eventTeams.EventID).
		SetTeamID(eventTeams.TeamID).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.EventTeams{EventTeams: *e}, nil
}

// Delete permet de supprimer une entrée EventTeams par ID
func (repo *EventTeamsService) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.EventTeams.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// List permet de récupérer toutes les entrées EventTeams
func (repo *EventTeamsService) List(ctx context.Context) ([]*ent.EventTeams, error) {
	return repo.db.EventTeams.Query().All(ctx)
}
