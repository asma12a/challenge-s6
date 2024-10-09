package presenter

import (
	"time"

	"github.com/google/uuid"
)

// User data
type Event struct {
	ID         uuid.UUID `json:"id,omitempty"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	EventCode  int16     `json:"event_code"`
	Date       time.Time `json:"date"`
	IsPublic   bool      `json:"is_public"`
	IsFinished bool      `json:"is_finished"`
}
