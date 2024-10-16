package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type Event struct {
	ent.Event
	EventTypeID string `json:"event_type_id"` // Ajoutez cette ligne
	SportID     string `json:"sport_id"`      // Assurez-vous que cela existe aussi
}

func NewEvent(name string, address string, eventCode int16, date string, eventTypeId string, sportId string) *Event {
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
