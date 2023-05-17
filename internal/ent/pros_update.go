// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kala/internal/ent/comment"
	"kala/internal/ent/predicate"
	"kala/internal/ent/pros"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProsUpdate is the builder for updating Pros entities.
type ProsUpdate struct {
	config
	hooks    []Hook
	mutation *ProsMutation
}

// Where appends a list predicates to the ProsUpdate builder.
func (pu *ProsUpdate) Where(ps ...predicate.Pros) *ProsUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPro sets the "pro" field.
func (pu *ProsUpdate) SetPro(s string) *ProsUpdate {
	pu.mutation.SetPro(s)
	return pu
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (pu *ProsUpdate) SetCommentID(id int) *ProsUpdate {
	pu.mutation.SetCommentID(id)
	return pu
}

// SetNillableCommentID sets the "comment" edge to the Comment entity by ID if the given value is not nil.
func (pu *ProsUpdate) SetNillableCommentID(id *int) *ProsUpdate {
	if id != nil {
		pu = pu.SetCommentID(*id)
	}
	return pu
}

// SetComment sets the "comment" edge to the Comment entity.
func (pu *ProsUpdate) SetComment(c *Comment) *ProsUpdate {
	return pu.SetCommentID(c.ID)
}

// Mutation returns the ProsMutation object of the builder.
func (pu *ProsUpdate) Mutation() *ProsMutation {
	return pu.mutation
}

// ClearComment clears the "comment" edge to the Comment entity.
func (pu *ProsUpdate) ClearComment() *ProsUpdate {
	pu.mutation.ClearComment()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProsUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProsUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProsUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProsUpdate) check() error {
	if v, ok := pu.mutation.Pro(); ok {
		if err := pros.ProValidator(v); err != nil {
			return &ValidationError{Name: "pro", err: fmt.Errorf(`ent: validator failed for field "Pros.pro": %w`, err)}
		}
	}
	return nil
}

func (pu *ProsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pros.Table, pros.Columns, sqlgraph.NewFieldSpec(pros.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Pro(); ok {
		_spec.SetField(pros.FieldPro, field.TypeString, value)
	}
	if pu.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pros.CommentTable,
			Columns: []string{pros.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pros.CommentTable,
			Columns: []string{pros.CommentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pros.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProsUpdateOne is the builder for updating a single Pros entity.
type ProsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProsMutation
}

// SetPro sets the "pro" field.
func (puo *ProsUpdateOne) SetPro(s string) *ProsUpdateOne {
	puo.mutation.SetPro(s)
	return puo
}

// SetCommentID sets the "comment" edge to the Comment entity by ID.
func (puo *ProsUpdateOne) SetCommentID(id int) *ProsUpdateOne {
	puo.mutation.SetCommentID(id)
	return puo
}

// SetNillableCommentID sets the "comment" edge to the Comment entity by ID if the given value is not nil.
func (puo *ProsUpdateOne) SetNillableCommentID(id *int) *ProsUpdateOne {
	if id != nil {
		puo = puo.SetCommentID(*id)
	}
	return puo
}

// SetComment sets the "comment" edge to the Comment entity.
func (puo *ProsUpdateOne) SetComment(c *Comment) *ProsUpdateOne {
	return puo.SetCommentID(c.ID)
}

// Mutation returns the ProsMutation object of the builder.
func (puo *ProsUpdateOne) Mutation() *ProsMutation {
	return puo.mutation
}

// ClearComment clears the "comment" edge to the Comment entity.
func (puo *ProsUpdateOne) ClearComment() *ProsUpdateOne {
	puo.mutation.ClearComment()
	return puo
}

// Where appends a list predicates to the ProsUpdate builder.
func (puo *ProsUpdateOne) Where(ps ...predicate.Pros) *ProsUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProsUpdateOne) Select(field string, fields ...string) *ProsUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pros entity.
func (puo *ProsUpdateOne) Save(ctx context.Context) (*Pros, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProsUpdateOne) SaveX(ctx context.Context) *Pros {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProsUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProsUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProsUpdateOne) check() error {
	if v, ok := puo.mutation.Pro(); ok {
		if err := pros.ProValidator(v); err != nil {
			return &ValidationError{Name: "pro", err: fmt.Errorf(`ent: validator failed for field "Pros.pro": %w`, err)}
		}
	}
	return nil
}

func (puo *ProsUpdateOne) sqlSave(ctx context.Context) (_node *Pros, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pros.Table, pros.Columns, sqlgraph.NewFieldSpec(pros.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pros.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pros.FieldID)
		for _, f := range fields {
			if !pros.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pros.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Pro(); ok {
		_spec.SetField(pros.FieldPro, field.TypeString, value)
	}
	if puo.mutation.CommentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pros.CommentTable,
			Columns: []string{pros.CommentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CommentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pros.CommentTable,
			Columns: []string{pros.CommentColumn},
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
	_node = &Pros{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pros.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}