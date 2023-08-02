// Code generated by ent, DO NOT EDIT.

package brand

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hossein1376/kala/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldDescription, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldStatus, v))
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldRating, v))
}

// RatingCount applies equality check predicate on the "rating_count" field. It's identical to RatingCountEQ.
func RatingCount(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldRatingCount, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Brand {
	return predicate.Brand(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Brand {
	return predicate.Brand(sql.FieldContainsFold(FieldDescription, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldStatus, v))
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldRating, v))
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldRating, v))
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...float64) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldRating, vs...))
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...float64) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldRating, vs...))
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldRating, v))
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldRating, v))
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldRating, v))
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v float64) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldRating, v))
}

// RatingCountEQ applies the EQ predicate on the "rating_count" field.
func RatingCountEQ(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldEQ(FieldRatingCount, v))
}

// RatingCountNEQ applies the NEQ predicate on the "rating_count" field.
func RatingCountNEQ(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldNEQ(FieldRatingCount, v))
}

// RatingCountIn applies the In predicate on the "rating_count" field.
func RatingCountIn(vs ...int32) predicate.Brand {
	return predicate.Brand(sql.FieldIn(FieldRatingCount, vs...))
}

// RatingCountNotIn applies the NotIn predicate on the "rating_count" field.
func RatingCountNotIn(vs ...int32) predicate.Brand {
	return predicate.Brand(sql.FieldNotIn(FieldRatingCount, vs...))
}

// RatingCountGT applies the GT predicate on the "rating_count" field.
func RatingCountGT(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldGT(FieldRatingCount, v))
}

// RatingCountGTE applies the GTE predicate on the "rating_count" field.
func RatingCountGTE(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldGTE(FieldRatingCount, v))
}

// RatingCountLT applies the LT predicate on the "rating_count" field.
func RatingCountLT(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldLT(FieldRatingCount, v))
}

// RatingCountLTE applies the LTE predicate on the "rating_count" field.
func RatingCountLTE(v int32) predicate.Brand {
	return predicate.Brand(sql.FieldLTE(FieldRatingCount, v))
}

// HasImage applies the HasEdge predicate on the "image" edge.
func HasImage() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ImageTable, ImagePrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasImageWith applies the HasEdge predicate on the "image" edge with a given conditions (other predicates).
func HasImageWith(preds ...predicate.Image) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := newImageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CategoryTable, CategoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductTable, ProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Brand) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Brand) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Brand) predicate.Brand {
	return predicate.Brand(func(s *sql.Selector) {
		p(s.Not())
	})
}
