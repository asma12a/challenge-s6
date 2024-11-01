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

func (TeamUser) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the TeamUser.
func (TeamUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").GoType(ulid.ID("")).NotEmpty(),
		field.String("team_id").GoType(ulid.ID("")).NotEmpty(),
		field.Strings("roles").Default([]string{"player"}),
	}
}

// Edges of the TeamUser.
func (TeamUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).Ref("team_users").Unique().Required().Field("team_id"),
		edge.From("user", User.Type).Ref("team_users").Unique().Required().Field("user_id"),
	}
}
