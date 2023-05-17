// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/address"
	"kala/internal/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AddressCreate is the builder for creating a Address entity.
type AddressCreate struct {
	config
	mutation *AddressMutation
	hooks    []Hook
}

// SetAddress sets the "address" field.
func (ac *AddressCreate) SetAddress(s string) *AddressCreate {
	ac.mutation.SetAddress(s)
	return ac
}

// SetZipCode sets the "zip_code" field.
func (ac *AddressCreate) SetZipCode(s string) *AddressCreate {
	ac.mutation.SetZipCode(s)
	return ac
}

// SetPhone sets the "phone" field.
func (ac *AddressCreate) SetPhone(s string) *AddressCreate {
	ac.mutation.SetPhone(s)
	return ac
}

// SetCoordinates sets the "coordinates" field.
func (ac *AddressCreate) SetCoordinates(s string) *AddressCreate {
	ac.mutation.SetCoordinates(s)
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *AddressCreate) SetUserID(id int) *AddressCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ac *AddressCreate) SetNillableUserID(id *int) *AddressCreate {
	if id != nil {
		ac = ac.SetUserID(*id)
	}
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *AddressCreate) SetUser(u *User) *AddressCreate {
	return ac.SetUserID(u.ID)
}

// Mutation returns the AddressMutation object of the builder.
func (ac *AddressCreate) Mutation() *AddressMutation {
	return ac.mutation
}

// Save creates the Address in the database.
func (ac *AddressCreate) Save(ctx context.Context) (*Address, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AddressCreate) SaveX(ctx context.Context) *Address {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AddressCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AddressCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AddressCreate) check() error {
	if _, ok := ac.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Address.address"`)}
	}
	if v, ok := ac.mutation.Address(); ok {
		if err := address.AddressValidator(v); err != nil {
			return &ValidationError{Name: "address", err: fmt.Errorf(`ent: validator failed for field "Address.address": %w`, err)}
		}
	}
	if _, ok := ac.mutation.ZipCode(); !ok {
		return &ValidationError{Name: "zip_code", err: errors.New(`ent: missing required field "Address.zip_code"`)}
	}
	if v, ok := ac.mutation.ZipCode(); ok {
		if err := address.ZipCodeValidator(v); err != nil {
			return &ValidationError{Name: "zip_code", err: fmt.Errorf(`ent: validator failed for field "Address.zip_code": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent: missing required field "Address.phone"`)}
	}
	if v, ok := ac.mutation.Phone(); ok {
		if err := address.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "Address.phone": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Coordinates(); !ok {
		return &ValidationError{Name: "coordinates", err: errors.New(`ent: missing required field "Address.coordinates"`)}
	}
	if v, ok := ac.mutation.Coordinates(); ok {
		if err := address.CoordinatesValidator(v); err != nil {
			return &ValidationError{Name: "coordinates", err: fmt.Errorf(`ent: validator failed for field "Address.coordinates": %w`, err)}
		}
	}
	return nil
}

func (ac *AddressCreate) sqlSave(ctx context.Context) (*Address, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AddressCreate) createSpec() (*Address, *sqlgraph.CreateSpec) {
	var (
		_node = &Address{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(address.Table, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.Address(); ok {
		_spec.SetField(address.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := ac.mutation.ZipCode(); ok {
		_spec.SetField(address.FieldZipCode, field.TypeString, value)
		_node.ZipCode = value
	}
	if value, ok := ac.mutation.Phone(); ok {
		_spec.SetField(address.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := ac.mutation.Coordinates(); ok {
		_spec.SetField(address.FieldCoordinates, field.TypeString, value)
		_node.Coordinates = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   address.UserTable,
			Columns: []string{address.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AddressCreateBulk is the builder for creating many Address entities in bulk.
type AddressCreateBulk struct {
	config
	builders []*AddressCreate
}

// Save creates the Address entities in the database.
func (acb *AddressCreateBulk) Save(ctx context.Context) ([]*Address, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Address, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AddressMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AddressCreateBulk) SaveX(ctx context.Context) []*Address {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AddressCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AddressCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
