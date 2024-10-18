package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type TeamUser struct {
	ent.TeamUser
	UserID ulid.ID `json:"user_id"`
	TeamID ulid.ID `json:"team_id"`
}

func NewTeamUser(userID ulid.ID, teamID ulid.ID) *TeamUser {
	return &TeamUser{
		UserID: userID,
		TeamID: teamID,
	}
}
