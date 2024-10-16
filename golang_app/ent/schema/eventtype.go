package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

// EventType holds the schema definition for the EventType entity.
type EventType struct {
	ent.Schema
}

// Fields of the EventType.
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

// Edges of the EventType.
func (EventType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type), // Un EventType peut avoir plusieurs événements
	}
}
