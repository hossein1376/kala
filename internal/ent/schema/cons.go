package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Cons holds the schema definition for the Cons entity.
type Cons struct {
	ent.Schema
}

// Fields of the Cons.
func (Cons) Fields() []ent.Field {
	return []ent.Field{
		field.String("con").NotEmpty(),
	}
}

// Edges of the Cons.
func (Cons) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", Comment.Type).
			Ref("cons").
			Unique(),
	}
}
