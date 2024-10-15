package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
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

func (repo *Event) Create(ctx context.Context, event *entity.Event) (*ent.Event, error) {
	entEvent, err := repo.db.Event.Create().
		SetName(event.Name).
		SetAddress(event.Address).
		SetEventCode(event.EventCode).
		SetDate(event.Date).
		SetIsPublic(event.IsPublic).
		SetIsFinished(event.IsFinished).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeCreated
	}

	return entEvent, nil
}

func (e *Event) FindOne(ctx context.Context, id entity.ID) (*ent.Event, error) {
	return e.db.Event.Query().Where(event.IDEQ(id)).Only(ctx)
}

// Update a pet
func (repo *Event) Update(ctx context.Context, event *entity.Event) (*ent.Event, error) {

	// Prepare the update query
	entEvent, err := repo.db.Event.
		UpdateOneID(event.ID).
		SetName(event.Name).
		SetAddress(event.Address).
		SetEventCode(event.EventCode).
		SetDate(event.Date).
		SetIsPublic(event.IsPublic).
		SetIsFinished(event.IsFinished).Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return entEvent, nil
}

func (e *Event) Delete(ctx context.Context, id entity.ID) error {
	err := e.db.Event.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (e *Event) List(ctx context.Context) ([]*ent.Event, error) {
	return e.db.Event.Query().All(ctx)
}
