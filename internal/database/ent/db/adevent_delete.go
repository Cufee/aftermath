// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/adevent"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// AdEventDelete is the builder for deleting a AdEvent entity.
type AdEventDelete struct {
	config
	hooks    []Hook
	mutation *AdEventMutation
}

// Where appends a list predicates to the AdEventDelete builder.
func (aed *AdEventDelete) Where(ps ...predicate.AdEvent) *AdEventDelete {
	aed.mutation.Where(ps...)
	return aed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aed *AdEventDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, aed.sqlExec, aed.mutation, aed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (aed *AdEventDelete) ExecX(ctx context.Context) int {
	n, err := aed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aed *AdEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(adevent.Table, sqlgraph.NewFieldSpec(adevent.FieldID, field.TypeString))
	if ps := aed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	aed.mutation.done = true
	return affected, err
}

// AdEventDeleteOne is the builder for deleting a single AdEvent entity.
type AdEventDeleteOne struct {
	aed *AdEventDelete
}

// Where appends a list predicates to the AdEventDelete builder.
func (aedo *AdEventDeleteOne) Where(ps ...predicate.AdEvent) *AdEventDeleteOne {
	aedo.aed.mutation.Where(ps...)
	return aedo
}

// Exec executes the deletion query.
func (aedo *AdEventDeleteOne) Exec(ctx context.Context) error {
	n, err := aedo.aed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{adevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aedo *AdEventDeleteOne) ExecX(ctx context.Context) {
	if err := aedo.Exec(ctx); err != nil {
		panic(err)
	}
}