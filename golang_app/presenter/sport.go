package presenter

import "github.com/asma12a/challenge-s6/ent/schema/ulid"

// User data
type Sport struct {
	ID ulid.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ImageURL holds the value of the "image_url" field.
	ImageURL string `json:"image_url,omitempty"`
}
