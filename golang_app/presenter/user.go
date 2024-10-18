package presenter

import (

	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Event représente le type de présentation pour un événement.
type User struct {
	ID         ulid.ID   `json:"id,omitempty"`
	Name       string    `json:"name,omitempty"`
	Email    	string    `json:"address,omitempty"` 
	Roles      []string  `json:"roles,omitempty"`
    // Utiliser le type Sport personnalisé
}
