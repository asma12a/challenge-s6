package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// EventType holds the schema definition for the EventType entity.
type EventType struct {
	ent.Schema
}

// Fields of the EventType.
func (EventType) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("name").NotEmpty(),
	}
}

// Edges of the EventType.
func (EventType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).StorageKey(edge.Column("event_type_id")),
	}
}
