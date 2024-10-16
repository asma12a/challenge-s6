package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/footevent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
)

type FootEvent struct {
	db *ent.Client
}

func NewFootEventService(client *ent.Client) *FootEvent {
	return &FootEvent{
		db: client,
	}
}

func (repo *FootEvent) Create(ctx context.Context, footEvent *entity.FootEvent) (*ent.FootEvent, error) {
	entFootEvent, err := repo.db.FootEvent.Create().
		SetEventID(footEvent.EventID).
		SetTeamAID(footEvent.TeamAID).
		SetTeamBID(footEvent.TeamBID).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeCreated
	}

	return entFootEvent, nil
}

func (repo *FootEvent) FindOne(ctx context.Context, id ulid.ID) (*ent.FootEvent, error) {
	return repo.db.FootEvent.Query().Where(footevent.IDEQ(id)).Only(ctx)
}

func (repo *FootEvent) Update(ctx context.Context, footEvent *entity.FootEvent) (*ent.FootEvent, error) {
	entFootEvent, err := repo.db.FootEvent.
		UpdateOneID(footEvent.ID).
		SetEventID(footEvent.EventID).
		SetTeamAID(footEvent.TeamAID).
		SetTeamBID(footEvent.TeamBID).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return entFootEvent, nil
}

func (repo *FootEvent) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.FootEvent.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (repo *FootEvent) List(ctx context.Context) ([]*ent.FootEvent, error) {
	return repo.db.FootEvent.Query().All(ctx)
}
