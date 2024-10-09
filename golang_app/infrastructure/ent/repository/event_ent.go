package repository

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/entity"
	usecase "github.com/asma12a/challenge-s6/usecase/event"
	"github.com/google/uuid"
)

// EventRepoEnt Ent repo
type EventRepoEnt struct {
	client *ent.Client
}

// NewEventRepoEnt is specific implementation of the interface
func NewEventRepoEnt(client *ent.Client) usecase.Repository {
	return &EventRepoEnt{
		client: client,
	}
}

// Create a Event
func (r *EventRepoEnt) Create(ctx context.Context, e *entity.Event) (*entity.Event, error) {
	entEvent, err := r.client.Event.Create().
		SetName(e.Name).
		SetAddress(e.Address).
		SetEventCode(e.EventCode).
		SetDate(e.Date).
		SetIsPublic(e.IsPublic).
		SetIsFinished(e.IsFinished).Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeCreated
	}

	Event := &entity.Event{Event: *entEvent}

	return Event, nil
}

// Get a Event
func (r *EventRepoEnt) Get(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	entEvent, err := r.client.Event.
		Query().
		Where(event.IDEQ(id)).
		Only(ctx) // `Only` fails if no Event found, or more than 1 Event returned.

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Event{*entEvent}, nil
}

// Update a Event
func (r *EventRepoEnt) Update(ctx context.Context, e *entity.Event) (*entity.Event, error) {

	// Prepare the update query
	entEvent, err := r.client.Event.
		UpdateOneID(e.ID).
		SetName(e.Name).
		SetAddress(e.Address).
		SetEventCode(e.EventCode).
		SetDate(e.Date).
		SetIsPublic(e.IsPublic).
		SetIsFinished(e.IsFinished).Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Event{Event: *entEvent}, nil
}

// Delete a Event
func (r *EventRepoEnt) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.client.Event.
		DeleteOneID(id).
		Exec(ctx)

	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// List Events
func (r *EventRepoEnt) List(ctx context.Context) ([]*entity.Event, error) {
	entEvents, err := r.client.Event.
		Query().
		All(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}

	events := make([]*entity.Event, len(entEvents))
	for i, e := range entEvents {
		events[i] = &entity.Event{Event: *e}
	}

	return events, nil
}

// Search Events
func (r *EventRepoEnt) Search(ctx context.Context, name string) ([]*entity.Event, error) {
	entEvents, err := r.client.Event.
		Query().
		Where(event.NameEQ(name)).
		All(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}

	events := make([]*entity.Event, len(entEvents))
	for i, e := range entEvents {
		events[i] = &entity.Event{Event: *e}
	}

	return events, nil
}
