// Code generated by ent, DO NOT EDIT.

package product

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hossein1376/kala/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdateTime, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// Review applies equality check predicate on the "review" field. It's identical to ReviewEQ.
func Review(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldReview, v))
}

// Rating applies equality check predicate on the "rating" field. It's identical to RatingEQ.
func Rating(v float64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldRating, v))
}

// RatingCount applies equality check predicate on the "rating_count" field. It's identical to RatingCountEQ.
func RatingCount(v int32) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldRatingCount, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v int32) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int32) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldQuantity, v))
}

// Available applies equality check predicate on the "available" field. It's identical to AvailableEQ.
func Available(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldAvailable, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStatus, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldUpdateTime, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldDescription, v))
}

// ReviewEQ applies the EQ predicate on the "review" field.
func ReviewEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldReview, v))
}

// ReviewNEQ applies the NEQ predicate on the "review" field.
func ReviewNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldReview, v))
}

// ReviewIn applies the In predicate on the "review" field.
func ReviewIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldReview, vs...))
}

// ReviewNotIn applies the NotIn predicate on the "review" field.
func ReviewNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldReview, vs...))
}

// ReviewGT applies the GT predicate on the "review" field.
func ReviewGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldReview, v))
}

// ReviewGTE applies the GTE predicate on the "review" field.
func ReviewGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldReview, v))
}

// ReviewLT applies the LT predicate on the "review" field.
func ReviewLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldReview, v))
}

// ReviewLTE applies the LTE predicate on the "review" field.
func ReviewLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldReview, v))
}

// ReviewContains applies the Contains predicate on the "review" field.
func ReviewContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldReview, v))
}

// ReviewHasPrefix applies the HasPrefix predicate on the "review" field.
func ReviewHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldReview, v))
}

// ReviewHasSuffix applies the HasSuffix predicate on the "review" field.
func ReviewHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldReview, v))
}

// ReviewEqualFold applies the EqualFold predicate on the "review" field.
func ReviewEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldReview, v))
}

// ReviewContainsFold applies the ContainsFold predicate on the "review" field.
func ReviewContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldReview, v))
}

// RatingEQ applies the EQ predicate on the "rating" field.
func RatingEQ(v float64) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldRating, v))
}

// RatingNEQ applies the NEQ predicate on the "rating" field.
func RatingNEQ(v float64) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldRating, v))
}

// RatingIn applies the In predicate on the "rating" field.
func RatingIn(vs ...float64) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldRating, vs...))
}

// RatingNotIn applies the NotIn predicate on the "rating" field.
func RatingNotIn(vs ...float64) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldRating, vs...))
}

// RatingGT applies the GT predicate on the "rating" field.
func RatingGT(v float64) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldRating, v))
}

// RatingGTE applies the GTE predicate on the "rating" field.
func RatingGTE(v float64) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldRating, v))
}

// RatingLT applies the LT predicate on the "rating" field.
func RatingLT(v float64) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldRating, v))
}

// RatingLTE applies the LTE predicate on the "rating" field.
func RatingLTE(v float64) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldRating, v))
}

// RatingCountEQ applies the EQ predicate on the "rating_count" field.
func RatingCountEQ(v int32) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldRatingCount, v))
}

// RatingCountNEQ applies the NEQ predicate on the "rating_count" field.
func RatingCountNEQ(v int32) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldRatingCount, v))
}

// RatingCountIn applies the In predicate on the "rating_count" field.
func RatingCountIn(vs ...int32) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldRatingCount, vs...))
}

// RatingCountNotIn applies the NotIn predicate on the "rating_count" field.
func RatingCountNotIn(vs ...int32) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldRatingCount, vs...))
}

// RatingCountGT applies the GT predicate on the "rating_count" field.
func RatingCountGT(v int32) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldRatingCount, v))
}

// RatingCountGTE applies the GTE predicate on the "rating_count" field.
func RatingCountGTE(v int32) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldRatingCount, v))
}

// RatingCountLT applies the LT predicate on the "rating_count" field.
func RatingCountLT(v int32) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldRatingCount, v))
}

// RatingCountLTE applies the LTE predicate on the "rating_count" field.
func RatingCountLTE(v int32) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldRatingCount, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v int32) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v int32) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...int32) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...int32) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v int32) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v int32) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v int32) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v int32) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPrice, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int32) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int32) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int32) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int32) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int32) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int32) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int32) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int32) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldQuantity, v))
}

// AvailableEQ applies the EQ predicate on the "available" field.
func AvailableEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldAvailable, v))
}

// AvailableNEQ applies the NEQ predicate on the "available" field.
func AvailableNEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldAvailable, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldStatus, v))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OrderTable, OrderPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}
