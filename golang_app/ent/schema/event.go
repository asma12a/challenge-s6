package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Int16("event_code").Positive(),
		field.String("date").NotEmpty(),
		field.Time("created_at").
			Default(time.Now),
		field.Bool("is_public").Default(false),
		field.Bool("is_finished").Default(false),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_stats_id", UserStats.Type).StorageKey(edge.Column("event_id")),
	}
}
