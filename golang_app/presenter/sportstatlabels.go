package presenter

import (
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)


type SportStatLabels struct {
	ID         ulid.ID          `json:"id,omitempty"`
	Label      string           `json:"label,omitempty"`
	Unit       string           `json:"unit,omitempty"`
	IsMain     bool             `json:"is_main,omitempty"`
	Sport      *Sport            `json:"sport,omitempty"` 
}
