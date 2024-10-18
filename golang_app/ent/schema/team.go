package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teams_id", EventTeams.Type).
			StorageKey(edge.Column("team_id")),
	}
}
