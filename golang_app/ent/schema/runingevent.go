package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// RunningEvent holds the schema definition for the RunningEvent entity.
type RunningEvent struct {
	ent.Schema
}

// Fields of the RunningEvent.
func (RunningEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("event_running_id").NotEmpty(),
		field.String("team_id").NotEmpty(),
	}

}

// Edges of the RunningEvent.
func (RunningEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
