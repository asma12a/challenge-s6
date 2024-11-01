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

func (EventTeams) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Event.
func (EventTeams) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("team_id").GoType(ulid.ID("")).NotEmpty(),
	}
}

// Edges of the Event.
func (EventTeams) Edges() []ent.Edge {
	return nil
}
