// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/predicate"
	"github.com/Tracy-coder/e-mall/data/ent/productimg"
)

// ProductImgDelete is the builder for deleting a ProductImg entity.
type ProductImgDelete struct {
	config
	hooks    []Hook
	mutation *ProductImgMutation
}

// Where appends a list predicates to the ProductImgDelete builder.
func (pid *ProductImgDelete) Where(ps ...predicate.ProductImg) *ProductImgDelete {
	pid.mutation.Where(ps...)
	return pid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pid *ProductImgDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pid.sqlExec, pid.mutation, pid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pid *ProductImgDelete) ExecX(ctx context.Context) int {
	n, err := pid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pid *ProductImgDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(productimg.Table, sqlgraph.NewFieldSpec(productimg.FieldID, field.TypeUint64))
	if ps := pid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pid.mutation.done = true
	return affected, err
}

// ProductImgDeleteOne is the builder for deleting a single ProductImg entity.
type ProductImgDeleteOne struct {
	pid *ProductImgDelete
}

// Where appends a list predicates to the ProductImgDelete builder.
func (pido *ProductImgDeleteOne) Where(ps ...predicate.ProductImg) *ProductImgDeleteOne {
	pido.pid.mutation.Where(ps...)
	return pido
}

// Exec executes the deletion query.
func (pido *ProductImgDeleteOne) Exec(ctx context.Context) error {
	n, err := pido.pid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productimg.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pido *ProductImgDeleteOne) ExecX(ctx context.Context) {
	if err := pido.Exec(ctx); err != nil {
		panic(err)
	}
}
