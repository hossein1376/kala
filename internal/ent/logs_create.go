// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hossein1376/kala/internal/ent/logs"
	"github.com/hossein1376/kala/internal/ent/user"
)

// LogsCreate is the builder for creating a Logs entity.
type LogsCreate struct {
	config
	mutation *LogsMutation
	hooks    []Hook
}

// SetAction sets the "action" field.
func (lc *LogsCreate) SetAction(s string) *LogsCreate {
	lc.mutation.SetAction(s)
	return lc
}

// SetNillableAction sets the "action" field if the given value is not nil.
func (lc *LogsCreate) SetNillableAction(s *string) *LogsCreate {
	if s != nil {
		lc.SetAction(*s)
	}
	return lc
}

// SetIP sets the "IP" field.
func (lc *LogsCreate) SetIP(s string) *LogsCreate {
	lc.mutation.SetIP(s)
	return lc
}

// SetAgent sets the "agent" field.
func (lc *LogsCreate) SetAgent(s string) *LogsCreate {
	lc.mutation.SetAgent(s)
	return lc
}

// SetDate sets the "date" field.
func (lc *LogsCreate) SetDate(t time.Time) *LogsCreate {
	lc.mutation.SetDate(t)
	return lc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lc *LogsCreate) SetUserID(id int) *LogsCreate {
	lc.mutation.SetUserID(id)
	return lc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lc *LogsCreate) SetNillableUserID(id *int) *LogsCreate {
	if id != nil {
		lc = lc.SetUserID(*id)
	}
	return lc
}

// SetUser sets the "user" edge to the User entity.
func (lc *LogsCreate) SetUser(u *User) *LogsCreate {
	return lc.SetUserID(u.ID)
}

// Mutation returns the LogsMutation object of the builder.
func (lc *LogsCreate) Mutation() *LogsMutation {
	return lc.mutation
}

// Save creates the Logs in the database.
func (lc *LogsCreate) Save(ctx context.Context) (*Logs, error) {
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LogsCreate) SaveX(ctx context.Context) *Logs {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LogsCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LogsCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LogsCreate) check() error {
	if _, ok := lc.mutation.IP(); !ok {
		return &ValidationError{Name: "IP", err: errors.New(`ent: missing required field "Logs.IP"`)}
	}
	if _, ok := lc.mutation.Agent(); !ok {
		return &ValidationError{Name: "agent", err: errors.New(`ent: missing required field "Logs.agent"`)}
	}
	if _, ok := lc.mutation.Date(); !ok {
		return &ValidationError{Name: "date", err: errors.New(`ent: missing required field "Logs.date"`)}
	}
	return nil
}

func (lc *LogsCreate) sqlSave(ctx context.Context) (*Logs, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LogsCreate) createSpec() (*Logs, *sqlgraph.CreateSpec) {
	var (
		_node = &Logs{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(logs.Table, sqlgraph.NewFieldSpec(logs.FieldID, field.TypeInt))
	)
	if value, ok := lc.mutation.Action(); ok {
		_spec.SetField(logs.FieldAction, field.TypeString, value)
		_node.Action = &value
	}
	if value, ok := lc.mutation.IP(); ok {
		_spec.SetField(logs.FieldIP, field.TypeString, value)
		_node.IP = value
	}
	if value, ok := lc.mutation.Agent(); ok {
		_spec.SetField(logs.FieldAgent, field.TypeString, value)
		_node.Agent = value
	}
	if value, ok := lc.mutation.Date(); ok {
		_spec.SetField(logs.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if nodes := lc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   logs.UserTable,
			Columns: []string{logs.UserColumn},
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

// LogsCreateBulk is the builder for creating many Logs entities in bulk.
type LogsCreateBulk struct {
	config
	err      error
	builders []*LogsCreate
}

// Save creates the Logs entities in the database.
func (lcb *LogsCreateBulk) Save(ctx context.Context) ([]*Logs, error) {
	if lcb.err != nil {
		return nil, lcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Logs, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LogsMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LogsCreateBulk) SaveX(ctx context.Context) []*Logs {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LogsCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LogsCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
