package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SportStatLabels holds the schema definition for the SportStatLabels entity.
type SportStatLabels struct {
	ent.Schema
}

func (SportStatLabels) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the SportStatLabels.
func (SportStatLabels) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").NotEmpty(),
		field.String("unit").Optional(),
		field.Bool("is_main").Default(false),
	}
}

// Edges of the SportStatLabels.
func (SportStatLabels) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sport", Sport.Type).Ref("sport_stat_labels").Unique().Required(),
		edge.To("user_stats", UserStats.Type).
			StorageKey(edge.Column("stat_id")),
	}
}
