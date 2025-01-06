package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.Float("latitude").StructTag(`validate:"required"`),
		field.Float("longitude").StructTag(`validate:"required"`),
		field.String("date").NotEmpty().StructTag(`validate:"required"`),
		field.String("event_code").NotEmpty().Unique(),
		field.Bool("is_public").Default(true),
		field.Enum("event_type").Values("match", "training").Default("match").Nillable(), // Permet de ne pas demander le champ lors de la création, à condition de gérer partout le pointeur
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sport", Sport.Type).Ref("events").Unique().Required(),
		edge.To("user_stats", UserStats.Type).StorageKey(edge.Column("event_id")),
		edge.To("messages", Message.Type).StorageKey(edge.Column("event_id")),
		edge.To("teams", Team.Type).StorageKey(edge.Column("event_id")).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
