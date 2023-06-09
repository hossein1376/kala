// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/brand"
	"kala/internal/ent/category"
	"kala/internal/ent/image"
	"kala/internal/ent/predicate"
	"kala/internal/ent/product"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BrandUpdate is the builder for updating Brand entities.
type BrandUpdate struct {
	config
	hooks    []Hook
	mutation *BrandMutation
}

// Where appends a list predicates to the BrandUpdate builder.
func (bu *BrandUpdate) Where(ps ...predicate.Brand) *BrandUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdateTime sets the "update_time" field.
func (bu *BrandUpdate) SetUpdateTime(t time.Time) *BrandUpdate {
	bu.mutation.SetUpdateTime(t)
	return bu
}

// SetName sets the "name" field.
func (bu *BrandUpdate) SetName(s string) *BrandUpdate {
	bu.mutation.SetName(s)
	return bu
}

// SetDescription sets the "description" field.
func (bu *BrandUpdate) SetDescription(s string) *BrandUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BrandUpdate) SetStatus(b bool) *BrandUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// SetRating sets the "rating" field.
func (bu *BrandUpdate) SetRating(f float64) *BrandUpdate {
	bu.mutation.ResetRating()
	bu.mutation.SetRating(f)
	return bu
}

// AddRating adds f to the "rating" field.
func (bu *BrandUpdate) AddRating(f float64) *BrandUpdate {
	bu.mutation.AddRating(f)
	return bu
}

// SetRatingCount sets the "rating_count" field.
func (bu *BrandUpdate) SetRatingCount(i int32) *BrandUpdate {
	bu.mutation.ResetRatingCount()
	bu.mutation.SetRatingCount(i)
	return bu
}

// AddRatingCount adds i to the "rating_count" field.
func (bu *BrandUpdate) AddRatingCount(i int32) *BrandUpdate {
	bu.mutation.AddRatingCount(i)
	return bu
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (bu *BrandUpdate) AddImageIDs(ids ...int) *BrandUpdate {
	bu.mutation.AddImageIDs(ids...)
	return bu
}

// AddImage adds the "image" edges to the Image entity.
func (bu *BrandUpdate) AddImage(i ...*Image) *BrandUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.AddImageIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (bu *BrandUpdate) AddCategoryIDs(ids ...int) *BrandUpdate {
	bu.mutation.AddCategoryIDs(ids...)
	return bu
}

// AddCategory adds the "category" edges to the Category entity.
func (bu *BrandUpdate) AddCategory(c ...*Category) *BrandUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.AddCategoryIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (bu *BrandUpdate) AddProductIDs(ids ...int) *BrandUpdate {
	bu.mutation.AddProductIDs(ids...)
	return bu
}

// AddProduct adds the "product" edges to the Product entity.
func (bu *BrandUpdate) AddProduct(p ...*Product) *BrandUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (bu *BrandUpdate) Mutation() *BrandMutation {
	return bu.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (bu *BrandUpdate) ClearImage() *BrandUpdate {
	bu.mutation.ClearImage()
	return bu
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (bu *BrandUpdate) RemoveImageIDs(ids ...int) *BrandUpdate {
	bu.mutation.RemoveImageIDs(ids...)
	return bu
}

// RemoveImage removes "image" edges to Image entities.
func (bu *BrandUpdate) RemoveImage(i ...*Image) *BrandUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return bu.RemoveImageIDs(ids...)
}

// ClearCategory clears all "category" edges to the Category entity.
func (bu *BrandUpdate) ClearCategory() *BrandUpdate {
	bu.mutation.ClearCategory()
	return bu
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (bu *BrandUpdate) RemoveCategoryIDs(ids ...int) *BrandUpdate {
	bu.mutation.RemoveCategoryIDs(ids...)
	return bu
}

// RemoveCategory removes "category" edges to Category entities.
func (bu *BrandUpdate) RemoveCategory(c ...*Category) *BrandUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return bu.RemoveCategoryIDs(ids...)
}

// ClearProduct clears all "product" edges to the Product entity.
func (bu *BrandUpdate) ClearProduct() *BrandUpdate {
	bu.mutation.ClearProduct()
	return bu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (bu *BrandUpdate) RemoveProductIDs(ids ...int) *BrandUpdate {
	bu.mutation.RemoveProductIDs(ids...)
	return bu
}

// RemoveProduct removes "product" edges to Product entities.
func (bu *BrandUpdate) RemoveProduct(p ...*Product) *BrandUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BrandUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BrandUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BrandUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BrandUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BrandUpdate) defaults() {
	if _, ok := bu.mutation.UpdateTime(); !ok {
		v := brand.UpdateDefaultUpdateTime()
		bu.mutation.SetUpdateTime(v)
	}
}

func (bu *BrandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(brand.Table, brand.Columns, sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdateTime(); ok {
		_spec.SetField(brand.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Name(); ok {
		_spec.SetField(brand.FieldName, field.TypeString, value)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(brand.FieldDescription, field.TypeString, value)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(brand.FieldStatus, field.TypeBool, value)
	}
	if value, ok := bu.mutation.Rating(); ok {
		_spec.SetField(brand.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.AddedRating(); ok {
		_spec.AddField(brand.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := bu.mutation.RatingCount(); ok {
		_spec.SetField(brand.FieldRatingCount, field.TypeInt32, value)
	}
	if value, ok := bu.mutation.AddedRatingCount(); ok {
		_spec.AddField(brand.FieldRatingCount, field.TypeInt32, value)
	}
	if bu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.ImageTable,
			Columns: brand.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedImageIDs(); len(nodes) > 0 && !bu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.ImageTable,
			Columns: brand.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.ImageTable,
			Columns: brand.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.CategoryTable,
			Columns: brand.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !bu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.CategoryTable,
			Columns: brand.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.CategoryTable,
			Columns: brand.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedProductIDs(); len(nodes) > 0 && !bu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BrandUpdateOne is the builder for updating a single Brand entity.
type BrandUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BrandMutation
}

// SetUpdateTime sets the "update_time" field.
func (buo *BrandUpdateOne) SetUpdateTime(t time.Time) *BrandUpdateOne {
	buo.mutation.SetUpdateTime(t)
	return buo
}

// SetName sets the "name" field.
func (buo *BrandUpdateOne) SetName(s string) *BrandUpdateOne {
	buo.mutation.SetName(s)
	return buo
}

// SetDescription sets the "description" field.
func (buo *BrandUpdateOne) SetDescription(s string) *BrandUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BrandUpdateOne) SetStatus(b bool) *BrandUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// SetRating sets the "rating" field.
func (buo *BrandUpdateOne) SetRating(f float64) *BrandUpdateOne {
	buo.mutation.ResetRating()
	buo.mutation.SetRating(f)
	return buo
}

// AddRating adds f to the "rating" field.
func (buo *BrandUpdateOne) AddRating(f float64) *BrandUpdateOne {
	buo.mutation.AddRating(f)
	return buo
}

// SetRatingCount sets the "rating_count" field.
func (buo *BrandUpdateOne) SetRatingCount(i int32) *BrandUpdateOne {
	buo.mutation.ResetRatingCount()
	buo.mutation.SetRatingCount(i)
	return buo
}

// AddRatingCount adds i to the "rating_count" field.
func (buo *BrandUpdateOne) AddRatingCount(i int32) *BrandUpdateOne {
	buo.mutation.AddRatingCount(i)
	return buo
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (buo *BrandUpdateOne) AddImageIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.AddImageIDs(ids...)
	return buo
}

// AddImage adds the "image" edges to the Image entity.
func (buo *BrandUpdateOne) AddImage(i ...*Image) *BrandUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.AddImageIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (buo *BrandUpdateOne) AddCategoryIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.AddCategoryIDs(ids...)
	return buo
}

// AddCategory adds the "category" edges to the Category entity.
func (buo *BrandUpdateOne) AddCategory(c ...*Category) *BrandUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.AddCategoryIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (buo *BrandUpdateOne) AddProductIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.AddProductIDs(ids...)
	return buo
}

// AddProduct adds the "product" edges to the Product entity.
func (buo *BrandUpdateOne) AddProduct(p ...*Product) *BrandUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddProductIDs(ids...)
}

// Mutation returns the BrandMutation object of the builder.
func (buo *BrandUpdateOne) Mutation() *BrandMutation {
	return buo.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (buo *BrandUpdateOne) ClearImage() *BrandUpdateOne {
	buo.mutation.ClearImage()
	return buo
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (buo *BrandUpdateOne) RemoveImageIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.RemoveImageIDs(ids...)
	return buo
}

// RemoveImage removes "image" edges to Image entities.
func (buo *BrandUpdateOne) RemoveImage(i ...*Image) *BrandUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return buo.RemoveImageIDs(ids...)
}

// ClearCategory clears all "category" edges to the Category entity.
func (buo *BrandUpdateOne) ClearCategory() *BrandUpdateOne {
	buo.mutation.ClearCategory()
	return buo
}

// RemoveCategoryIDs removes the "category" edge to Category entities by IDs.
func (buo *BrandUpdateOne) RemoveCategoryIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.RemoveCategoryIDs(ids...)
	return buo
}

// RemoveCategory removes "category" edges to Category entities.
func (buo *BrandUpdateOne) RemoveCategory(c ...*Category) *BrandUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return buo.RemoveCategoryIDs(ids...)
}

// ClearProduct clears all "product" edges to the Product entity.
func (buo *BrandUpdateOne) ClearProduct() *BrandUpdateOne {
	buo.mutation.ClearProduct()
	return buo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (buo *BrandUpdateOne) RemoveProductIDs(ids ...int) *BrandUpdateOne {
	buo.mutation.RemoveProductIDs(ids...)
	return buo
}

// RemoveProduct removes "product" edges to Product entities.
func (buo *BrandUpdateOne) RemoveProduct(p ...*Product) *BrandUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the BrandUpdate builder.
func (buo *BrandUpdateOne) Where(ps ...predicate.Brand) *BrandUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BrandUpdateOne) Select(field string, fields ...string) *BrandUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Brand entity.
func (buo *BrandUpdateOne) Save(ctx context.Context) (*Brand, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BrandUpdateOne) SaveX(ctx context.Context) *Brand {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BrandUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BrandUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BrandUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdateTime(); !ok {
		v := brand.UpdateDefaultUpdateTime()
		buo.mutation.SetUpdateTime(v)
	}
}

func (buo *BrandUpdateOne) sqlSave(ctx context.Context) (_node *Brand, err error) {
	_spec := sqlgraph.NewUpdateSpec(brand.Table, brand.Columns, sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Brand.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, brand.FieldID)
		for _, f := range fields {
			if !brand.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != brand.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdateTime(); ok {
		_spec.SetField(brand.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Name(); ok {
		_spec.SetField(brand.FieldName, field.TypeString, value)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(brand.FieldDescription, field.TypeString, value)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(brand.FieldStatus, field.TypeBool, value)
	}
	if value, ok := buo.mutation.Rating(); ok {
		_spec.SetField(brand.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.AddedRating(); ok {
		_spec.AddField(brand.FieldRating, field.TypeFloat64, value)
	}
	if value, ok := buo.mutation.RatingCount(); ok {
		_spec.SetField(brand.FieldRatingCount, field.TypeInt32, value)
	}
	if value, ok := buo.mutation.AddedRatingCount(); ok {
		_spec.AddField(brand.FieldRatingCount, field.TypeInt32, value)
	}
	if buo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.ImageTable,
			Columns: brand.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedImageIDs(); len(nodes) > 0 && !buo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.ImageTable,
			Columns: brand.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.ImageTable,
			Columns: brand.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.CategoryTable,
			Columns: brand.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !buo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.CategoryTable,
			Columns: brand.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   brand.CategoryTable,
			Columns: brand.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedProductIDs(); len(nodes) > 0 && !buo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   brand.ProductTable,
			Columns: []string{brand.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Brand{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{brand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
