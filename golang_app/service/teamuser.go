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

// @Summary Create a new TeamUser
// @Description Create a new team-user relation with userID and teamID
// @Tags team_users
// @Accept  json
// @Produce  json
// @Param teamUser body entity.TeamUser true "Team-User relation to be created"
// @Success 201 {object} entity.TeamUser "TeamUser created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /team_users [post]
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

// @Summary Get a TeamUser by userID and teamID
// @Description Get a specific team-user relation by userID and teamID
// @Tags team_users
// @Accept  json
// @Produce  json
// @Param userID path string true "User ID"
// @Param teamID path string true "Team ID"
// @Success 200 {object} entity.TeamUser "TeamUser details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "TeamUser Not Found"
// @Router /team_users/{userID}/{teamID} [get]
func (repo *TeamUser) FindOne(ctx context.Context, userID ulid.ID, teamID ulid.ID) (*entity.TeamUser, error) {
	teamUser, err := repo.db.TeamUser.Query().Where(teamuser.IDEQ(userID), teamuser.IDEQ(teamID)).Clone().WithUser().WithTeam().
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.TeamUser{TeamUser: *teamUser}, nil
}

// @Summary Update a TeamUser relation
// @Description Update a specific team-user relation by userID and teamID
// @Tags team_users
// @Accept  json
// @Produce  json
// @Param id path string true "TeamUser ID"
// @Param teamUser body entity.TeamUser true "Updated Team-User relation"
// @Success 200 {object} entity.TeamUser "Updated TeamUser"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "TeamUser Not Found"
// @Router /team_users/{id} [put]
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

// @Summary Delete a TeamUser relation
// @Description Delete a team-user relation by ID
// @Tags team_users
// @Accept  json
// @Produce  json
// @Param id path string true "TeamUser ID"
// @Success 200 {object} map[string]interface{} "TeamUser deleted"
// @Failure 404 {object} map[string]interface{} "TeamUser Not Found"
// @Router /team_users/{id} [delete]
func (repo *TeamUser) Delete(ctx context.Context, id ulid.ID) error {
	err := repo.db.TeamUser.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// @Summary List all TeamUser relations
// @Description Get a list of all team-user relations
// @Tags team_users
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.TeamUser "List of TeamUser relations"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /team_users [get]
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
