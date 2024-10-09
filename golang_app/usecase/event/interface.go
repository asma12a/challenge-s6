package event

import (
	"context"
	"time"

	"github.com/asma12a/challenge-s6/entity"
	"github.com/google/uuid"
)

// Reader interface
type Reader interface {
	Get(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	Search(ctx context.Context, query string) ([]*entity.Event, error)
	List(ctx context.Context) ([]*entity.Event, error)
}

// Writer user writer
type Writer interface {
	Create(ctx context.Context, e *entity.Event) (*entity.Event, error)
	Update(ctx context.Context, e *entity.Event) (*entity.Event, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreateEvent(ctx context.Context, name string, address string, eventCode int16, date time.Time, isPublic bool, isFinished bool) (*entity.Event, error)
	GetEvent(ctx context.Context, id uuid.UUID) (*entity.Event, error)
	UpdateEvent(ctx context.Context, e *entity.Event) (*entity.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	SearchEvents(ctx context.Context, query string) ([]*entity.Event, error)
	ListEvents(ctx context.Context) ([]*entity.Event, error)
}
