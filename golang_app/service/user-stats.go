package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
)

type UserStats struct {
	db *ent.Client
}

func NewUserStatsService(client *ent.Client) *UserStats {
	return &UserStats{
		db: client,
	}
}

func (repo *UserStats) Create(ctx context.Context, userStats *entity.UserStats) error {

	_, err := repo.db.UserStats.Create().
		SetUserID(userStats.UserID).
		SetEventID(userStats.EventID).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}
