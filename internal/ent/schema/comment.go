package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Mixins of the Comment
func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").NotEmpty(),
		field.Enum("status").
			Values("published",
				"draft",
				"deleted").Default("draft"),
		field.Int32("likes").Default(0).Min(0),
		field.Int32("dislikes").Default(0).Min(0),
		field.Float("rating").Default(0).Range(0, 5),
		field.Int32("rating_count").Positive(),
		field.Bool("verified_buyer"),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type).
			StorageKey(edge.Column("image")).
			Unique(),
		edge.To("cons", Cons.Type).
			StorageKey(edge.Column("comment")),
		edge.To("pros", Pros.Type).
			StorageKey(edge.Column("comment")),

		edge.From("user", User.Type).
			Ref("comment"),
		edge.From("product", Product.Type).
			Ref("comment"),
	}
}
