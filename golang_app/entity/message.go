package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Message struct {
	ent.Message
}

func NewMessage(eventId ulid.ID, teamId ulid.ID, content string, createdAt date ) *Message {
	return &Message{
		Message: ent.Message{
			EventID:   eventId,
			UserId:    userId,
			Content:   content,
			CreatedAt: createdAt,
		},
	}
}
