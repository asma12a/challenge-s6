package service

import (
	"context"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/team"
	"github.com/asma12a/challenge-s6/ent/teamuser"
	"github.com/asma12a/challenge-s6/ent/user"
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

// @Summary Create a new team
// @Description Create a new team with name and max players
// @Tags teams
// @Accept  json
// @Produce  json
// @Param team body entity.Team true "Team to be created"
// @Success 201 {object} entity.Team "Team created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /teams [post]
func (e *Team) AddTeam(ctx context.Context, eventID ulid.ID, teamInput entity.Team) error {
	tx, err := e.db.Tx(ctx)
	if err != nil {
		return err
	}

	eventFound, err := tx.Event.Query().Where(event.IDEQ(eventID)).WithSport().Only(ctx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	existingTeams, err := tx.Team.Query().
		Where(team.HasEventWith(event.IDEQ(eventID))).
		All(ctx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if eventFound.Edges.Sport.MaxTeams != nil && len(existingTeams) >= *eventFound.Edges.Sport.MaxTeams {
		_ = tx.Rollback()
		return entity.ErrCannotBeCreated
	}

	for _, existingTeam := range existingTeams {
		if existingTeam.Name == teamInput.Name {
			_ = tx.Rollback()
			return entity.ErrCannotBeCreated
		}
	}

	_, err = tx.Team.Create().
		SetName(teamInput.Name).
		SetMaxPlayers(teamInput.MaxPlayers).
		SetEventID(eventID). // Assuming Team has an EventID field
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (t *Team) FindAll(ctx context.Context, eventID ulid.ID) ([]*ent.Team, error) {
	return t.db.Team.Query().Where(team.HasEventWith(event.IDEQ(eventID))).
		WithTeamUsers(func(q *ent.TeamUserQuery) {
			q.WithUser()
		}).
		All(ctx)
}

// @Summary Get a team by ID
// @Description Get a specific team by its ID
// @Tags teams
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Success 200 {object} entity.Team "Team details"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Team Not Found"
// @Router /teams/{id} [get]
func (t *Team) FindOne(ctx context.Context, id ulid.ID) (*entity.Team, error) {
	team, err := t.db.Team.Query().Where(team.IDEQ(id)).
		Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.Team{Team: *team}, nil
}

// @Summary Update a team
// @Description Update a team's details by ID
// @Tags teams
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Param team body entity.Team true "Updated team data"
// @Success 200 {object} entity.Team "Updated team"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 404 {object} map[string]interface{} "Team Not Found"
// @Router /teams/{id} [put]
func (repo *Team) Update(ctx context.Context, t *entity.Team) (*entity.Team, error) {

	team, err := repo.db.Team.
		UpdateOneID(t.ID).
		SetName(t.Name).
		SetMaxPlayers(t.MaxPlayers).
		Save(ctx)

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return &entity.Team{Team: *team}, nil
}

// @Summary Delete a team
// @Description Delete a team by ID
// @Tags teams
// @Accept  json
// @Produce  json
// @Param id path string true "Team ID"
// @Success 200 {object} map[string]interface{} "Team deleted"
// @Failure 404 {object} map[string]interface{} "Team Not Found"
// @Router /teams/{id} [delete]
func (t *Team) Delete(ctx context.Context, id ulid.ID) error {
	err := t.db.Team.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return entity.ErrCannotBeDeleted
	}
	return nil
}

// @Summary List all teams
// @Description Get a list of all teams
// @Tags teams
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.Team "List of teams"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /teams [get]

func (t *Team) List(ctx context.Context) ([]*ent.Team, error) {
	return t.db.Team.Query().All(ctx)
}

func (e *Team) JoinTeam(ctx context.Context, eventID, teamID, userID ulid.ID) error {
	// Check if the team exists and get the current number of players
	teamFound, err := e.db.Team.Query().
		Where(team.IDEQ(teamID)).
		WithTeamUsers().
		Only(ctx)
	if err != nil {
		return entity.ErrEntityNotFound("Team")
	}

	// Check if the team has a max players limit and if it's full
	if teamFound.MaxPlayers > 0 && len(teamFound.Edges.TeamUsers) >= teamFound.MaxPlayers {
		return entity.ErrTeamFull
	}

	// Check if the user exists
	userFound, uErr := e.db.User.Get(ctx, userID)
	if uErr != nil {
		return entity.ErrNotFound
	}

	// Check if the user is already in another team for the same event
	existingTeamUser, err := e.db.TeamUser.Query().
		Where(
			teamuser.HasTeamWith(team.HasEventWith(event.IDEQ(eventID))),
			teamuser.HasUserWith(user.IDEQ(userID)),
		).
		Only(ctx)
	if err == nil && existingTeamUser != nil {
		return entity.ErrUserAlreadyInATeam
	}

	// Add the user to the team
	_, err = e.db.TeamUser.Create().
		SetTeamID(teamID).
		SetUserID(userID).
		SetEmail(userFound.Email).
		SetStatus("valid").
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (e *Team) SwitchTeam(ctx context.Context, eventID, teamID, userID ulid.ID) error {
	// Check if the team exists and get the current number of players
	teamFound, err := e.db.Team.Query().
		Where(team.IDEQ(teamID)).
		WithTeamUsers().
		Only(ctx)
	if err != nil {
		return entity.ErrNotFound
	}

	// Check if the team has a max players limit and if it's full
	if teamFound.MaxPlayers > 0 && len(teamFound.Edges.TeamUsers) >= teamFound.MaxPlayers {
		return entity.ErrTeamFull
	}

	// Check if the user exists
	userFound, uErr := e.db.User.Get(ctx, userID)
	if uErr != nil {
		return entity.ErrNotFound
	}

	// Check if the user is already in another team for the same event
	existingTeamUser, err := e.db.TeamUser.Query().
		Where(
			teamuser.HasTeamWith(team.HasEventWith(event.IDEQ(eventID))),
			teamuser.HasUserWith(user.IDEQ(userID)),
		).
		WithTeam().
		Only(ctx)
	if err == nil && existingTeamUser != nil {
		// Check if the user is already in the requested team
		if existingTeamUser.Edges.Team.ID == teamID {
			return entity.ErrUserAlreadyInThisTeam
		}
		// Remove the user from the current team
		err = e.db.TeamUser.DeleteOne(existingTeamUser).Exec(ctx)
		if err != nil {
			return err
		}
	}

	// Add the user to the new team
	_, err = e.db.TeamUser.Create().
		SetTeamID(teamID).
		SetUserID(userID).
		SetEmail(userFound.Email).
		SetStatus("valid").
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
