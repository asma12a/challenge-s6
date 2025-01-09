package entity

import (
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type ActionLog struct {
	ent.ActionLog
	UserID    *ulid.ID  `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func NewActionLog(
	userID *ulid.ID, action string, description string,
) *ActionLog {
	return &ActionLog{
		ActionLog: ent.ActionLog{
			Action:      action,
			Description: description,
			CreatedAt:   time.Now(),
		},
		UserID: userID,
	}
}
