package presenter

import (
	"time"

	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Event représente le type de présentation pour un événement.
type Event struct {
	ID         ulid.ID          `json:"id,omitempty"`
	Name       string           `json:"name,omitempty"`
	Address    string           `json:"address,omitempty"`
	Latitude   float64          `json:"latitude,omitempty"`
	Longitude  float64          `json:"longitude,omitempty"`
	EventCode  string           `json:"event_code,omitempty"`
	Date       string           `json:"date,omitempty"`
	CreatedAt  time.Time        `json:"created_at,omitempty"`
	CreatedBy  ulid.ID          `json:"created_by,omitempty"`
	IsPublic   bool             `json:"is_public,omitempty"`
	IsFinished bool             `json:"is_finished,omitempty"`
	EventType  *event.EventType `json:"event_type,omitempty"` // Utiliser le type EventType personnalisé
	Sport      Sport            `json:"sport,omitempty"`
	Teams      []Team           `json:"teams,omitempty"`
}
