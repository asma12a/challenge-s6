package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type FootEvent struct {
	ent.FootEvent
}

func NewFootEvent(eventID string, teamAID string, teamBID string) *FootEvent {
	return &FootEvent{
		FootEvent: ent.FootEvent{
			EventID: eventID,
			TeamAID: teamAID,
			TeamBID: teamBID,
		},
	}
}
