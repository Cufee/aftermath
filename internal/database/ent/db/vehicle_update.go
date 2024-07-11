// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/vehicle"
	"golang.org/x/text/language"
)

// VehicleUpdate is the builder for updating Vehicle entities.
type VehicleUpdate struct {
	config
	hooks     []Hook
	mutation  *VehicleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the VehicleUpdate builder.
func (vu *VehicleUpdate) Where(ps ...predicate.Vehicle) *VehicleUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetUpdatedAt sets the "updated_at" field.
func (vu *VehicleUpdate) SetUpdatedAt(t time.Time) *VehicleUpdate {
	vu.mutation.SetUpdatedAt(t)
	return vu
}

// SetTier sets the "tier" field.
func (vu *VehicleUpdate) SetTier(i int) *VehicleUpdate {
	vu.mutation.ResetTier()
	vu.mutation.SetTier(i)
	return vu
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (vu *VehicleUpdate) SetNillableTier(i *int) *VehicleUpdate {
	if i != nil {
		vu.SetTier(*i)
	}
	return vu
}

// AddTier adds i to the "tier" field.
func (vu *VehicleUpdate) AddTier(i int) *VehicleUpdate {
	vu.mutation.AddTier(i)
	return vu
}

// SetLocalizedNames sets the "localized_names" field.
func (vu *VehicleUpdate) SetLocalizedNames(m map[language.Tag]string) *VehicleUpdate {
	vu.mutation.SetLocalizedNames(m)
	return vu
}

// Mutation returns the VehicleMutation object of the builder.
func (vu *VehicleUpdate) Mutation() *VehicleMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VehicleUpdate) Save(ctx context.Context) (int, error) {
	vu.defaults()
	return withHooks(ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VehicleUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VehicleUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VehicleUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vu *VehicleUpdate) defaults() {
	if _, ok := vu.mutation.UpdatedAt(); !ok {
		v := vehicle.UpdateDefaultUpdatedAt()
		vu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vu *VehicleUpdate) check() error {
	if v, ok := vu.mutation.Tier(); ok {
		if err := vehicle.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`db: validator failed for field "Vehicle.tier": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vu *VehicleUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VehicleUpdate {
	vu.modifiers = append(vu.modifiers, modifiers...)
	return vu
}

func (vu *VehicleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(vehicle.Table, vehicle.Columns, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeString))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.UpdatedAt(); ok {
		_spec.SetField(vehicle.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vu.mutation.Tier(); ok {
		_spec.SetField(vehicle.FieldTier, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedTier(); ok {
		_spec.AddField(vehicle.FieldTier, field.TypeInt, value)
	}
	if value, ok := vu.mutation.LocalizedNames(); ok {
		_spec.SetField(vehicle.FieldLocalizedNames, field.TypeJSON, value)
	}
	_spec.AddModifiers(vu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VehicleUpdateOne is the builder for updating a single Vehicle entity.
type VehicleUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *VehicleMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (vuo *VehicleUpdateOne) SetUpdatedAt(t time.Time) *VehicleUpdateOne {
	vuo.mutation.SetUpdatedAt(t)
	return vuo
}

// SetTier sets the "tier" field.
func (vuo *VehicleUpdateOne) SetTier(i int) *VehicleUpdateOne {
	vuo.mutation.ResetTier()
	vuo.mutation.SetTier(i)
	return vuo
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (vuo *VehicleUpdateOne) SetNillableTier(i *int) *VehicleUpdateOne {
	if i != nil {
		vuo.SetTier(*i)
	}
	return vuo
}

// AddTier adds i to the "tier" field.
func (vuo *VehicleUpdateOne) AddTier(i int) *VehicleUpdateOne {
	vuo.mutation.AddTier(i)
	return vuo
}

// SetLocalizedNames sets the "localized_names" field.
func (vuo *VehicleUpdateOne) SetLocalizedNames(m map[language.Tag]string) *VehicleUpdateOne {
	vuo.mutation.SetLocalizedNames(m)
	return vuo
}

// Mutation returns the VehicleMutation object of the builder.
func (vuo *VehicleUpdateOne) Mutation() *VehicleMutation {
	return vuo.mutation
}

// Where appends a list predicates to the VehicleUpdate builder.
func (vuo *VehicleUpdateOne) Where(ps ...predicate.Vehicle) *VehicleUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VehicleUpdateOne) Select(field string, fields ...string) *VehicleUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Vehicle entity.
func (vuo *VehicleUpdateOne) Save(ctx context.Context) (*Vehicle, error) {
	vuo.defaults()
	return withHooks(ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VehicleUpdateOne) SaveX(ctx context.Context) *Vehicle {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VehicleUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VehicleUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vuo *VehicleUpdateOne) defaults() {
	if _, ok := vuo.mutation.UpdatedAt(); !ok {
		v := vehicle.UpdateDefaultUpdatedAt()
		vuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuo *VehicleUpdateOne) check() error {
	if v, ok := vuo.mutation.Tier(); ok {
		if err := vehicle.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`db: validator failed for field "Vehicle.tier": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vuo *VehicleUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VehicleUpdateOne {
	vuo.modifiers = append(vuo.modifiers, modifiers...)
	return vuo
}

func (vuo *VehicleUpdateOne) sqlSave(ctx context.Context) (_node *Vehicle, err error) {
	if err := vuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(vehicle.Table, vehicle.Columns, sqlgraph.NewFieldSpec(vehicle.FieldID, field.TypeString))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Vehicle.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vehicle.FieldID)
		for _, f := range fields {
			if !vehicle.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != vehicle.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.UpdatedAt(); ok {
		_spec.SetField(vehicle.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vuo.mutation.Tier(); ok {
		_spec.SetField(vehicle.FieldTier, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedTier(); ok {
		_spec.AddField(vehicle.FieldTier, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.LocalizedNames(); ok {
		_spec.SetField(vehicle.FieldLocalizedNames, field.TypeJSON, value)
	}
	_spec.AddModifiers(vuo.modifiers...)
	_node = &Vehicle{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vehicle.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
