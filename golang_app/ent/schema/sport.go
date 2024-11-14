package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Sport struct {
	ent.Schema
}

func (Sport) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

func (Sport) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().StructTag(`validate:"required"`),
		field.String("color").Optional(),
		field.String("image_url").Optional(),
		field.Int("max_teams").Optional(),
		field.Enum("type").Values("individual", "team").Default("team").Optional(),
	}
}

func (Sport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).
			StorageKey(edge.Column("sport_id")),
		edge.To("sport_stat_labels", SportStatLabels.Type).
			StorageKey(edge.Column("stat_id")),
	}
}
