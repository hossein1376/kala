// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/comment"
	"kala/internal/ent/cons"
	"kala/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConsUpdate is the builder for updating Cons entities.
type ConsUpdate struct {
	config
	hooks    []Hook
	mutation *ConsMutation
}

// Where appends a list predicates to the ConsUpdate builder.
func (cu *ConsUpdate) Where(ps ...predicate.Cons) *ConsUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCon sets the "con" field.
func (cu *ConsUpdate) SetCon(s string) *ConsUpdate {
	cu.mutation.SetCon(s)
	return cu
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (cu *ConsUpdate) SetCommentID(id int) *ConsUpdate {
	cu.mutation.SetCommentID(id)
	return cu
}

// SetNillableCommentID sets the "comment" edge to the Comment entity by ID if the given value is not nil.
func (cu *ConsUpdate) SetNillableCommentID(id *int) *ConsUpdate {
	if id != nil {
		cu = cu.SetCommentID(*id)
	}
	return cu
}

// SetComment sets the "comment" edge to the Comment entity.
func (cu *ConsUpdate) SetComment(c *Comment) *ConsUpdate {
	return cu.SetCommentID(c.ID)
}

// Mutation returns the ConsMutation object of the builder.
func (cu *ConsUpdate) Mutation() *ConsMutation {
	return cu.mutation
}

// ClearComment clears the "comment" edge to the Comment entity.
func (cu *ConsUpdate) ClearComment() *ConsUpdate {
	cu.mutation.ClearComment()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConsUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConsUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConsUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConsUpdate) check() error {
	if v, ok := cu.mutation.Con(); ok {
		if err := cons.ConValidator(v); err != nil {
			return &ValidationError{Name: "con", err: fmt.Errorf(`ent: validator failed for field "Cons.con": %w`, err)}
		}
	}
	return nil
}

func (cu *ConsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(cons.Table, cons.Columns, sqlgraph.NewFieldSpec(cons.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Con(); ok {
		_spec.SetField(cons.FieldCon, field.TypeString, value)
	}
	if cu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cons.CommentTable,
			Columns: []string{cons.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cons.CommentTable,
			Columns: []string{cons.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cons.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConsUpdateOne is the builder for updating a single Cons entity.
type ConsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConsMutation
}

// SetCon sets the "con" field.
func (cuo *ConsUpdateOne) SetCon(s string) *ConsUpdateOne {
	cuo.mutation.SetCon(s)
	return cuo
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (cuo *ConsUpdateOne) SetCommentID(id int) *ConsUpdateOne {
	cuo.mutation.SetCommentID(id)
	return cuo
}

// SetNillableCommentID sets the "comment" edge to the Comment entity by ID if the given value is not nil.
func (cuo *ConsUpdateOne) SetNillableCommentID(id *int) *ConsUpdateOne {
	if id != nil {
		cuo = cuo.SetCommentID(*id)
	}
	return cuo
}

// SetComment sets the "comment" edge to the Comment entity.
func (cuo *ConsUpdateOne) SetComment(c *Comment) *ConsUpdateOne {
	return cuo.SetCommentID(c.ID)
}

// Mutation returns the ConsMutation object of the builder.
func (cuo *ConsUpdateOne) Mutation() *ConsMutation {
	return cuo.mutation
}

// ClearComment clears the "comment" edge to the Comment entity.
func (cuo *ConsUpdateOne) ClearComment() *ConsUpdateOne {
	cuo.mutation.ClearComment()
	return cuo
}

// Where appends a list predicates to the ConsUpdate builder.
func (cuo *ConsUpdateOne) Where(ps ...predicate.Cons) *ConsUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConsUpdateOne) Select(field string, fields ...string) *ConsUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cons entity.
func (cuo *ConsUpdateOne) Save(ctx context.Context) (*Cons, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConsUpdateOne) SaveX(ctx context.Context) *Cons {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConsUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConsUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConsUpdateOne) check() error {
	if v, ok := cuo.mutation.Con(); ok {
		if err := cons.ConValidator(v); err != nil {
			return &ValidationError{Name: "con", err: fmt.Errorf(`ent: validator failed for field "Cons.con": %w`, err)}
		}
	}
	return nil
}

func (cuo *ConsUpdateOne) sqlSave(ctx context.Context) (_node *Cons, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(cons.Table, cons.Columns, sqlgraph.NewFieldSpec(cons.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Cons.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cons.FieldID)
		for _, f := range fields {
			if !cons.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cons.FieldID {
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
	if value, ok := cuo.mutation.Con(); ok {
		_spec.SetField(cons.FieldCon, field.TypeString, value)
	}
	if cuo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cons.CommentTable,
			Columns: []string{cons.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   cons.CommentTable,
			Columns: []string{cons.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Cons{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cons.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
