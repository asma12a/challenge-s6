package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
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

func (repo *Event) Create(ctx context.Context, event *entity.Event) error {

	create := repo.db.Event.Create().
		SetName(event.Name).
		SetAddress(event.Address).
		SetEventCode(event.EventCode).
		SetDate(event.Date).
		SetSportID(event.SportID)

	if event.EventType != nil {
		create.SetEventType(*event.EventType)
	}

	_, err := create.Save(ctx)
	if err != nil {
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
