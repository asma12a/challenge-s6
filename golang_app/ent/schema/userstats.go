package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// UserStats holds the schema definition for the UserStats entity.
type UserStats struct {
	ent.Schema
}

func (UserStats) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the UserStats.
func (UserStats) Fields() []ent.Field {
	return []ent.Field{
		field.Int("stat_value").Positive().Default(0),
	}
}

// Edges of the UserStats.
func (UserStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("user_stats").Field("user_id").Unique().Required(),
		edge.From("event", Event.Type).Ref("user_stats").Field("event_id").Unique().Required(),
		edge.From("stat", SportStatLabels.Type).Ref("user_stats").Field("stat_id").Unique().Required(),
	}
}

// Indexes of the UserStats.
func (UserStats) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("user", "event", "stat").Unique(),
	}
}
