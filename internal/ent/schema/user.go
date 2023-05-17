package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixins of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
		field.Bytes("password").Sensitive(),
		field.String("first_name").Optional(),
		field.String("last_name").Optional(),
		field.String("email").Optional(),
		field.String("phone").Optional(),
		field.Bool("is_seller").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comment", Comment.Type).
			StorageKey(edge.Table("user_comments"),
				edge.Columns("user", "comment"),
				edge.Symbols("user_id", "comment_id")),
		edge.To("image", Image.Type),
		edge.To("seller", Seller.Type).
			StorageKey(edge.Column("user_id")),
		edge.To("order", Order.Type),
		edge.To("logs", Logs.Type).
			StorageKey(edge.Column("user")),
		edge.To("address", Address.Type).
			StorageKey(edge.Column("user")),
	}
}
