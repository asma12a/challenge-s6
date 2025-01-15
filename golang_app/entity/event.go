package entity

import (
	"math/rand"
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	geo "github.com/kellydunn/golang-geo"
)

type Event struct {
	ent.Event
	SportID ulid.ID `json:"sport_id"`
}

func NewEvent(name string, address string, latitude, longitude float64, date time.Time, sportId ulid.ID, eventType *event.EventType) *Event {
	event := &Event{
		Event: ent.Event{
			Name:      name,
			EventCode: GenerateEventCode(),
			Date:      date,
			EventType: eventType,
			Address:   address,
			Latitude:  latitude,
			Longitude: longitude,
		},
		SportID: sportId,
	}

	return event
}

func GenerateEventCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	code := make([]byte, length)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}

	return string(code)
}

func SortEventsByDistance(events []*ent.Event, latitude, longitude float64) []*ent.Event {
	referencePoint := geo.NewPoint(latitude, longitude)
	points := make([]*geo.Point, len(events))
	for i, event := range events {
		points[i] = geo.NewPoint(event.Latitude, event.Longitude)
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if referencePoint.GreatCircleDistance(points[i]) > referencePoint.GreatCircleDistance(points[j]) {
				points[i], points[j] = points[j], points[i]
				events[i], events[j] = events[j], events[i]
			}
		}
	}

	return events
}
