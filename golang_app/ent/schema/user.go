package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().StructTag(`validate:"required"`),
		field.String("email").NotEmpty().Unique().StructTag(`validate:"required,email"`).Match(regexp.MustCompile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$")),
		field.String("password").NotEmpty(),
		field.Strings("roles").Default([]string{"user"}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_stats", UserStats.Type).
			StorageKey(edge.Column("user_id")),
		edge.To("team_users", TeamUser.Type).StorageKey(edge.Column("user_id")),
		edge.To("user_message_id", Message.Type).StorageKey(edge.Column("user_id")),
	}
}
