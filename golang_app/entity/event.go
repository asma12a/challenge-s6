package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Event struct {
	ent.Event
	SportID ulid.ID     `json:"sport_id"`
}

func NewEvent(name string, address string, date string, sportId ulid.ID, eventType *event.EventType) *Event {
	event := &Event{
		Event: ent.Event{
			Name:      name,
			EventCode: GenerateEventCode(),
			Date:      date,
			EventType: eventType,
			Address:   address,
		},
		SportID: sportId,
	}

	return event
}

func GenerateEventCode() string {
	return "TEST"
}
