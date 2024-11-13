package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.String("email").Unique().Nillable().Optional(),
		field.Enum("role").Values("player", "coach").Default("player"),
		field.String("status").Default("pending"),
	}
}

// Edges of the TeamUser.
func (TeamUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).Ref("team_users").Unique().Required(),
		edge.From("user", User.Type).Ref("team_users").Unique(),
	}
}
