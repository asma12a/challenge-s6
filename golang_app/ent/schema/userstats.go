package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// UserStats holds the schema definition for the UserStats entity.
type UserStats struct {
	ent.Schema
}

// Fields of the UserStats.
func (UserStats) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("user_id").NotEmpty(),
		field.String("event_id").NotEmpty(),
	}

}

// Edges of the UserStats.
func (UserStats) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_stats_foot", UserStatsFoot.Type).
			StorageKey(edge.Column("stats_id")),
		edge.To("user_stats_basket", UserStatsBasket.Type).
			StorageKey(edge.Column("stats_id")),
		edge.To("user_stats_running", UserStatsRunning.Type).
			StorageKey(edge.Column("stats_id")),
		edge.To("user_stats_tennis", UserStatsTennis.Type).
			StorageKey(edge.Column("stats_id")),
	}
}
