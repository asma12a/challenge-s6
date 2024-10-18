package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Event struct {
	ent.Event
	EventTypeID ulid.ID `json:"event_type_id"` // Ajoutez cette ligne
	SportID     ulid.ID `json:"sport_id"`      // Assurez-vous que cela existe aussi
}

func NewEvent(name string, address string, eventCode int16, date string, eventTypeId ulid.ID, sportId ulid.ID) *Event {
	return &Event{
		Event: ent.Event{
			Name:      name,
			Address:   address,
			EventCode: eventCode,
			Date:      date,
		},
		EventTypeID: eventTypeId,
		SportID:     sportId,
	}
}
