package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("name").NotEmpty(),
		field.String("address").NotEmpty(),
		field.Int16("event_code").Positive(),
		field.String("date").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Bool("is_public").Default(false),
		field.Bool("is_finished").Default(false),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event_type", EventType.Type).Ref("events").Unique(),
		edge.From("sport", Sport.Type).Ref("events").Unique(),
		edge.To("user_stats_id", UserStats.Type).StorageKey(edge.Column("event_id")),
		edge.To("basket_event_id", BasketEvent.Type).StorageKey(edge.Column("event_id")),
		edge.To("tennis_event_id", TennisEvent.Type).StorageKey(edge.Column("event_id")),
		edge.To("running_event_id", RunningEvent.Type).StorageKey(edge.Column("event_id")),
		edge.To("training_event_id", TrainingEvent.Type).StorageKey(edge.Column("event_id")),
	}
}
