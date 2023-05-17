package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pros holds the schema definition for the Pros entity.
type Pros struct {
	ent.Schema
}

// Fields of the Pros.
func (Pros) Fields() []ent.Field {
	return []ent.Field{
		field.String("pro").NotEmpty(),
	}
}

// Edges of the Pros.
func (Pros) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("comment", Comment.Type).
			Ref("pros").
			Unique(),
	}
}
