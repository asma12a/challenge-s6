package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Event struct {
	ent.Event
	SportID ulid.ID `json:"sport_id" validate:"required"`
}

func NewEvent(name string, address string, eventCode string, date string, eventType *event.EventType, sportId ulid.ID) *Event {
	return &Event{
		Event: ent.Event{
			Name:      name,
			Address:   address,
			EventCode: eventCode,
			Date:      date,
			EventType: eventType,
		},
		SportID: sportId,
	}
}
