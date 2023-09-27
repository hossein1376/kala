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
		field.Enum("role").
			Values("admin",
				"seller",
				"user").Default("user"),
		field.Bool("status").Default(true),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type).
			StorageKey(edge.Column("user_id")),
		edge.To("logs", Logs.Type).
			StorageKey(edge.Column("user")),
	}
}
