package presenter

import "github.com/asma12a/challenge-s6/ent/schema/ulid"

// User data
type Sport struct {
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// MaxTeams holds the value of the "max_teams" field.
	MaxTeams int `json:"max_teams,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
}
