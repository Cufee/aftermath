// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// VehicleSnapshotCreate is the builder for creating a VehicleSnapshot entity.
type VehicleSnapshotCreate struct {
	config
	mutation *VehicleSnapshotMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vsc *VehicleSnapshotCreate) SetCreatedAt(t time.Time) *VehicleSnapshotCreate {
	vsc.mutation.SetCreatedAt(t)
	return vsc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vsc *VehicleSnapshotCreate) SetNillableCreatedAt(t *time.Time) *VehicleSnapshotCreate {
	if t != nil {
		vsc.SetCreatedAt(*t)
	}
	return vsc
}

// SetUpdatedAt sets the "updated_at" field.
func (vsc *VehicleSnapshotCreate) SetUpdatedAt(t time.Time) *VehicleSnapshotCreate {
	vsc.mutation.SetUpdatedAt(t)
	return vsc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vsc *VehicleSnapshotCreate) SetNillableUpdatedAt(t *time.Time) *VehicleSnapshotCreate {
	if t != nil {
		vsc.SetUpdatedAt(*t)
	}
	return vsc
}

// SetType sets the "type" field.
func (vsc *VehicleSnapshotCreate) SetType(mt models.SnapshotType) *VehicleSnapshotCreate {
	vsc.mutation.SetType(mt)
	return vsc
}

// SetAccountID sets the "account_id" field.
func (vsc *VehicleSnapshotCreate) SetAccountID(s string) *VehicleSnapshotCreate {
	vsc.mutation.SetAccountID(s)
	return vsc
}

// SetVehicleID sets the "vehicle_id" field.
func (vsc *VehicleSnapshotCreate) SetVehicleID(s string) *VehicleSnapshotCreate {
	vsc.mutation.SetVehicleID(s)
	return vsc
}

// SetReferenceID sets the "reference_id" field.
func (vsc *VehicleSnapshotCreate) SetReferenceID(s string) *VehicleSnapshotCreate {
	vsc.mutation.SetReferenceID(s)
	return vsc
}

// SetBattles sets the "battles" field.
func (vsc *VehicleSnapshotCreate) SetBattles(i int) *VehicleSnapshotCreate {
	vsc.mutation.SetBattles(i)
	return vsc
}

// SetLastBattleTime sets the "last_battle_time" field.
func (vsc *VehicleSnapshotCreate) SetLastBattleTime(t time.Time) *VehicleSnapshotCreate {
	vsc.mutation.SetLastBattleTime(t)
	return vsc
}

// SetFrame sets the "frame" field.
func (vsc *VehicleSnapshotCreate) SetFrame(ff frame.StatsFrame) *VehicleSnapshotCreate {
	vsc.mutation.SetFrame(ff)
	return vsc
}

// SetID sets the "id" field.
func (vsc *VehicleSnapshotCreate) SetID(s string) *VehicleSnapshotCreate {
	vsc.mutation.SetID(s)
	return vsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (vsc *VehicleSnapshotCreate) SetNillableID(s *string) *VehicleSnapshotCreate {
	if s != nil {
		vsc.SetID(*s)
	}
	return vsc
}

// SetAccount sets the "account" edge to the Account entity.
func (vsc *VehicleSnapshotCreate) SetAccount(a *Account) *VehicleSnapshotCreate {
	return vsc.SetAccountID(a.ID)
}

// Mutation returns the VehicleSnapshotMutation object of the builder.
func (vsc *VehicleSnapshotCreate) Mutation() *VehicleSnapshotMutation {
	return vsc.mutation
}

// Save creates the VehicleSnapshot in the database.
func (vsc *VehicleSnapshotCreate) Save(ctx context.Context) (*VehicleSnapshot, error) {
	vsc.defaults()
	return withHooks(ctx, vsc.sqlSave, vsc.mutation, vsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vsc *VehicleSnapshotCreate) SaveX(ctx context.Context) *VehicleSnapshot {
	v, err := vsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vsc *VehicleSnapshotCreate) Exec(ctx context.Context) error {
	_, err := vsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vsc *VehicleSnapshotCreate) ExecX(ctx context.Context) {
	if err := vsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vsc *VehicleSnapshotCreate) defaults() {
	if _, ok := vsc.mutation.CreatedAt(); !ok {
		v := vehiclesnapshot.DefaultCreatedAt()
		vsc.mutation.SetCreatedAt(v)
	}
	if _, ok := vsc.mutation.UpdatedAt(); !ok {
		v := vehiclesnapshot.DefaultUpdatedAt()
		vsc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vsc.mutation.ID(); !ok {
		v := vehiclesnapshot.DefaultID()
		vsc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vsc *VehicleSnapshotCreate) check() error {
	if _, ok := vsc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "VehicleSnapshot.created_at"`)}
	}
	if _, ok := vsc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "VehicleSnapshot.updated_at"`)}
	}
	if _, ok := vsc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "VehicleSnapshot.type"`)}
	}
	if v, ok := vsc.mutation.GetType(); ok {
		if err := vehiclesnapshot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "VehicleSnapshot.type": %w`, err)}
		}
	}
	if _, ok := vsc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`db: missing required field "VehicleSnapshot.account_id"`)}
	}
	if v, ok := vsc.mutation.AccountID(); ok {
		if err := vehiclesnapshot.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`db: validator failed for field "VehicleSnapshot.account_id": %w`, err)}
		}
	}
	if _, ok := vsc.mutation.VehicleID(); !ok {
		return &ValidationError{Name: "vehicle_id", err: errors.New(`db: missing required field "VehicleSnapshot.vehicle_id"`)}
	}
	if v, ok := vsc.mutation.VehicleID(); ok {
		if err := vehiclesnapshot.VehicleIDValidator(v); err != nil {
			return &ValidationError{Name: "vehicle_id", err: fmt.Errorf(`db: validator failed for field "VehicleSnapshot.vehicle_id": %w`, err)}
		}
	}
	if _, ok := vsc.mutation.ReferenceID(); !ok {
		return &ValidationError{Name: "reference_id", err: errors.New(`db: missing required field "VehicleSnapshot.reference_id"`)}
	}
	if v, ok := vsc.mutation.ReferenceID(); ok {
		if err := vehiclesnapshot.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "VehicleSnapshot.reference_id": %w`, err)}
		}
	}
	if _, ok := vsc.mutation.Battles(); !ok {
		return &ValidationError{Name: "battles", err: errors.New(`db: missing required field "VehicleSnapshot.battles"`)}
	}
	if _, ok := vsc.mutation.LastBattleTime(); !ok {
		return &ValidationError{Name: "last_battle_time", err: errors.New(`db: missing required field "VehicleSnapshot.last_battle_time"`)}
	}
	if _, ok := vsc.mutation.Frame(); !ok {
		return &ValidationError{Name: "frame", err: errors.New(`db: missing required field "VehicleSnapshot.frame"`)}
	}
	if _, ok := vsc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`db: missing required edge "VehicleSnapshot.account"`)}
	}
	return nil
}

func (vsc *VehicleSnapshotCreate) sqlSave(ctx context.Context) (*VehicleSnapshot, error) {
	if err := vsc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected VehicleSnapshot.ID type: %T", _spec.ID.Value)
		}
	}
	vsc.mutation.id = &_node.ID
	vsc.mutation.done = true
	return _node, nil
}

func (vsc *VehicleSnapshotCreate) createSpec() (*VehicleSnapshot, *sqlgraph.CreateSpec) {
	var (
		_node = &VehicleSnapshot{config: vsc.config}
		_spec = sqlgraph.NewCreateSpec(vehiclesnapshot.Table, sqlgraph.NewFieldSpec(vehiclesnapshot.FieldID, field.TypeString))
	)
	if id, ok := vsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vsc.mutation.CreatedAt(); ok {
		_spec.SetField(vehiclesnapshot.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vsc.mutation.UpdatedAt(); ok {
		_spec.SetField(vehiclesnapshot.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vsc.mutation.GetType(); ok {
		_spec.SetField(vehiclesnapshot.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := vsc.mutation.VehicleID(); ok {
		_spec.SetField(vehiclesnapshot.FieldVehicleID, field.TypeString, value)
		_node.VehicleID = value
	}
	if value, ok := vsc.mutation.ReferenceID(); ok {
		_spec.SetField(vehiclesnapshot.FieldReferenceID, field.TypeString, value)
		_node.ReferenceID = value
	}
	if value, ok := vsc.mutation.Battles(); ok {
		_spec.SetField(vehiclesnapshot.FieldBattles, field.TypeInt, value)
		_node.Battles = value
	}
	if value, ok := vsc.mutation.LastBattleTime(); ok {
		_spec.SetField(vehiclesnapshot.FieldLastBattleTime, field.TypeTime, value)
		_node.LastBattleTime = value
	}
	if value, ok := vsc.mutation.Frame(); ok {
		_spec.SetField(vehiclesnapshot.FieldFrame, field.TypeJSON, value)
		_node.Frame = value
	}
	if nodes := vsc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vehiclesnapshot.AccountTable,
			Columns: []string{vehiclesnapshot.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VehicleSnapshotCreateBulk is the builder for creating many VehicleSnapshot entities in bulk.
type VehicleSnapshotCreateBulk struct {
	config
	err      error
	builders []*VehicleSnapshotCreate
}

// Save creates the VehicleSnapshot entities in the database.
func (vscb *VehicleSnapshotCreateBulk) Save(ctx context.Context) ([]*VehicleSnapshot, error) {
	if vscb.err != nil {
		return nil, vscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vscb.builders))
	nodes := make([]*VehicleSnapshot, len(vscb.builders))
	mutators := make([]Mutator, len(vscb.builders))
	for i := range vscb.builders {
		func(i int, root context.Context) {
			builder := vscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VehicleSnapshotMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vscb *VehicleSnapshotCreateBulk) SaveX(ctx context.Context) []*VehicleSnapshot {
	v, err := vscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vscb *VehicleSnapshotCreateBulk) Exec(ctx context.Context) error {
	_, err := vscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vscb *VehicleSnapshotCreateBulk) ExecX(ctx context.Context) {
	if err := vscb.Exec(ctx); err != nil {
		panic(err)
	}
}
