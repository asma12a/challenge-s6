package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type EventType struct {
	ent.EventType
}

func NewEventType(name string) *EventType {
	return &EventType{
		EventType: ent.EventType{
			Name: name,
		},
	}
}
