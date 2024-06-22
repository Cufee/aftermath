// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/userconnection"
)

// UserConnectionDelete is the builder for deleting a UserConnection entity.
type UserConnectionDelete struct {
	config
	hooks    []Hook
	mutation *UserConnectionMutation
}

// Where appends a list predicates to the UserConnectionDelete builder.
func (ucd *UserConnectionDelete) Where(ps ...predicate.UserConnection) *UserConnectionDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserConnectionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ucd.sqlExec, ucd.mutation, ucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserConnectionDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserConnectionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userconnection.Table, sqlgraph.NewFieldSpec(userconnection.FieldID, field.TypeString))
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucd.mutation.done = true
	return affected, err
}

// UserConnectionDeleteOne is the builder for deleting a single UserConnection entity.
type UserConnectionDeleteOne struct {
	ucd *UserConnectionDelete
}

// Where appends a list predicates to the UserConnectionDelete builder.
func (ucdo *UserConnectionDeleteOne) Where(ps ...predicate.UserConnection) *UserConnectionDeleteOne {
	ucdo.ucd.mutation.Where(ps...)
	return ucdo
}

// Exec executes the deletion query.
func (ucdo *UserConnectionDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userconnection.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserConnectionDeleteOne) ExecX(ctx context.Context) {
	if err := ucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
