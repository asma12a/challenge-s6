package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/team"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/entity"
)

type Team struct {
	db *ent.Client
}

func NewTeamService(client *ent.Client) *Team {
	return &Team{
		db: client,
	}
}

func (repo *Team) Create(ctx context.Context, t *entity.Team) error {
	_, err := repo.db.Team.Create().
		SetName(t.Name).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

func (t *Team) FindOne(ctx context.Context, id ulid.ID) (*entity.Team, error) {
	team, err := t.db.Team.Query().Where(team.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Team{Team: *team}, nil
}

func (repo *Team) Update(ctx context.Context, t *entity.Team) (*entity.Team, error) {

	team, err := repo.db.Team.
		UpdateOneID(t.ID).
		SetName(t.Name).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Team{Team: *team}, nil
}

func (t *Team) Delete(ctx context.Context, id ulid.ID) error {
	err := t.db.Team.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (t *Team) List(ctx context.Context) ([]*ent.Team, error) {
	return t.db.Team.Query().All(ctx)
}
