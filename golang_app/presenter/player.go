package presenter

import (
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Role string

const (
	PlayerRole Role = "player"
	CoachRole  Role = "coach"
	OrgRole    Role = "org"
)

type Status string

const (
	PendingStatus Status = "pending"
	ValidStatus   Status = "valid"
)

type Player struct {
	ID     ulid.ID `json:"id,omitempty"`
	Name   string  `json:"name,omitempty"`
	Email  string  `json:"email,omitempty"`
	Role   Role    `json:"role,omitempty"`
	Status Status  `json:"status,omitempty"`
	UserID ulid.ID `json:"user_id,omitempty"`
}
