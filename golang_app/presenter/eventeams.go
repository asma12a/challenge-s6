package presenter

import (
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type EventTeams struct {
	ID      ulid.ID `json:"id"`
	EventID ulid.ID `json:"event_id"`
	TeamID  ulid.ID `json:"team_id"`
}
