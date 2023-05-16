package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin holds the schema definition for the TimeMixin entity.
type TimeMixin struct {
	mixin.Schema
}

// Fields of the TimeMixin.
func (TimeMixin) Fields() []ent.Field {
	createTimeMixin := mixin.AnnotateFields(mixin.CreateTime{})
	updateTimeMixin := mixin.AnnotateFields(mixin.UpdateTime{})
	return append(createTimeMixin.Fields(), updateTimeMixin.Fields()...)
}
