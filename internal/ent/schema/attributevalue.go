package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AttributeValue holds the schema definition for the AttributeValue entity.
type AttributeValue struct {
	ent.Schema
}

// Fields of the AttributeValue.
func (AttributeValue) Fields() []ent.Field {
	return []ent.Field{
		field.String("value").NotEmpty(),
		field.Int("attribute"),
		field.Int("product"),
	}
}

// Edges of the AttributeValue.
func (AttributeValue) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("attributes", Attribute.Type).
			Ref("values").
			Field("attribute").
			Required().
			Unique(),
		edge.From("products", Product.Type).
			Ref("values").
			Field("product").
			Required().
			Unique(),
	}
}
