package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/sport"
	"github.com/asma12a/challenge-s6/entity"
)

type Sport struct {
	db *ent.Client
}

func NewSportService(client *ent.Client) *Sport {
	return &Sport{
		db: client,
	}
}

func (repo *Sport) Create(ctx context.Context, sport *entity.Sport) error {
	_, err := repo.db.Sport.Create().
		SetName(sport.Name).
		SetImageURL(sport.ImageURL).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

func (sp *Sport) FindOne(ctx context.Context, id string) (*entity.Sport, error) {
	sport, err := sp.db.Sport.Query().Where(sport.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Sport{Sport: *sport}, nil
}

func (repo *Sport) Update(ctx context.Context, sp *entity.Sport) (*entity.Sport, error) {

	// Prepare the update query
	sport, err := repo.db.Sport.
		UpdateOneID(sp.ID).
		SetName(sp.Name).
		SetImageURL(sp.ImageURL).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Sport{Sport: *sport}, nil
}

func (sp *Sport) Delete(ctx context.Context, id string) error {
	err := sp.db.Sport.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (sp *Sport) List(ctx context.Context) ([]*ent.Sport, error) {
	return sp.db.Sport.Query().All(ctx)
}
