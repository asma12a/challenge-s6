package service

import (
	"context"
	"time"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/entity"
)

type ActionLogService struct {
	db *ent.Client
}

// NewActionLogService crée un nouveau service pour l'entité ActionLog
func NewActionLogService(client *ent.Client) *ActionLogService {
	return &ActionLogService{
		db: client,
	}
}

// Create crée un nouvel ActionLog
func (repo *ActionLogService) Create(ctx context.Context, actionLog *entity.ActionLog) error {
	_, err := repo.db.ActionLog.Create().
		SetUserID(*actionLog.UserID).
		SetAction(actionLog.Action).
		SetDescription(actionLog.Description).
		SetCreatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

// List permet de lister tous les ActionLogs
func (repo *ActionLogService) List(ctx context.Context) ([]*entity.ActionLog, error) {
	logs, err := repo.db.ActionLog.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*entity.ActionLog
	for _, log := range logs {
		result = append(result, &entity.ActionLog{ActionLog: *log})
	}

	return result, nil
}
