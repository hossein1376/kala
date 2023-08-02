package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("path").NotEmpty(),
		field.String("caption").Optional(),
		field.Int32("width").Positive(),
		field.Int32("height").Positive(),
		field.Float("size_kb").Positive(),
		field.Time("uploaded_at").GoType(time.Now()),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("image"),
		edge.From("comment", Comment.Type).
			Ref("image"),
		edge.From("brand", Brand.Type).
			Ref("image"),
		edge.From("product", Product.Type).
			Ref("image"),
		edge.From("category", Category.Type).
			Ref("image"),
		edge.From("sub_category", SubCategory.Type).
			Ref("image"),
	}
}
