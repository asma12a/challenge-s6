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

// @Summary Create a new UserStats
// @Description Create a new UserStats entry with userID and eventID
// @Tags user_stats
// @Accept  json
// @Produce  json
// @Param userStats body entity.UserStats true "UserStats to be created"
// @Success 201 {object} entity.UserStats "UserStats created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /user_stats [post]
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
