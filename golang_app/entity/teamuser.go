package entity

import (
	"github.com/asma12a/challenge-s6/ent"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
	"github.com/asma12a/challenge-s6/ent/teamuser"
)

type TeamUser struct {
	ent.TeamUser
	UserID *ulid.ID `json:"user_id,omitempty"`
	TeamID ulid.ID  `json:"team_id"`
}

func NewTeamUser(
	email string, role teamuser.Role, status teamuser.Status,
	userId *ulid.ID, teamId ulid.ID,
) *TeamUser {
	return &TeamUser{
		TeamUser: ent.TeamUser{
			Email:  email,
			Role:   role,
			Status: status,
		},
		UserID: userId,
		TeamID: teamId,
	}
}
