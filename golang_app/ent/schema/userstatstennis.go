package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// UserStatsTennis holds the schema definition for the UserStatsTennis entity.
type UserStatsTennis struct {
	ent.Schema
}

// Fields of the UserStatsTennis.
func (UserStatsTennis) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("stats_id").NotEmpty(),
		field.Int8("sets").NonNegative().Default(0),
		field.Int8("rating").Min(0).Max(10),
	}
}

// Edges of the UserStatsTennis.
func (UserStatsTennis) Edges() []ent.Edge {
	return nil
}
