package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubCategory holds the schema definition for the SubCategory entity.
type SubCategory struct {
	ent.Schema
}

// Mixins of the SubCategory
func (SubCategory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the SubCategory.
func (SubCategory) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").NotEmpty(),
	}
}

// Edges of the SubCategory.
func (SubCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("sub_category"),
		edge.From("category", Category.Type).
			Ref("sub_category").
			Unique(),
	}
}
