package presenter

import (
	"time"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Event représente le type de présentation pour un événement.
type Event struct {
	ID         ulid.ID   `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Address    string    `json:"address,omitempty"`
	EventCode  int16     `json:"event_code,omitempty"`
	Date       string    `json:"date,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	IsPublic   bool      `json:"is_public,omitempty"`
	IsFinished bool      `json:"is_finished,omitempty"`
	Sport      Sport     `json:"sport_id,omitempty"`      // Utiliser le type Sport personnalisé
}
