package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

type EventType struct {
	ent.Schema
}

func (EventType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(
			func() string {
				return ulid.Make().String()
			},
		).NotEmpty().Unique().Immutable(),
		field.String("name").NotEmpty(),
	}
}

func (EventType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).StorageKey(edge.Column("event_type_id")),
	}
}
