// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicle"
)

// VehicleDelete is the builder for deleting a Vehicle entity.
type VehicleDelete struct {
	config
	hooks    []Hook
	mutation *VehicleMutation
}

// Where appends a list predicates to the VehicleDelete builder.
func (vd *VehicleDelete) Where(ps ...predicate.Vehicle) *VehicleDelete {
	vd.mutation.Where(ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VehicleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vd.sqlExec, vd.mutation, vd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VehicleDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VehicleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vehicle.Table, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeString))
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vd.mutation.done = true
	return affected, err
}

// VehicleDeleteOne is the builder for deleting a single Vehicle entity.
type VehicleDeleteOne struct {
	vd *VehicleDelete
}

// Where appends a list predicates to the VehicleDelete builder.
func (vdo *VehicleDeleteOne) Where(ps ...predicate.Vehicle) *VehicleDeleteOne {
	vdo.vd.mutation.Where(ps...)
	return vdo
}

// Exec executes the deletion query.
func (vdo *VehicleDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vehicle.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VehicleDeleteOne) ExecX(ctx context.Context) {
	if err := vdo.Exec(ctx); err != nil {
		panic(err)
	}
}
