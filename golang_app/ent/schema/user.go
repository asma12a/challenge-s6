package schema

import (
	"entgo.io/ent"
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
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.JSON("role", []string{}).Default([]string{"user"}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
