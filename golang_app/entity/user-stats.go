package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type UserStats struct {
	ent.UserStats
}

func NewUserStats(userID, eventID ulid.ID) *UserStats {
	return &UserStats{
		UserStats: ent.UserStats{
			UserID:  userID,
			EventID: eventID,
		},
	}
}
