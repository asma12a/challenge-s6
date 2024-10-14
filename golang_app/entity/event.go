package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type Event struct {
	ent.Event
}

func NewEvent(name string, address string, eventCode int16, date string, isPublic bool, isFinished bool) *Event {
	return &Event{
		Event: ent.Event{
			Name:       name,
			Address:    address,
			EventCode:  eventCode,
			Date:       date,
			IsPublic:   isPublic,
			IsFinished: isFinished,
		},
	}
}
