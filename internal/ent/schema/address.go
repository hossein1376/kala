package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").NotEmpty(),
		field.String("city").NotEmpty(),
		field.String("state").NotEmpty(),
		field.String("first_name").Optional().Nillable(),
		field.String("last_name").Optional().Nillable(),
		field.String("phone").NotEmpty(),
		field.String("zip_code").NotEmpty(),
		field.String("coordinates").NotEmpty(),
		field.Bool("is_seller").Default(false),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("address").
			Unique(),
		edge.From("seller", Seller.Type).
			Ref("address").
			Unique(),
	}
}
