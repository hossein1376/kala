// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/attributevalue"
	"kala/internal/ent/brand"
	"kala/internal/ent/category"
	"kala/internal/ent/comment"
	"kala/internal/ent/image"
	"kala/internal/ent/order"
	"kala/internal/ent/product"
	"kala/internal/ent/subcategory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *ProductCreate) SetCreateTime(t time.Time) *ProductCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreateTime(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProductCreate) SetUpdateTime(t time.Time) *ProductCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdateTime(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetReview sets the "review" field.
func (pc *ProductCreate) SetReview(s string) *ProductCreate {
	pc.mutation.SetReview(s)
	return pc
}

// SetRating sets the "rating" field.
func (pc *ProductCreate) SetRating(f float64) *ProductCreate {
	pc.mutation.SetRating(f)
	return pc
}

// SetRatingCount sets the "rating_count" field.
func (pc *ProductCreate) SetRatingCount(i int32) *ProductCreate {
	pc.mutation.SetRatingCount(i)
	return pc
}

// SetPrice sets the "price" field.
func (pc *ProductCreate) SetPrice(i int32) *ProductCreate {
	pc.mutation.SetPrice(i)
	return pc
}

// SetQuantity sets the "quantity" field.
func (pc *ProductCreate) SetQuantity(i int32) *ProductCreate {
	pc.mutation.SetQuantity(i)
	return pc
}

// SetAvailable sets the "available" field.
func (pc *ProductCreate) SetAvailable(b bool) *ProductCreate {
	pc.mutation.SetAvailable(b)
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProductCreate) SetStatus(b bool) *ProductCreate {
	pc.mutation.SetStatus(b)
	return pc
}

// AddValueIDs adds the "values" edge to the AttributeValue entity by IDs.
func (pc *ProductCreate) AddValueIDs(ids ...int) *ProductCreate {
	pc.mutation.AddValueIDs(ids...)
	return pc
}

// AddValues adds the "values" edges to the AttributeValue entity.
func (pc *ProductCreate) AddValues(a ...*AttributeValue) *ProductCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return pc.AddValueIDs(ids...)
}

// AddCommentIDs adds the "comment" edge to the Comment entity by IDs.
func (pc *ProductCreate) AddCommentIDs(ids ...int) *ProductCreate {
	pc.mutation.AddCommentIDs(ids...)
	return pc
}

// AddComment adds the "comment" edges to the Comment entity.
func (pc *ProductCreate) AddComment(c ...*Comment) *ProductCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCommentIDs(ids...)
}

// AddImageIDs adds the "image" edge to the Image entity by IDs.
func (pc *ProductCreate) AddImageIDs(ids ...int) *ProductCreate {
	pc.mutation.AddImageIDs(ids...)
	return pc
}

// AddImage adds the "image" edges to the Image entity.
func (pc *ProductCreate) AddImage(i ...*Image) *ProductCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return pc.AddImageIDs(ids...)
}

// AddOrderIDs adds the "order" edge to the Order entity by IDs.
func (pc *ProductCreate) AddOrderIDs(ids ...int) *ProductCreate {
	pc.mutation.AddOrderIDs(ids...)
	return pc
}

// AddOrder adds the "order" edges to the Order entity.
func (pc *ProductCreate) AddOrder(o ...*Order) *ProductCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrderIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (pc *ProductCreate) AddCategoryIDs(ids ...int) *ProductCreate {
	pc.mutation.AddCategoryIDs(ids...)
	return pc
}

// AddCategory adds the "category" edges to the Category entity.
func (pc *ProductCreate) AddCategory(c ...*Category) *ProductCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCategoryIDs(ids...)
}

// AddSubCategoryIDs adds the "sub_category" edge to the SubCategory entity by IDs.
func (pc *ProductCreate) AddSubCategoryIDs(ids ...int) *ProductCreate {
	pc.mutation.AddSubCategoryIDs(ids...)
	return pc
}

// AddSubCategory adds the "sub_category" edges to the SubCategory entity.
func (pc *ProductCreate) AddSubCategory(s ...*SubCategory) *ProductCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSubCategoryIDs(ids...)
}

// SetBrandID sets the "brand" edge to the Brand entity by ID.
func (pc *ProductCreate) SetBrandID(id int) *ProductCreate {
	pc.mutation.SetBrandID(id)
	return pc
}

// SetNillableBrandID sets the "brand" edge to the Brand entity by ID if the given value is not nil.
func (pc *ProductCreate) SetNillableBrandID(id *int) *ProductCreate {
	if id != nil {
		pc = pc.SetBrandID(*id)
	}
	return pc
}

// SetBrand sets the "brand" edge to the Brand entity.
func (pc *ProductCreate) SetBrand(b *Brand) *ProductCreate {
	return pc.SetBrandID(b.ID)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := product.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := product.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Product.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Product.update_time"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Product.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Review(); !ok {
		return &ValidationError{Name: "review", err: errors.New(`ent: missing required field "Product.review"`)}
	}
	if _, ok := pc.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New(`ent: missing required field "Product.rating"`)}
	}
	if v, ok := pc.mutation.Rating(); ok {
		if err := product.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Product.rating": %w`, err)}
		}
	}
	if _, ok := pc.mutation.RatingCount(); !ok {
		return &ValidationError{Name: "rating_count", err: errors.New(`ent: missing required field "Product.rating_count"`)}
	}
	if v, ok := pc.mutation.RatingCount(); ok {
		if err := product.RatingCountValidator(v); err != nil {
			return &ValidationError{Name: "rating_count", err: fmt.Errorf(`ent: validator failed for field "Product.rating_count": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Product.price"`)}
	}
	if v, ok := pc.mutation.Price(); ok {
		if err := product.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Product.price": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Product.quantity"`)}
	}
	if v, ok := pc.mutation.Quantity(); ok {
		if err := product.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "Product.quantity": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Available(); !ok {
		return &ValidationError{Name: "available", err: errors.New(`ent: missing required field "Product.available"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Product.status"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(product.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(product.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Review(); ok {
		_spec.SetField(product.FieldReview, field.TypeString, value)
		_node.Review = value
	}
	if value, ok := pc.mutation.Rating(); ok {
		_spec.SetField(product.FieldRating, field.TypeFloat64, value)
		_node.Rating = value
	}
	if value, ok := pc.mutation.RatingCount(); ok {
		_spec.SetField(product.FieldRatingCount, field.TypeInt32, value)
		_node.RatingCount = value
	}
	if value, ok := pc.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeInt32, value)
		_node.Price = value
	}
	if value, ok := pc.mutation.Quantity(); ok {
		_spec.SetField(product.FieldQuantity, field.TypeInt32, value)
		_node.Quantity = value
	}
	if value, ok := pc.mutation.Available(); ok {
		_spec.SetField(product.FieldAvailable, field.TypeBool, value)
		_node.Available = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeBool, value)
		_node.Status = value
	}
	if nodes := pc.mutation.ValuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.ValuesTable,
			Columns: []string{product.ValuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attributevalue.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.CommentTable,
			Columns: product.CommentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.ImageTable,
			Columns: product.ImagePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.OrderTable,
			Columns: product.OrderPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.CategoryTable,
			Columns: product.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(category.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubCategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.SubCategoryTable,
			Columns: product.SubCategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subcategory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BrandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   product.BrandTable,
			Columns: []string{product.BrandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(brand.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.brand_product = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
