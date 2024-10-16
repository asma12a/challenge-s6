package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/asma12a/challenge-s6/ent/schema/ulid"
)

// UserStatsRunning holds the schema definition for the UserStatsRunning entity.
type UserStatsRunning struct {
	ent.Schema
}

// Fields of the UserStatsRunning.
func (UserStatsRunning) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(ulid.ID("")).
			DefaultFunc(
				func() ulid.ID {
					return ulid.MustNew("")
				},
			),
		field.String("stats_id").NotEmpty(),
		field.Int8("distance").NonNegative().Default(0), // in meters
		field.Int8("time").NonNegative().Default(0),     // in seconds
	}
}

// Edges of the UserStatsRunning.
func (UserStatsRunning) Edges() []ent.Edge {
	return nil
}
