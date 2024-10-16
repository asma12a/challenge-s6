package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

type Sport struct {
	ent.Schema
}

func (Sport) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("name").NotEmpty(),
		field.String("image_url").Optional(),
	}
}

func (Sport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).StorageKey(edge.Column("sport_id")),
	}
}
