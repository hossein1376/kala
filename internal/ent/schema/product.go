package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Mixins of the Product
func (Product) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Text("description").NotEmpty(),
		field.Text("review"),
		field.Float("rating").Range(0, 5),
		field.Int32("rating_count").Positive(),
		field.Int32("price").Positive(),
		field.Int32("quantity").Positive(),
		field.Bool("available"),
		field.Bool("status"),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("values", AttributeValue.Type),
		edge.To("comment", Comment.Type).
			StorageKey(edge.Table("product_comments")),
		edge.To("image", Image.Type),
		edge.To("order", Order.Type),

		edge.From("brand", Brand.Type).
			Ref("product").
			Unique(),
	}
}
