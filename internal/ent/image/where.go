// Code generated by ent, DO NOT EDIT.

package image

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/hossein1376/kala/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldName, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldPath, v))
}

// Caption applies equality check predicate on the "caption" field. It's identical to CaptionEQ.
func Caption(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldCaption, v))
}

// Width applies equality check predicate on the "width" field. It's identical to WidthEQ.
func Width(v int32) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldWidth, v))
}

// Height applies equality check predicate on the "height" field. It's identical to HeightEQ.
func Height(v int32) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldHeight, v))
}

// SizeKB applies equality check predicate on the "size_kb" field. It's identical to SizeKBEQ.
func SizeKB(v float64) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldSizeKB, v))
}

// UploadedAt applies equality check predicate on the "uploaded_at" field. It's identical to UploadedAtEQ.
func UploadedAt(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldEQ(FieldUploadedAt, vc))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldName, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldPath, v))
}

// CaptionEQ applies the EQ predicate on the "caption" field.
func CaptionEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldCaption, v))
}

// CaptionNEQ applies the NEQ predicate on the "caption" field.
func CaptionNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldCaption, v))
}

// CaptionIn applies the In predicate on the "caption" field.
func CaptionIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldCaption, vs...))
}

// CaptionNotIn applies the NotIn predicate on the "caption" field.
func CaptionNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldCaption, vs...))
}

// CaptionGT applies the GT predicate on the "caption" field.
func CaptionGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldCaption, v))
}

// CaptionGTE applies the GTE predicate on the "caption" field.
func CaptionGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldCaption, v))
}

// CaptionLT applies the LT predicate on the "caption" field.
func CaptionLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldCaption, v))
}

// CaptionLTE applies the LTE predicate on the "caption" field.
func CaptionLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldCaption, v))
}

// CaptionContains applies the Contains predicate on the "caption" field.
func CaptionContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldCaption, v))
}

// CaptionHasPrefix applies the HasPrefix predicate on the "caption" field.
func CaptionHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldCaption, v))
}

// CaptionHasSuffix applies the HasSuffix predicate on the "caption" field.
func CaptionHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldCaption, v))
}

// CaptionIsNil applies the IsNil predicate on the "caption" field.
func CaptionIsNil() predicate.Image {
	return predicate.Image(sql.FieldIsNull(FieldCaption))
}

// CaptionNotNil applies the NotNil predicate on the "caption" field.
func CaptionNotNil() predicate.Image {
	return predicate.Image(sql.FieldNotNull(FieldCaption))
}

// CaptionEqualFold applies the EqualFold predicate on the "caption" field.
func CaptionEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldCaption, v))
}

// CaptionContainsFold applies the ContainsFold predicate on the "caption" field.
func CaptionContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldCaption, v))
}

// WidthEQ applies the EQ predicate on the "width" field.
func WidthEQ(v int32) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldWidth, v))
}

// WidthNEQ applies the NEQ predicate on the "width" field.
func WidthNEQ(v int32) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldWidth, v))
}

// WidthIn applies the In predicate on the "width" field.
func WidthIn(vs ...int32) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldWidth, vs...))
}

// WidthNotIn applies the NotIn predicate on the "width" field.
func WidthNotIn(vs ...int32) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldWidth, vs...))
}

// WidthGT applies the GT predicate on the "width" field.
func WidthGT(v int32) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldWidth, v))
}

// WidthGTE applies the GTE predicate on the "width" field.
func WidthGTE(v int32) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldWidth, v))
}

// WidthLT applies the LT predicate on the "width" field.
func WidthLT(v int32) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldWidth, v))
}

// WidthLTE applies the LTE predicate on the "width" field.
func WidthLTE(v int32) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldWidth, v))
}

// HeightEQ applies the EQ predicate on the "height" field.
func HeightEQ(v int32) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldHeight, v))
}

// HeightNEQ applies the NEQ predicate on the "height" field.
func HeightNEQ(v int32) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldHeight, v))
}

// HeightIn applies the In predicate on the "height" field.
func HeightIn(vs ...int32) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldHeight, vs...))
}

// HeightNotIn applies the NotIn predicate on the "height" field.
func HeightNotIn(vs ...int32) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldHeight, vs...))
}

// HeightGT applies the GT predicate on the "height" field.
func HeightGT(v int32) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldHeight, v))
}

// HeightGTE applies the GTE predicate on the "height" field.
func HeightGTE(v int32) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldHeight, v))
}

// HeightLT applies the LT predicate on the "height" field.
func HeightLT(v int32) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldHeight, v))
}

// HeightLTE applies the LTE predicate on the "height" field.
func HeightLTE(v int32) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldHeight, v))
}

// SizeKBEQ applies the EQ predicate on the "size_kb" field.
func SizeKBEQ(v float64) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldSizeKB, v))
}

// SizeKBNEQ applies the NEQ predicate on the "size_kb" field.
func SizeKBNEQ(v float64) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldSizeKB, v))
}

// SizeKBIn applies the In predicate on the "size_kb" field.
func SizeKBIn(vs ...float64) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldSizeKB, vs...))
}

// SizeKBNotIn applies the NotIn predicate on the "size_kb" field.
func SizeKBNotIn(vs ...float64) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldSizeKB, vs...))
}

// SizeKBGT applies the GT predicate on the "size_kb" field.
func SizeKBGT(v float64) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldSizeKB, v))
}

// SizeKBGTE applies the GTE predicate on the "size_kb" field.
func SizeKBGTE(v float64) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldSizeKB, v))
}

// SizeKBLT applies the LT predicate on the "size_kb" field.
func SizeKBLT(v float64) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldSizeKB, v))
}

// SizeKBLTE applies the LTE predicate on the "size_kb" field.
func SizeKBLTE(v float64) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldSizeKB, v))
}

// UploadedAtEQ applies the EQ predicate on the "uploaded_at" field.
func UploadedAtEQ(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldEQ(FieldUploadedAt, vc))
}

// UploadedAtNEQ applies the NEQ predicate on the "uploaded_at" field.
func UploadedAtNEQ(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldNEQ(FieldUploadedAt, vc))
}

// UploadedAtIn applies the In predicate on the "uploaded_at" field.
func UploadedAtIn(vs ...time.Time) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.Image(sql.FieldIn(FieldUploadedAt, v...))
}

// UploadedAtNotIn applies the NotIn predicate on the "uploaded_at" field.
func UploadedAtNotIn(vs ...time.Time) predicate.Image {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = time.Time(vs[i])
	}
	return predicate.Image(sql.FieldNotIn(FieldUploadedAt, v...))
}

// UploadedAtGT applies the GT predicate on the "uploaded_at" field.
func UploadedAtGT(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldGT(FieldUploadedAt, vc))
}

// UploadedAtGTE applies the GTE predicate on the "uploaded_at" field.
func UploadedAtGTE(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldGTE(FieldUploadedAt, vc))
}

// UploadedAtLT applies the LT predicate on the "uploaded_at" field.
func UploadedAtLT(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldLT(FieldUploadedAt, vc))
}

// UploadedAtLTE applies the LTE predicate on the "uploaded_at" field.
func UploadedAtLTE(v time.Time) predicate.Image {
	vc := time.Time(v)
	return predicate.Image(sql.FieldLTE(FieldUploadedAt, vc))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UserTable, UserPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComment applies the HasEdge predicate on the "comment" edge.
func HasComment() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CommentTable, CommentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentWith applies the HasEdge predicate on the "comment" edge with a given conditions (other predicates).
func HasCommentWith(preds ...predicate.Comment) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newCommentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBrand applies the HasEdge predicate on the "brand" edge.
func HasBrand() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, BrandTable, BrandPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBrandWith applies the HasEdge predicate on the "brand" edge with a given conditions (other predicates).
func HasBrandWith(preds ...predicate.Brand) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newBrandStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, CategoryTable, CategoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSubCategory applies the HasEdge predicate on the "sub_category" edge.
func HasSubCategory() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubCategoryTable, SubCategoryPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubCategoryWith applies the HasEdge predicate on the "sub_category" edge with a given conditions (other predicates).
func HasSubCategoryWith(preds ...predicate.SubCategory) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newSubCategoryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
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
func Not(p predicate.Image) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		p(s.Not())
	})
}
