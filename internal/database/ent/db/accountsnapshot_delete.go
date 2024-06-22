// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// AccountSnapshotDelete is the builder for deleting a AccountSnapshot entity.
type AccountSnapshotDelete struct {
	config
	hooks    []Hook
	mutation *AccountSnapshotMutation
}

// Where appends a list predicates to the AccountSnapshotDelete builder.
func (asd *AccountSnapshotDelete) Where(ps ...predicate.AccountSnapshot) *AccountSnapshotDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AccountSnapshotDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AccountSnapshotDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AccountSnapshotDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(accountsnapshot.Table, sqlgraph.NewFieldSpec(accountsnapshot.FieldID, field.TypeString))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// AccountSnapshotDeleteOne is the builder for deleting a single AccountSnapshot entity.
type AccountSnapshotDeleteOne struct {
	asd *AccountSnapshotDelete
}

// Where appends a list predicates to the AccountSnapshotDelete builder.
func (asdo *AccountSnapshotDeleteOne) Where(ps ...predicate.AccountSnapshot) *AccountSnapshotDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *AccountSnapshotDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{accountsnapshot.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AccountSnapshotDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}
