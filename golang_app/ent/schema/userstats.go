package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
		field.String("user_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("event_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("stat_id").GoType(ulid.ID("")).NotEmpty(),
		field.Int("stat_value").Positive().Default(0),
	}
}

// Edges of the UserStats.
func (UserStats) Edges() []ent.Edge {
	return nil
}
