package schema

import "C"
import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Mixins of the Category
func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sub_category", SubCategory.Type).
			StorageKey(edge.Column("category")),

		edge.From("product", Product.Type).Ref("category"),
		edge.From("brand", Brand.Type).Ref("category"),
		edge.From("seller", Seller.Type).Ref("category"),
	}
}
