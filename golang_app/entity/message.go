package entity

import (
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Message struct {
	ent.Message
	EventID ulid.ID `json:"event_id"`
	UserID  ulid.ID `json:"user_id"`
}

func NewMessage(eventId ulid.ID, userId ulid.ID, userName string, content string, createdAt time.Time) *Message {
	return &Message{
		Message: ent.Message{
			UserName:  userName,
			Content:   content,
			CreatedAt: createdAt,
		},
		EventID: eventId,
		UserID:  userId,
	}
}
