package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type EventTeams struct {
	ent.EventTeams
}

func NewEventTeams(eventId ulid.ID, teamId ulid.ID) *EventTeams {
	return &EventTeams{
		EventTeams: ent.EventTeams{
			EventID: eventId,
			TeamID:  teamId},
	}

}
