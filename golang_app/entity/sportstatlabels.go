package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"

)

type SportStatLabels struct {
	ent.SportStatLabels
	SportID ulid.ID     `json:"sport_id"`
}

func NewSportStatLabels(label string, unit string, isMain bool, sportId ulid.ID) *SportStatLabels {
	sportStatLabels := &SportStatLabels{
		SportStatLabels: ent.SportStatLabels{
			Label:  label,
			Unit:   unit,
			IsMain: isMain,
		},
		SportID: sportId,
	}

	return sportStatLabels
}
