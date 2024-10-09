package event

import (
	"context"
	"time"

	"github.com/asma12a/challenge-s6/entity"
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateEvent Create a Event
func (s *Service) CreateEvent(ctx context.Context, name string, address string, eventCode int16, date time.Time, isPublic bool, isFinished bool) (*entity.Event, error) {
	e := entity.NewEvent(name, address, eventCode, date, isPublic, isFinished)
	return s.repo.Create(ctx, e)
}

// GetEvent Get a Event by ID
func (s *Service) GetEvent(ctx context.Context, id uuid.UUID) (*entity.Event, error) {
	return s.repo.Get(ctx, id)
}

// UpdateEvent Update a Event
func (s *Service) UpdateEvent(ctx context.Context, e *entity.Event) (*entity.Event, error) {
	_, err := s.repo.Get(ctx, e.ID)
	if err != nil {
		return nil, entity.ErrNotFound
	}

	return s.repo.Update(ctx, e)
}

// DeleteEvent Delete a Event
func (s *Service) DeleteEvent(ctx context.Context, id uuid.UUID) error {

	e, err := s.repo.Get(ctx, id)
	if e == nil || err != nil {
		return entity.ErrNotFound
	}
	return s.repo.Delete(ctx, id)
}

// ListEvents List Events
func (s *Service) ListEvents(ctx context.Context) ([]*entity.Event, error) {
	return s.repo.List(ctx)
}

// SearchEvents Search Events
func (s *Service) SearchEvents(ctx context.Context, query string) ([]*entity.Event, error) {
	return s.repo.Search(ctx, query)
}
