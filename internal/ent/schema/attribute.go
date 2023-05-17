package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

// Fields of the Attribute.
func (Attribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
	}
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("values", AttributeValue.Type),
	}
}
