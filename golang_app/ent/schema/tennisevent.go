package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// TennisEvent holds the schema definition for the TennisEvent entity.
type TennisEvent struct {
	ent.Schema
}

// Fields of the TennisEvent.
func (TennisEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("event_tennis_id").NotEmpty(),
		field.String("team_A").NotEmpty(),
		field.String("team_B").NotEmpty(),
	}

}

// Edges of the TennisEvent.
func (TennisEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
