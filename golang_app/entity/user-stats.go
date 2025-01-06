package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type UserStats struct {
	ent.UserStats
	UserID  ulid.ID `json:"user_id"`
	EventID ulid.ID `json:"event_id"`
	StatID  ulid.ID `json:"stat_id"`
}

func NewUserStats(userID, eventID, statId ulid.ID, statValue int) *UserStats {
	return &UserStats{
		UserStats: ent.UserStats{
			StatValue: statValue,
		},
		UserID:  userID,
		EventID: eventID,
		StatID:  statId,
	}
}
