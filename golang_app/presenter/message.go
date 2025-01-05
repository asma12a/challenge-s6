package presenter

import (
	"time"

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Message struct {
	ID        ulid.ID   `json:"id"`
	EventID   ulid.ID   `json:"event_id"`
	UserName  string    `json:"user_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
}
