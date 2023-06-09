// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kala/internal/ent/predicate"
	"kala/internal/ent/seller"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SellerDelete is the builder for deleting a Seller entity.
type SellerDelete struct {
	config
	hooks    []Hook
	mutation *SellerMutation
}

// Where appends a list predicates to the SellerDelete builder.
func (sd *SellerDelete) Where(ps ...predicate.Seller) *SellerDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SellerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SellerDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *SellerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(seller.Table, sqlgraph.NewFieldSpec(seller.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SellerDeleteOne is the builder for deleting a single Seller entity.
type SellerDeleteOne struct {
	sd *SellerDelete
}

// Where appends a list predicates to the SellerDelete builder.
func (sdo *SellerDeleteOne) Where(ps ...predicate.Seller) *SellerDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SellerDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{seller.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SellerDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}
