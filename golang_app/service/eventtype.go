package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/eventtype"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
)

type EventType struct {
	db *ent.Client
}

func NewEventTypeService(client *ent.Client) *EventType {
	return &EventType{
		db: client,
	}
}

func (repo *EventType) Create(ctx context.Context, et *entity.EventType) error {
	_, err := repo.db.EventType.Create().
		SetName(et.Name).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

func (e *EventType) FindOne(ctx context.Context, id ulid.ID) (*entity.EventType, error) {
	et, err := e.db.EventType.Query().Where(eventtype.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.EventType{EventType: *et}, nil
}

func (repo *EventType) Update(ctx context.Context, et *entity.EventType) (*entity.EventType, error) {

	// Prepare the update query
	eventType, err := repo.db.EventType.
		UpdateOneID(et.ID).
		SetName(et.Name).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.EventType{EventType: *eventType}, nil
}

func (et *EventType) Delete(ctx context.Context, id ulid.ID) error {
	err := et.db.EventType.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (et *EventType) List(ctx context.Context) ([]*ent.EventType, error) {
	return et.db.EventType.Query().All(ctx)
}
