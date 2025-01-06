package presenter

import (
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type User struct {
	ID    ulid.ID  `json:"id,omitempty"`
	Name  string   `json:"name,omitempty"`
	Email string   `json:"email,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
