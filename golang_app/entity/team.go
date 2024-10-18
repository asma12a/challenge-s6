package entity

import (
	"github.com/asma12a/challenge-s6/ent"
)

type Team struct {
	ent.Team
}

func NewTeam(name string, maxPlayers int) *Team {
	return &Team{
		Team: ent.Team{
			Name:       name,
			MaxPlayers: maxPlayers,
		},
	}
}
