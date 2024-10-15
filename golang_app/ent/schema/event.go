package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(
			func() string {
				return ulid.Make().String()
			},
		).NotEmpty().Unique().Immutable(),
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

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event_type", EventType.Type).
			Ref("event").
			Unique(),
		edge.From("sport", Sport.Type).
			Ref("event").
			Unique(),
	}
}
