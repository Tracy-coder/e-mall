// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tracy-coder/e-mall/data/ent/favourite"
	"github.com/Tracy-coder/e-mall/data/ent/predicate"
)

// FavouriteDelete is the builder for deleting a Favourite entity.
type FavouriteDelete struct {
	config
	hooks    []Hook
	mutation *FavouriteMutation
}

// Where appends a list predicates to the FavouriteDelete builder.
func (fd *FavouriteDelete) Where(ps ...predicate.Favourite) *FavouriteDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FavouriteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FavouriteDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FavouriteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(favourite.Table, sqlgraph.NewFieldSpec(favourite.FieldID, field.TypeUint64))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FavouriteDeleteOne is the builder for deleting a single Favourite entity.
type FavouriteDeleteOne struct {
	fd *FavouriteDelete
}

// Where appends a list predicates to the FavouriteDelete builder.
func (fdo *FavouriteDeleteOne) Where(ps ...predicate.Favourite) *FavouriteDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FavouriteDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{favourite.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FavouriteDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
