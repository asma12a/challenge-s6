package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/teamuser"
	"github.com/asma12a/challenge-s6/entity"
)

type TeamUser struct {
	db *ent.Client
}

func NewTeamUserService(client *ent.Client) *TeamUser {
	return &TeamUser{
		db: client,
	}
}

func (repo *TeamUser) Create(ctx context.Context, teamUser *entity.TeamUser) error {
	_, err := repo.db.TeamUser.Create().
		SetUserID(teamUser.UserID).
		SetTeamID(teamUser.TeamID).
		Save(ctx)

	if err != nil {
		return entity.ErrCannotBeCreated
	}

	return nil
}

func (repo *TeamUser) FindOne(ctx context.Context, userID ulid.ID, teamID ulid.ID) (*entity.TeamUser, error) {
	teamUser, err := repo.db.TeamUser.Query().Where(teamuser.IDEQ(userID), teamuser.IDEQ(teamID)).Clone().WithUser().WithTeam().
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.TeamUser{TeamUser: *teamUser}, nil
}

func (repo *TeamUser) Update(ctx context.Context, teamUser *entity.TeamUser) (*entity.TeamUser, error) {

	// Prepare the update query
	e, err := repo.db.TeamUser.
		UpdateOneID(teamUser.ID).
		SetUserID(teamUser.UserID).
		SetTeamID(teamUser.TeamID).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrCannotBeUpdated
	}

	return &entity.TeamUser{TeamUser: *e}, nil
}

func (repo *TeamUser) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.TeamUser.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

func (repo *TeamUser) List(ctx context.Context) ([]*entity.TeamUser, error) {
	teamUsers, err := repo.db.TeamUser.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*entity.TeamUser
	for _, tu := range teamUsers {
		result = append(result, &entity.TeamUser{TeamUser: *tu})
	}

	return result, nil
}
