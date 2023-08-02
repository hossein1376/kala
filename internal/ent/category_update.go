// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hossein1376/kala/internal/ent/brand"
	"github.com/hossein1376/kala/internal/ent/category"
	"github.com/hossein1376/kala/internal/ent/image"
	"github.com/hossein1376/kala/internal/ent/predicate"
	"github.com/hossein1376/kala/internal/ent/product"
)

// CategoryUpdate is the builder for updating Category entities.
type CategoryUpdate struct {
	config
	hooks    []Hook
	mutation *CategoryMutation
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cu *CategoryUpdate) Where(ps ...predicate.Category) *CategoryUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *CategoryUpdate) SetUpdateTime(t time.Time) *CategoryUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetName sets the "name" field.
func (cu *CategoryUpdate) SetName(s string) *CategoryUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *CategoryUpdate) SetDescription(s string) *CategoryUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (cu *CategoryUpdate) AddImageIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddImageIDs(ids...)
	return cu
}

// AddImage adds the "image" edges to the Image entity.
func (cu *CategoryUpdate) AddImage(i ...*Image) *CategoryUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.AddImageIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (cu *CategoryUpdate) AddProductIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddProductIDs(ids...)
	return cu
}

// AddProduct adds the "product" edges to the Product entity.
func (cu *CategoryUpdate) AddProduct(p ...*Product) *CategoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProductIDs(ids...)
}

// AddBrandIDs adds the "brand" edge to the Brand entity by IDs.
func (cu *CategoryUpdate) AddBrandIDs(ids ...int) *CategoryUpdate {
	cu.mutation.AddBrandIDs(ids...)
	return cu
}

// AddBrand adds the "brand" edges to the Brand entity.
func (cu *CategoryUpdate) AddBrand(b ...*Brand) *CategoryUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.AddBrandIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cu *CategoryUpdate) Mutation() *CategoryMutation {
	return cu.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (cu *CategoryUpdate) ClearImage() *CategoryUpdate {
	cu.mutation.ClearImage()
	return cu
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (cu *CategoryUpdate) RemoveImageIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveImageIDs(ids...)
	return cu
}

// RemoveImage removes "image" edges to Image entities.
func (cu *CategoryUpdate) RemoveImage(i ...*Image) *CategoryUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cu.RemoveImageIDs(ids...)
}

// ClearProduct clears all "product" edges to the Product entity.
func (cu *CategoryUpdate) ClearProduct() *CategoryUpdate {
	cu.mutation.ClearProduct()
	return cu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (cu *CategoryUpdate) RemoveProductIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveProductIDs(ids...)
	return cu
}

// RemoveProduct removes "product" edges to Product entities.
func (cu *CategoryUpdate) RemoveProduct(p ...*Product) *CategoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProductIDs(ids...)
}

// ClearBrand clears all "brand" edges to the Brand entity.
func (cu *CategoryUpdate) ClearBrand() *CategoryUpdate {
	cu.mutation.ClearBrand()
	return cu
}

// RemoveBrandIDs removes the "brand" edge to Brand entities by IDs.
func (cu *CategoryUpdate) RemoveBrandIDs(ids ...int) *CategoryUpdate {
	cu.mutation.RemoveBrandIDs(ids...)
	return cu
}

// RemoveBrand removes "brand" edges to Brand entities.
func (cu *CategoryUpdate) RemoveBrand(b ...*Brand) *CategoryUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cu.RemoveBrandIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CategoryUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CategoryUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CategoryUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CategoryUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := category.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CategoryUpdate) check() error {
	if v, ok := cu.mutation.Name(); ok {
		if err := category.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Category.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Description(); ok {
		if err := category.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Category.description": %w`, err)}
		}
	}
	return nil
}

func (cu *CategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.SetField(category.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if cu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ImageTable,
			Columns: category.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedImageIDs(); len(nodes) > 0 && !cu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ImageTable,
			Columns: category.ImagePrimaryKey,
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
	if nodes := cu.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ImageTable,
			Columns: category.ImagePrimaryKey,
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
	if cu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductTable,
			Columns: category.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProductIDs(); len(nodes) > 0 && !cu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductTable,
			Columns: category.ProductPrimaryKey,
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
	if nodes := cu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductTable,
			Columns: category.ProductPrimaryKey,
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
	if cu.mutation.BrandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.BrandTable,
			Columns: category.BrandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedBrandIDs(); len(nodes) > 0 && !cu.mutation.BrandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.BrandTable,
			Columns: category.BrandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.BrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.BrandTable,
			Columns: category.BrandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CategoryUpdateOne is the builder for updating a single Category entity.
type CategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CategoryMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *CategoryUpdateOne) SetUpdateTime(t time.Time) *CategoryUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetName sets the "name" field.
func (cuo *CategoryUpdateOne) SetName(s string) *CategoryUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *CategoryUpdateOne) SetDescription(s string) *CategoryUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (cuo *CategoryUpdateOne) AddImageIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddImageIDs(ids...)
	return cuo
}

// AddImage adds the "image" edges to the Image entity.
func (cuo *CategoryUpdateOne) AddImage(i ...*Image) *CategoryUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.AddImageIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (cuo *CategoryUpdateOne) AddProductIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddProductIDs(ids...)
	return cuo
}

// AddProduct adds the "product" edges to the Product entity.
func (cuo *CategoryUpdateOne) AddProduct(p ...*Product) *CategoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProductIDs(ids...)
}

// AddBrandIDs adds the "brand" edge to the Brand entity by IDs.
func (cuo *CategoryUpdateOne) AddBrandIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.AddBrandIDs(ids...)
	return cuo
}

// AddBrand adds the "brand" edges to the Brand entity.
func (cuo *CategoryUpdateOne) AddBrand(b ...*Brand) *CategoryUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.AddBrandIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cuo *CategoryUpdateOne) Mutation() *CategoryMutation {
	return cuo.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (cuo *CategoryUpdateOne) ClearImage() *CategoryUpdateOne {
	cuo.mutation.ClearImage()
	return cuo
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (cuo *CategoryUpdateOne) RemoveImageIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveImageIDs(ids...)
	return cuo
}

// RemoveImage removes "image" edges to Image entities.
func (cuo *CategoryUpdateOne) RemoveImage(i ...*Image) *CategoryUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return cuo.RemoveImageIDs(ids...)
}

// ClearProduct clears all "product" edges to the Product entity.
func (cuo *CategoryUpdateOne) ClearProduct() *CategoryUpdateOne {
	cuo.mutation.ClearProduct()
	return cuo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (cuo *CategoryUpdateOne) RemoveProductIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveProductIDs(ids...)
	return cuo
}

// RemoveProduct removes "product" edges to Product entities.
func (cuo *CategoryUpdateOne) RemoveProduct(p ...*Product) *CategoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProductIDs(ids...)
}

// ClearBrand clears all "brand" edges to the Brand entity.
func (cuo *CategoryUpdateOne) ClearBrand() *CategoryUpdateOne {
	cuo.mutation.ClearBrand()
	return cuo
}

// RemoveBrandIDs removes the "brand" edge to Brand entities by IDs.
func (cuo *CategoryUpdateOne) RemoveBrandIDs(ids ...int) *CategoryUpdateOne {
	cuo.mutation.RemoveBrandIDs(ids...)
	return cuo
}

// RemoveBrand removes "brand" edges to Brand entities.
func (cuo *CategoryUpdateOne) RemoveBrand(b ...*Brand) *CategoryUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cuo.RemoveBrandIDs(ids...)
}

// Where appends a list predicates to the CategoryUpdate builder.
func (cuo *CategoryUpdateOne) Where(ps ...predicate.Category) *CategoryUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CategoryUpdateOne) Select(field string, fields ...string) *CategoryUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Category entity.
func (cuo *CategoryUpdateOne) Save(ctx context.Context) (*Category, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CategoryUpdateOne) SaveX(ctx context.Context) *Category {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CategoryUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CategoryUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := category.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CategoryUpdateOne) check() error {
	if v, ok := cuo.mutation.Name(); ok {
		if err := category.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Category.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Description(); ok {
		if err := category.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Category.description": %w`, err)}
		}
	}
	return nil
}

func (cuo *CategoryUpdateOne) sqlSave(ctx context.Context) (_node *Category, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(category.Table, category.Columns, sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Category.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, category.FieldID)
		for _, f := range fields {
			if !category.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != category.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.SetField(category.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(category.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.SetField(category.FieldDescription, field.TypeString, value)
	}
	if cuo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ImageTable,
			Columns: category.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedImageIDs(); len(nodes) > 0 && !cuo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ImageTable,
			Columns: category.ImagePrimaryKey,
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
	if nodes := cuo.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   category.ImageTable,
			Columns: category.ImagePrimaryKey,
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
	if cuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductTable,
			Columns: category.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProductIDs(); len(nodes) > 0 && !cuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductTable,
			Columns: category.ProductPrimaryKey,
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
	if nodes := cuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductTable,
			Columns: category.ProductPrimaryKey,
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
	if cuo.mutation.BrandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.BrandTable,
			Columns: category.BrandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedBrandIDs(); len(nodes) > 0 && !cuo.mutation.BrandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.BrandTable,
			Columns: category.BrandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.BrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.BrandTable,
			Columns: category.BrandPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Category{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{category.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
