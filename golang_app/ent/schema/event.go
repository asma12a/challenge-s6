package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

func (Event) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().StructTag(`validate:"required"`),
		field.String("address").NotEmpty().StructTag(`validate:"required"`),
		field.String("event_code").StructTag(`validate:"required"`),
		field.String("date").NotEmpty().StructTag(`validate:"required"`),
		field.Bool("is_public").Default(false),
		field.Bool("is_finished").Default(false),
		field.Enum("event_type").Values("match", "training").Default("match").Nillable(), // Permet de ne pas demander le champ lors de la création, à condition de gérer partout le pointeur
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sport", Sport.Type).Ref("events").Unique(),
		edge.To("user_stats", UserStats.Type).StorageKey(edge.Column("event_id")),
		edge.To("event_teams", EventTeams.Type).StorageKey(edge.Column("event_id")),
		edge.To("messages", Message.Type).StorageKey(edge.Column("event_id")),
	}
}
