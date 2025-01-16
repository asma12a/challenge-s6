package service

import (
	"context"
	"log"

	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/event"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/team"
	"github.com/asma12a/challenge-s6/ent/teamuser"
	"github.com/asma12a/challenge-s6/ent/user"
	"github.com/asma12a/challenge-s6/entity"
	"github.com/redis/go-redis/v9"

)

type TeamUser struct {
	db *ent.Client
}

func NewTeamUserService(client *ent.Client) *TeamUser {
	return &TeamUser{
		db: client,
	}
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
func (repo *TeamUser) FindOne(ctx context.Context, id ulid.ID) (*entity.TeamUser, error) {
	teamUser, err := repo.db.TeamUser.Query().Where(teamuser.IDEQ(id)).WithTeam().Only(ctx)

	if err != nil {
		return nil, entity.ErrNotFound
	}
	return &entity.TeamUser{TeamUser: *teamUser, TeamID: teamUser.Edges.Team.ID}, nil
}

func (e *TeamUser) AddPlayerToTeam(ctx context.Context, teamUserInput entity.TeamUser, eventID ulid.ID) error {
	// Check if the team exists and get the current number of players
	teamFound, err := e.db.Team.Query().
		Where(team.IDEQ(teamUserInput.TeamID)).
		WithTeamUsers().
		Only(ctx)
	if err != nil {
		return entity.ErrEntityNotFound("Team")
	}

	// Check if the team has a max players limit and if it's full
	if teamFound.MaxPlayers > 0 && len(teamFound.Edges.TeamUsers) >= teamFound.MaxPlayers {
		return entity.ErrTeamFull
	}

	// Check if the user is already in another team for the same event
	existingTeamUser, err := e.db.TeamUser.Query().
		Where(
			teamuser.HasTeamWith(team.HasEventWith(event.IDEQ(eventID))),
			teamuser.HasUserWith(user.EmailEQ(teamUserInput.Email)),
		).
		Only(ctx)
	if err == nil && existingTeamUser != nil {
		return entity.ErrUserAlreadyInATeam
	}

	tx, err := e.db.Tx(ctx)
	if err != nil {
		log.Println(err, "error creating transaction")
		return err
	}
	defer tx.Rollback()

	userFound, err := tx.User.Query().Where(user.EmailEQ(teamUserInput.Email)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
		// User not found, send an email to the provided email
		log.Println("user not found, sending email")
		// TODO: Send email
		// err = sendEmail(teamUserInput.Email)
		// if err != nil {
		// 	log.Println("error sending email:", err)
		// 	return err
		// }
	}

	teamUserCreate := tx.TeamUser.Create().
		SetTeamID(teamUserInput.TeamID).
		SetEmail(teamUserInput.Email).
		SetStatus("pending")

	if userFound != nil {
		teamUserCreate = teamUserCreate.SetUserID(userFound.ID).SetStatus("valid")
	}

	if teamUserInput.Role != "" {
		teamUserCreate = teamUserCreate.SetRole(teamUserInput.Role)
	}

	if _, err = teamUserCreate.Save(ctx); err != nil {
		log.Println("error adding player to team:", err)
		return err
	}
	if err := tx.Commit(); err != nil {
		log.Println("error committing transaction:", err)
		return err
	}

	return nil
}

// Can update team_id, role, status (if update team, check if team is full)
func (e *TeamUser) UpdatePlayer(ctx context.Context, teamUserInput entity.TeamUser) error {
	// Check if the team exists and get the current number of players
	teamFound, err := e.db.Team.Query().
		Where(team.IDEQ(teamUserInput.TeamID)).
		WithTeamUsers().
		Only(ctx)

	if err != nil {
		return entity.ErrEntityNotFound("Team")
	}

	// Check if the team has a max players limit and if it's full
	if teamFound.MaxPlayers > 0 && len(teamFound.Edges.TeamUsers) >= teamFound.MaxPlayers {
		return entity.ErrTeamFull
	}

	teamUserUpdate := e.db.TeamUser.UpdateOneID(teamUserInput.ID).
		SetTeamID(teamUserInput.TeamID)

	if teamUserInput.Role != "" {
		teamUserUpdate = teamUserUpdate.SetRole(teamUserInput.Role)
	}

	_, err = teamUserUpdate.Save(ctx)

	if err != nil {
		return err
	}

	return nil
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

func (repo *TeamUser) UpdateTeamUserWithUser(ctx context.Context, existingUser entity.User, serviceNotification NotificationService, rdb *redis.Client) error {
	teamUsers, err := repo.db.TeamUser.Query().Where(teamuser.EmailEQ(existingUser.Email)).
	WithTeam(func(q *ent.TeamQuery) { 
        q.WithEvent() 
    }).
	All(ctx)
	log.Println("teamUsers", teamUsers)
	if err != nil {
		return err
	}

	for _, teamUser := range teamUsers {
		_, err := repo.db.TeamUser.UpdateOneID(teamUser.ID).
			SetUserID(existingUser.ID).
			SetStatus("valid").
			Save(ctx)
		if err != nil {
			return err
		}
		createdBy := teamUser.Edges.Team.Edges.Event.CreatedBy
		createdByUser, err := repo.db.User.Query().Where(user.IDEQ(createdBy)).Only(ctx)
		if err != nil {
			return err
		}
		fcmToken, err := serviceNotification.GetTokenFromRedis(ctx,rdb, string(createdByUser.ID)+"_FCM") 
		if err == nil {
			serviceNotification.SendPushNotification(fcmToken,"Information", "Votre invité "+ existingUser.Email +" nous a rejoint")

		}
	}

	return nil
}
