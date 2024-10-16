package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// UserStatsBasket holds the schema definition for the UserStatsBasket entity.
type UserStatsBasket struct {
	ent.Schema
}

// Fields of the UserStatsBasket.
func (UserStatsBasket) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("stats_id").NotEmpty(),
		field.Int8("hoops").NonNegative().Default(0),
		field.Int8("assists").NonNegative().Default(0),
		field.Int8("rating").Min(0).Max(10),
	}
}

// Edges of the UserStatsBasket.
func (UserStatsBasket) Edges() []ent.Edge {
	return nil
}
