package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// SportStatLabels holds the schema definition for the SportStatLabels entity.
type SportStatLabels struct {
	ent.Schema
}

// Fields of the SportStatLabels.
func (SportStatLabels) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("sport_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("stat_label").NotEmpty(),
	}
}

// Edges of the SportStatLabels.
func (SportStatLabels) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_stats", UserStats.Type).
			StorageKey(edge.Column("stat_id")),
	}
}
