package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Mixins of the Order
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return nil
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("seller", Seller.Type).
			Ref("order").
			Unique(),
		edge.From("product", Product.Type).Ref("order"),
		edge.From("user", User.Type).
			Ref("order").
			Unique().
			Required(),
	}
}
