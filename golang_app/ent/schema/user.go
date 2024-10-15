package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/oklog/ulid/v2"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(
			func() string {
				return ulid.Make().String()
			},
		).NotEmpty().Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
		field.String("password").NotEmpty(),
		field.String("role").Default("user").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user_stats", UserStats.Type).
			StorageKey(edge.Column("user_id")),
	}
}
