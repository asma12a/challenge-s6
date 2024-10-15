package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type Event struct {
	ent.Event
}

func NewEvent(name string, address string, eventCode int16, date string, eventType *ent.EventType) *Event {
	return &Event{
		Event: ent.Event{
			Name:      name,
			Address:   address,
			EventCode: eventCode,
			Date:      date,
			Edges: ent.EventEdges{
				EventType: eventType,
			},
		},
	}
}
