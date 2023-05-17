package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Logs holds the schema definition for the Logs entity.
type Logs struct {
	ent.Schema
}

// Fields of the Logs.
func (Logs) Fields() []ent.Field {
	return []ent.Field{
		field.String("action").Optional().Nillable(),
		field.String("IP"),
		field.String("agent"),
		field.Time("date").GoType(time.Now()),
	}
}

// Edges of the Logs.
func (Logs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("logs").
			Unique(),
	}
}
