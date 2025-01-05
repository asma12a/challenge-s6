package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("email").NotEmpty().StructTag(`validate:"required,email"`).Match(regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")),
		field.Enum("role").Values("player", "coach", "org").Default("player"),
		field.Enum("status").Values("pending", "valid").Default("pending"),
	}
}

// Edges of the TeamUser.
func (TeamUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("team_users").Unique(),
		edge.From("team", Team.Type).Ref("team_users").Unique().Required(),
	}
}

// Indexes of the TeamUser.
func (TeamUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").Edges("team").Unique(),
	}
}
