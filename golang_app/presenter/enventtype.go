package presenter

import "github.com/asma12a/challenge-s6/ent/schema/ulid"

// User data
type EventType struct {
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}
