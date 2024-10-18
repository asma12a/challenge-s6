package entity

import (
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Message struct {
	ent.Message
}

func NewMessage(eventId ulid.ID, userId ulid.ID, content string, createdAt time.Time) *Message {
	return &Message{
		Message: ent.Message{
			EventID:   eventId,
			UserID:    userId,
			Content:   content,
			CreatedAt: createdAt,
		},
	}
}
