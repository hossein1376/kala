package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Seller holds the schema definition for the Seller entity.
type Seller struct {
	ent.Schema
}

// Mixins of the Seller
func (Seller) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Seller.
func (Seller) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
		field.Float("rating").Range(0, 5),
		field.Int32("rating_count").Positive(),
		field.String("phone").NotEmpty(),
	}
}

// Edges of the Seller.
func (Seller) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("product", Product.Type),
		edge.To("category", Category.Type),
		edge.To("address", Address.Type),

		edge.From("user", User.Type).
			Ref("seller").
			Unique(),
	}
}
