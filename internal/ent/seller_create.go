// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/address"
	"kala/internal/ent/category"
	"kala/internal/ent/order"
	"kala/internal/ent/product"
	"kala/internal/ent/seller"
	"kala/internal/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerCreate is the builder for creating a Seller entity.
type SellerCreate struct {
	config
	mutation *SellerMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *SellerCreate) SetCreateTime(t time.Time) *SellerCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SellerCreate) SetNillableCreateTime(t *time.Time) *SellerCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SellerCreate) SetUpdateTime(t time.Time) *SellerCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SellerCreate) SetNillableUpdateTime(t *time.Time) *SellerCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SellerCreate) SetName(s string) *SellerCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetDescription sets the "description" field.
func (sc *SellerCreate) SetDescription(s string) *SellerCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SellerCreate) SetNillableDescription(s *string) *SellerCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetRating sets the "rating" field.
func (sc *SellerCreate) SetRating(f float64) *SellerCreate {
	sc.mutation.SetRating(f)
	return sc
}

// SetRatingCount sets the "rating_count" field.
func (sc *SellerCreate) SetRatingCount(i int32) *SellerCreate {
	sc.mutation.SetRatingCount(i)
	return sc
}

// SetPhone sets the "phone" field.
func (sc *SellerCreate) SetPhone(s string) *SellerCreate {
	sc.mutation.SetPhone(s)
	return sc
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (sc *SellerCreate) AddProductIDs(ids ...int) *SellerCreate {
	sc.mutation.AddProductIDs(ids...)
	return sc
}

// AddProduct adds the "product" edges to the Product entity.
func (sc *SellerCreate) AddProduct(p ...*Product) *SellerCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddProductIDs(ids...)
}

// AddCategoryIDs adds the "category" edge to the Category entity by IDs.
func (sc *SellerCreate) AddCategoryIDs(ids ...int) *SellerCreate {
	sc.mutation.AddCategoryIDs(ids...)
	return sc
}

// AddCategory adds the "category" edges to the Category entity.
func (sc *SellerCreate) AddCategory(c ...*Category) *SellerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return sc.AddCategoryIDs(ids...)
}

// AddAddresIDs adds the "address" edge to the Address entity by IDs.
func (sc *SellerCreate) AddAddresIDs(ids ...int) *SellerCreate {
	sc.mutation.AddAddresIDs(ids...)
	return sc
}

// AddAddress adds the "address" edges to the Address entity.
func (sc *SellerCreate) AddAddress(a ...*Address) *SellerCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAddresIDs(ids...)
}

// AddOrderIDs adds the "order" edge to the Order entity by IDs.
func (sc *SellerCreate) AddOrderIDs(ids ...int) *SellerCreate {
	sc.mutation.AddOrderIDs(ids...)
	return sc
}

// AddOrder adds the "order" edges to the Order entity.
func (sc *SellerCreate) AddOrder(o ...*Order) *SellerCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return sc.AddOrderIDs(ids...)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (sc *SellerCreate) SetUserID(id int) *SellerCreate {
	sc.mutation.SetUserID(id)
	return sc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (sc *SellerCreate) SetNillableUserID(id *int) *SellerCreate {
	if id != nil {
		sc = sc.SetUserID(*id)
	}
	return sc
}

// SetUser sets the "user" edge to the User entity.
func (sc *SellerCreate) SetUser(u *User) *SellerCreate {
	return sc.SetUserID(u.ID)
}

// Mutation returns the SellerMutation object of the builder.
func (sc *SellerCreate) Mutation() *SellerMutation {
	return sc.mutation
}

// Save creates the Seller in the database.
func (sc *SellerCreate) Save(ctx context.Context) (*Seller, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SellerCreate) SaveX(ctx context.Context) *Seller {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SellerCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SellerCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SellerCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := seller.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := seller.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SellerCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Seller.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Seller.update_time"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Seller.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := seller.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Seller.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Rating(); !ok {
		return &ValidationError{Name: "rating", err: errors.New(`ent: missing required field "Seller.rating"`)}
	}
	if v, ok := sc.mutation.Rating(); ok {
		if err := seller.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "Seller.rating": %w`, err)}
		}
	}
	if _, ok := sc.mutation.RatingCount(); !ok {
		return &ValidationError{Name: "rating_count", err: errors.New(`ent: missing required field "Seller.rating_count"`)}
	}
	if v, ok := sc.mutation.RatingCount(); ok {
		if err := seller.RatingCountValidator(v); err != nil {
			return &ValidationError{Name: "rating_count", err: fmt.Errorf(`ent: validator failed for field "Seller.rating_count": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Seller.phone"`)}
	}
	if v, ok := sc.mutation.Phone(); ok {
		if err := seller.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Seller.phone": %w`, err)}
		}
	}
	return nil
}

func (sc *SellerCreate) sqlSave(ctx context.Context) (*Seller, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SellerCreate) createSpec() (*Seller, *sqlgraph.CreateSpec) {
	var (
		_node = &Seller{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(seller.Table, sqlgraph.NewFieldSpec(seller.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(seller.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(seller.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(seller.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(seller.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := sc.mutation.Rating(); ok {
		_spec.SetField(seller.FieldRating, field.TypeFloat64, value)
		_node.Rating = value
	}
	if value, ok := sc.mutation.RatingCount(); ok {
		_spec.SetField(seller.FieldRatingCount, field.TypeInt32, value)
		_node.RatingCount = value
	}
	if value, ok := sc.mutation.Phone(); ok {
		_spec.SetField(seller.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if nodes := sc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.ProductTable,
			Columns: []string{seller.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   seller.CategoryTable,
			Columns: seller.CategoryPrimaryKey,
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
	if nodes := sc.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.AddressTable,
			Columns: []string{seller.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   seller.OrderTable,
			Columns: []string{seller.OrderColumn},
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
	if nodes := sc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   seller.UserTable,
			Columns: []string{seller.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SellerCreateBulk is the builder for creating many Seller entities in bulk.
type SellerCreateBulk struct {
	config
	builders []*SellerCreate
}

// Save creates the Seller entities in the database.
func (scb *SellerCreateBulk) Save(ctx context.Context) ([]*Seller, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Seller, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SellerMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SellerCreateBulk) SaveX(ctx context.Context) []*Seller {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SellerCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SellerCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
