package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// BasketEvent holds the schema definition for the BasketEvent entity.
type BasketEvent struct {
	ent.Schema
}

// Fields of the BasketEvent.
func (BasketEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("event_basket_id").NotEmpty(),
		field.String("team_A_id").NotEmpty(),
		field.String("team_B_id").NotEmpty(),
	}

}

// Edges of the BasketEvent.
func (BasketEvent) Edges() []ent.Edge {
	return []ent.Edge{}
}
