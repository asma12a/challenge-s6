package presenter

import (

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Event représente le type de présentation pour un événement.
type Team struct {
	ID         ulid.ID   `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	MaxPlayers int       `json:"maxPlayers,omitempty"`
	// Utiliser le type Sport personnalisé
}
