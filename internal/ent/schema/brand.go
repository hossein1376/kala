package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Brand holds the schema definition for the Brand entity.
type Brand struct {
	ent.Schema
}

// Mixins of the Brand
func (Brand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Brand.
func (Brand) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Text("description"),
		field.Bool("status"),
		field.Float("rating"),
		field.Int32("rating_count"),
	}
}

// Edges of the Brand.
func (Brand) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type),
		edge.To("category", Category.Type),
		edge.To("product", Product.Type),
	}
}
