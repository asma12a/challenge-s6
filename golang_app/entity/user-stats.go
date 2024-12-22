package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type UserStats struct {
	ent.UserStats
	UserID  ulid.ID `json:"user_id"`
	EventID ulid.ID `json:"event_id"`
}

func NewUserStats(userID, eventID ulid.ID) *UserStats {
	return &UserStats{
		UserStats: ent.UserStats{},
		UserID:    userID,
		EventID:   eventID,
	}
}
