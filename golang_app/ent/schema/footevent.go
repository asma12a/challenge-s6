package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// FootEvent holds the schema definition for the FootEvent entity.
type FootEvent struct {
	ent.Schema
}

// Fields of the FootEvent.
func (FootEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("event_id").NotEmpty(),
		field.String("team_A_id").NotEmpty(),
		field.String("team_B_id").NotEmpty(),
	}

}

// Edges of the FootEvent.
func (FootEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
