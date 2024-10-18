package presenter

import (
	"time"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Message struct {
	ID        ulid.ID   `json:"id"`
	EventID   ulid.ID   `json:"event_id"`
	UserID    ulid.ID   `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}
