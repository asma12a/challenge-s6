package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

func (Message) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("user_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("content").NotEmpty(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
