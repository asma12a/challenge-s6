package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Event holds the schema definition for the EventTeams entity.
type EventTeams struct {
	ent.Schema
}

// Fields of the Event.
func (EventTeams) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("event_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("team_id").GoType(ulid.ID("")).NotEmpty(),
	}
}

// Edges of the Event.
func (EventTeams) Edges() []ent.Edge {
	return []ent.Edge{}
}
