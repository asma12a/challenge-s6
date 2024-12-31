package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type UserStats struct {
	ent.UserStats
}

func NewUserStats(userID, eventID, statId ulid.ID, statValue int) *UserStats {
	return &UserStats{
		UserStats: ent.UserStats{
			UserID:  userID,
			EventID: eventID,
			StatID:  statId,
			StatValue: statValue,
		},
	}
}
