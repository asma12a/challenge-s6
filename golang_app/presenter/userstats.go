package presenter

import (
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)


type UserStats struct {
	ID         ulid.ID          `json:"id,omitempty"`
	StatLabel  *SportStatLabels  `json:"stat_label,omitempty"`
	Value      int			  `json:"value,omitempty"`
	User       *User             `json:"user,omitempty"`
}
