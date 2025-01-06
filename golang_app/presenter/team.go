package presenter

import (
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Team représente le type de présentation pour une équipe.
type Team struct {
	ID         ulid.ID  `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	MaxPlayers int      `json:"maxPlayers,omitempty"`
	Players    []Player `json:"players,omitempty"`
}
