package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("password").NotEmpty(),
		field.JSON("role", []string{}).Default([]string{"user"})}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_stats", UserStats.Type).
			StorageKey(edge.Column("user_id")),
		edge.To("teamusers", TeamUser.Type).StorageKey(edge.Column("user_id")),
	}
}
