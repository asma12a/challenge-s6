package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ActionLog struct {
	ent.Schema
}

func (ActionLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("action").NotEmpty().StructTag(`validate:"required"`),

		field.String("description").NotEmpty().StructTag(`validate:"required"`),

		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

func (ActionLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("action_logs").Unique().Required(),
	}
}

func (ActionLog) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("action").Unique(),
	}
}
