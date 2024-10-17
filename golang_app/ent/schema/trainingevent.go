package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// TrainingEvent holds the schema definition for the TrainingEvent entity.
type TrainingEvent struct {
	ent.Schema
}

// Fields of the TrainingEvent.
func (TrainingEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("event_id").NotEmpty(),
		field.String("team_id").NotEmpty(),
	}

}

// Edges of the TrainingEvent.
func (TrainingEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
