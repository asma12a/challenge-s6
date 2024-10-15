package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// UserStats holds the schema definition for the UserStats entity.
type UserStats struct {
	ent.Schema
}

// Fields of the UserStats.
func (UserStats) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", ulid.ULID{}),
		field.String("user_id"),
		field.UUID("event_id", uuid.UUID{}),
	}

}

// Edges of the UserStats.
func (UserStats) Edges() []ent.Edge {
	return []ent.Edge{}
}
