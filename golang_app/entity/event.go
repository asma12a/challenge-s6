package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)



type Event struct {
	ent.Event 
	SportID     ulid.ID `json:"sport_id"` 
	Teams       []*ent.Team `json:"teams,omitempty"`  
}

func NewEvent(name string, address string, eventCode string, date string, sportId ulid.ID, eventType event.EventType, teams []*ent.Team ) *Event {
	return &Event{
		Event: ent.Event{
			Name:      name,
			EventCode: eventCode,
			Date:      date,
			EventType: &eventType,
			Address:  address,
		},
		SportID:     sportId,
		Teams:       teams,
	}
}
