package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// TeamUser holds the schema definition for the TeamUser entity.
type TeamUser struct {
	ent.Schema
}

// Fields of the TeamUser.
func (TeamUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
	}
}

// Edges of the TeamUser.
func (TeamUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).Ref("teamusers").Unique(),
		edge.From("user", User.Type).Ref("teamusers").Unique(),
	}
}
