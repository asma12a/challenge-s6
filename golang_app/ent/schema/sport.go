package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

// Sport holds the schema definition for the Sport entity.
type Sport struct {
	ent.Schema
}

// Fields of the Sport.
func (Sport) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(
			func() string {
				return ulid.Make().String()
			},
		).NotEmpty().Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.String("image_url").Optional(),
	}
}

// Edges of the Sport.
func (Sport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("event", Event.Type),
	}
}
