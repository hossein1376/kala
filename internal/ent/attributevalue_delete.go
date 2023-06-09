// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kala/internal/ent/attributevalue"
	"kala/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AttributeValueDelete is the builder for deleting a AttributeValue entity.
type AttributeValueDelete struct {
	config
	hooks    []Hook
	mutation *AttributeValueMutation
}

// Where appends a list predicates to the AttributeValueDelete builder.
func (avd *AttributeValueDelete) Where(ps ...predicate.AttributeValue) *AttributeValueDelete {
	avd.mutation.Where(ps...)
	return avd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (avd *AttributeValueDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, avd.sqlExec, avd.mutation, avd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (avd *AttributeValueDelete) ExecX(ctx context.Context) int {
	n, err := avd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (avd *AttributeValueDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(attributevalue.Table, sqlgraph.NewFieldSpec(attributevalue.FieldID, field.TypeInt))
	if ps := avd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, avd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	avd.mutation.done = true
	return affected, err
}

// AttributeValueDeleteOne is the builder for deleting a single AttributeValue entity.
type AttributeValueDeleteOne struct {
	avd *AttributeValueDelete
}

// Where appends a list predicates to the AttributeValueDelete builder.
func (avdo *AttributeValueDeleteOne) Where(ps ...predicate.AttributeValue) *AttributeValueDeleteOne {
	avdo.avd.mutation.Where(ps...)
	return avdo
}

// Exec executes the deletion query.
func (avdo *AttributeValueDeleteOne) Exec(ctx context.Context) error {
	n, err := avdo.avd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{attributevalue.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (avdo *AttributeValueDeleteOne) ExecX(ctx context.Context) {
	if err := avdo.Exec(ctx); err != nil {
		panic(err)
	}
}
