package presenter

import "time"

// User data
type Event struct {
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// EventCode holds the value of the "event_code" field.
	EventCode int16 `json:"event_code,omitempty"`
	// Date holds the value of the "date" field.
	Date string `json:"date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// IsPublic holds the value of the "is_public" field.
	IsPublic bool `json:"is_public,omitempty"`
	// IsFinished holds the value of the "is_finished" field.
	IsFinished bool `json:"is_finished,omitempty"`
	// EventTypeID holds the value of the "event_type_id" field.
	EventTypeID string `json:"event_type_id,omitempty"`
	// SportID holds the value of the "sport_id" field.
	SportID string `json:"sport_id,omitempty"`
}
