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
	"github.com/hossein1376/kala/internal/ent/category"
	"github.com/hossein1376/kala/internal/ent/image"
	"github.com/hossein1376/kala/internal/ent/predicate"
	"github.com/hossein1376/kala/internal/ent/product"
	"github.com/hossein1376/kala/internal/ent/subcategory"
)

// SubCategoryUpdate is the builder for updating SubCategory entities.
type SubCategoryUpdate struct {
	config
	hooks    []Hook
	mutation *SubCategoryMutation
}

// Where appends a list predicates to the SubCategoryUpdate builder.
func (scu *SubCategoryUpdate) Where(ps ...predicate.SubCategory) *SubCategoryUpdate {
	scu.mutation.Where(ps...)
	return scu
}

// SetUpdateTime sets the "update_time" field.
func (scu *SubCategoryUpdate) SetUpdateTime(t time.Time) *SubCategoryUpdate {
	scu.mutation.SetUpdateTime(t)
	return scu
}

// SetName sets the "name" field.
func (scu *SubCategoryUpdate) SetName(s string) *SubCategoryUpdate {
	scu.mutation.SetName(s)
	return scu
}

// SetDescription sets the "description" field.
func (scu *SubCategoryUpdate) SetDescription(s string) *SubCategoryUpdate {
	scu.mutation.SetDescription(s)
	return scu
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (scu *SubCategoryUpdate) AddImageIDs(ids ...int) *SubCategoryUpdate {
	scu.mutation.AddImageIDs(ids...)
	return scu
}

// AddImage adds the "image" edges to the Image entity.
func (scu *SubCategoryUpdate) AddImage(i ...*Image) *SubCategoryUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return scu.AddImageIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (scu *SubCategoryUpdate) AddProductIDs(ids ...int) *SubCategoryUpdate {
	scu.mutation.AddProductIDs(ids...)
	return scu
}

// AddProduct adds the "product" edges to the Product entity.
func (scu *SubCategoryUpdate) AddProduct(p ...*Product) *SubCategoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scu.AddProductIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (scu *SubCategoryUpdate) SetCategoryID(id int) *SubCategoryUpdate {
	scu.mutation.SetCategoryID(id)
	return scu
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (scu *SubCategoryUpdate) SetNillableCategoryID(id *int) *SubCategoryUpdate {
	if id != nil {
		scu = scu.SetCategoryID(*id)
	}
	return scu
}

// SetCategory sets the "category" edge to the Category entity.
func (scu *SubCategoryUpdate) SetCategory(c *Category) *SubCategoryUpdate {
	return scu.SetCategoryID(c.ID)
}

// Mutation returns the SubCategoryMutation object of the builder.
func (scu *SubCategoryUpdate) Mutation() *SubCategoryMutation {
	return scu.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (scu *SubCategoryUpdate) ClearImage() *SubCategoryUpdate {
	scu.mutation.ClearImage()
	return scu
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (scu *SubCategoryUpdate) RemoveImageIDs(ids ...int) *SubCategoryUpdate {
	scu.mutation.RemoveImageIDs(ids...)
	return scu
}

// RemoveImage removes "image" edges to Image entities.
func (scu *SubCategoryUpdate) RemoveImage(i ...*Image) *SubCategoryUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return scu.RemoveImageIDs(ids...)
}

// ClearProduct clears all "product" edges to the Product entity.
func (scu *SubCategoryUpdate) ClearProduct() *SubCategoryUpdate {
	scu.mutation.ClearProduct()
	return scu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (scu *SubCategoryUpdate) RemoveProductIDs(ids ...int) *SubCategoryUpdate {
	scu.mutation.RemoveProductIDs(ids...)
	return scu
}

// RemoveProduct removes "product" edges to Product entities.
func (scu *SubCategoryUpdate) RemoveProduct(p ...*Product) *SubCategoryUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scu.RemoveProductIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (scu *SubCategoryUpdate) ClearCategory() *SubCategoryUpdate {
	scu.mutation.ClearCategory()
	return scu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (scu *SubCategoryUpdate) Save(ctx context.Context) (int, error) {
	scu.defaults()
	return withHooks(ctx, scu.sqlSave, scu.mutation, scu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scu *SubCategoryUpdate) SaveX(ctx context.Context) int {
	affected, err := scu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (scu *SubCategoryUpdate) Exec(ctx context.Context) error {
	_, err := scu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scu *SubCategoryUpdate) ExecX(ctx context.Context) {
	if err := scu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scu *SubCategoryUpdate) defaults() {
	if _, ok := scu.mutation.UpdateTime(); !ok {
		v := subcategory.UpdateDefaultUpdateTime()
		scu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scu *SubCategoryUpdate) check() error {
	if v, ok := scu.mutation.Name(); ok {
		if err := subcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SubCategory.name": %w`, err)}
		}
	}
	if v, ok := scu.mutation.Description(); ok {
		if err := subcategory.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "SubCategory.description": %w`, err)}
		}
	}
	return nil
}

func (scu *SubCategoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := scu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(subcategory.Table, subcategory.Columns, sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeInt))
	if ps := scu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scu.mutation.UpdateTime(); ok {
		_spec.SetField(subcategory.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := scu.mutation.Name(); ok {
		_spec.SetField(subcategory.FieldName, field.TypeString, value)
	}
	if value, ok := scu.mutation.Description(); ok {
		_spec.SetField(subcategory.FieldDescription, field.TypeString, value)
	}
	if scu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subcategory.ImageTable,
			Columns: subcategory.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedImageIDs(); len(nodes) > 0 && !scu.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subcategory.ImageTable,
			Columns: subcategory.ImagePrimaryKey,
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
	if nodes := scu.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subcategory.ImageTable,
			Columns: subcategory.ImagePrimaryKey,
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
	if scu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subcategory.ProductTable,
			Columns: subcategory.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.RemovedProductIDs(); len(nodes) > 0 && !scu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subcategory.ProductTable,
			Columns: subcategory.ProductPrimaryKey,
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
	if nodes := scu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subcategory.ProductTable,
			Columns: subcategory.ProductPrimaryKey,
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
	if scu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, scu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	scu.mutation.done = true
	return n, nil
}

// SubCategoryUpdateOne is the builder for updating a single SubCategory entity.
type SubCategoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubCategoryMutation
}

// SetUpdateTime sets the "update_time" field.
func (scuo *SubCategoryUpdateOne) SetUpdateTime(t time.Time) *SubCategoryUpdateOne {
	scuo.mutation.SetUpdateTime(t)
	return scuo
}

// SetName sets the "name" field.
func (scuo *SubCategoryUpdateOne) SetName(s string) *SubCategoryUpdateOne {
	scuo.mutation.SetName(s)
	return scuo
}

// SetDescription sets the "description" field.
func (scuo *SubCategoryUpdateOne) SetDescription(s string) *SubCategoryUpdateOne {
	scuo.mutation.SetDescription(s)
	return scuo
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (scuo *SubCategoryUpdateOne) AddImageIDs(ids ...int) *SubCategoryUpdateOne {
	scuo.mutation.AddImageIDs(ids...)
	return scuo
}

// AddImage adds the "image" edges to the Image entity.
func (scuo *SubCategoryUpdateOne) AddImage(i ...*Image) *SubCategoryUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return scuo.AddImageIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (scuo *SubCategoryUpdateOne) AddProductIDs(ids ...int) *SubCategoryUpdateOne {
	scuo.mutation.AddProductIDs(ids...)
	return scuo
}

// AddProduct adds the "product" edges to the Product entity.
func (scuo *SubCategoryUpdateOne) AddProduct(p ...*Product) *SubCategoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scuo.AddProductIDs(ids...)
}

// SetCategoryID sets the "category" edge to the Category entity by ID.
func (scuo *SubCategoryUpdateOne) SetCategoryID(id int) *SubCategoryUpdateOne {
	scuo.mutation.SetCategoryID(id)
	return scuo
}

// SetNillableCategoryID sets the "category" edge to the Category entity by ID if the given value is not nil.
func (scuo *SubCategoryUpdateOne) SetNillableCategoryID(id *int) *SubCategoryUpdateOne {
	if id != nil {
		scuo = scuo.SetCategoryID(*id)
	}
	return scuo
}

// SetCategory sets the "category" edge to the Category entity.
func (scuo *SubCategoryUpdateOne) SetCategory(c *Category) *SubCategoryUpdateOne {
	return scuo.SetCategoryID(c.ID)
}

// Mutation returns the SubCategoryMutation object of the builder.
func (scuo *SubCategoryUpdateOne) Mutation() *SubCategoryMutation {
	return scuo.mutation
}

// ClearImage clears all "image" edges to the Image entity.
func (scuo *SubCategoryUpdateOne) ClearImage() *SubCategoryUpdateOne {
	scuo.mutation.ClearImage()
	return scuo
}

// RemoveImageIDs removes the "image" edge to Image entities by IDs.
func (scuo *SubCategoryUpdateOne) RemoveImageIDs(ids ...int) *SubCategoryUpdateOne {
	scuo.mutation.RemoveImageIDs(ids...)
	return scuo
}

// RemoveImage removes "image" edges to Image entities.
func (scuo *SubCategoryUpdateOne) RemoveImage(i ...*Image) *SubCategoryUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return scuo.RemoveImageIDs(ids...)
}

// ClearProduct clears all "product" edges to the Product entity.
func (scuo *SubCategoryUpdateOne) ClearProduct() *SubCategoryUpdateOne {
	scuo.mutation.ClearProduct()
	return scuo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (scuo *SubCategoryUpdateOne) RemoveProductIDs(ids ...int) *SubCategoryUpdateOne {
	scuo.mutation.RemoveProductIDs(ids...)
	return scuo
}

// RemoveProduct removes "product" edges to Product entities.
func (scuo *SubCategoryUpdateOne) RemoveProduct(p ...*Product) *SubCategoryUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return scuo.RemoveProductIDs(ids...)
}

// ClearCategory clears the "category" edge to the Category entity.
func (scuo *SubCategoryUpdateOne) ClearCategory() *SubCategoryUpdateOne {
	scuo.mutation.ClearCategory()
	return scuo
}

// Where appends a list predicates to the SubCategoryUpdate builder.
func (scuo *SubCategoryUpdateOne) Where(ps ...predicate.SubCategory) *SubCategoryUpdateOne {
	scuo.mutation.Where(ps...)
	return scuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (scuo *SubCategoryUpdateOne) Select(field string, fields ...string) *SubCategoryUpdateOne {
	scuo.fields = append([]string{field}, fields...)
	return scuo
}

// Save executes the query and returns the updated SubCategory entity.
func (scuo *SubCategoryUpdateOne) Save(ctx context.Context) (*SubCategory, error) {
	scuo.defaults()
	return withHooks(ctx, scuo.sqlSave, scuo.mutation, scuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (scuo *SubCategoryUpdateOne) SaveX(ctx context.Context) *SubCategory {
	node, err := scuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (scuo *SubCategoryUpdateOne) Exec(ctx context.Context) error {
	_, err := scuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scuo *SubCategoryUpdateOne) ExecX(ctx context.Context) {
	if err := scuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (scuo *SubCategoryUpdateOne) defaults() {
	if _, ok := scuo.mutation.UpdateTime(); !ok {
		v := subcategory.UpdateDefaultUpdateTime()
		scuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (scuo *SubCategoryUpdateOne) check() error {
	if v, ok := scuo.mutation.Name(); ok {
		if err := subcategory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "SubCategory.name": %w`, err)}
		}
	}
	if v, ok := scuo.mutation.Description(); ok {
		if err := subcategory.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "SubCategory.description": %w`, err)}
		}
	}
	return nil
}

func (scuo *SubCategoryUpdateOne) sqlSave(ctx context.Context) (_node *SubCategory, err error) {
	if err := scuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(subcategory.Table, subcategory.Columns, sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeInt))
	id, ok := scuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SubCategory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := scuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subcategory.FieldID)
		for _, f := range fields {
			if !subcategory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := scuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := scuo.mutation.UpdateTime(); ok {
		_spec.SetField(subcategory.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := scuo.mutation.Name(); ok {
		_spec.SetField(subcategory.FieldName, field.TypeString, value)
	}
	if value, ok := scuo.mutation.Description(); ok {
		_spec.SetField(subcategory.FieldDescription, field.TypeString, value)
	}
	if scuo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subcategory.ImageTable,
			Columns: subcategory.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedImageIDs(); len(nodes) > 0 && !scuo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subcategory.ImageTable,
			Columns: subcategory.ImagePrimaryKey,
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
	if nodes := scuo.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   subcategory.ImageTable,
			Columns: subcategory.ImagePrimaryKey,
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
	if scuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subcategory.ProductTable,
			Columns: subcategory.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.RemovedProductIDs(); len(nodes) > 0 && !scuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subcategory.ProductTable,
			Columns: subcategory.ProductPrimaryKey,
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
	if nodes := scuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   subcategory.ProductTable,
			Columns: subcategory.ProductPrimaryKey,
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
	if scuo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := scuo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subcategory.CategoryTable,
			Columns: []string{subcategory.CategoryColumn},
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
	_node = &SubCategory{config: scuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, scuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subcategory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	scuo.mutation.done = true
	return _node, nil
}
