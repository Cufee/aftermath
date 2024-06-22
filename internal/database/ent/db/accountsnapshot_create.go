// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// AccountSnapshotCreate is the builder for creating a AccountSnapshot entity.
type AccountSnapshotCreate struct {
	config
	mutation *AccountSnapshotMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (asc *AccountSnapshotCreate) SetCreatedAt(i int) *AccountSnapshotCreate {
	asc.mutation.SetCreatedAt(i)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *AccountSnapshotCreate) SetNillableCreatedAt(i *int) *AccountSnapshotCreate {
	if i != nil {
		asc.SetCreatedAt(*i)
	}
	return asc
}

// SetUpdatedAt sets the "updated_at" field.
func (asc *AccountSnapshotCreate) SetUpdatedAt(i int) *AccountSnapshotCreate {
	asc.mutation.SetUpdatedAt(i)
	return asc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (asc *AccountSnapshotCreate) SetNillableUpdatedAt(i *int) *AccountSnapshotCreate {
	if i != nil {
		asc.SetUpdatedAt(*i)
	}
	return asc
}

// SetType sets the "type" field.
func (asc *AccountSnapshotCreate) SetType(mt models.SnapshotType) *AccountSnapshotCreate {
	asc.mutation.SetType(mt)
	return asc
}

// SetLastBattleTime sets the "last_battle_time" field.
func (asc *AccountSnapshotCreate) SetLastBattleTime(i int) *AccountSnapshotCreate {
	asc.mutation.SetLastBattleTime(i)
	return asc
}

// SetAccountID sets the "account_id" field.
func (asc *AccountSnapshotCreate) SetAccountID(s string) *AccountSnapshotCreate {
	asc.mutation.SetAccountID(s)
	return asc
}

// SetReferenceID sets the "reference_id" field.
func (asc *AccountSnapshotCreate) SetReferenceID(s string) *AccountSnapshotCreate {
	asc.mutation.SetReferenceID(s)
	return asc
}

// SetRatingBattles sets the "rating_battles" field.
func (asc *AccountSnapshotCreate) SetRatingBattles(i int) *AccountSnapshotCreate {
	asc.mutation.SetRatingBattles(i)
	return asc
}

// SetRatingFrame sets the "rating_frame" field.
func (asc *AccountSnapshotCreate) SetRatingFrame(ff frame.StatsFrame) *AccountSnapshotCreate {
	asc.mutation.SetRatingFrame(ff)
	return asc
}

// SetRegularBattles sets the "regular_battles" field.
func (asc *AccountSnapshotCreate) SetRegularBattles(i int) *AccountSnapshotCreate {
	asc.mutation.SetRegularBattles(i)
	return asc
}

// SetRegularFrame sets the "regular_frame" field.
func (asc *AccountSnapshotCreate) SetRegularFrame(ff frame.StatsFrame) *AccountSnapshotCreate {
	asc.mutation.SetRegularFrame(ff)
	return asc
}

// SetID sets the "id" field.
func (asc *AccountSnapshotCreate) SetID(s string) *AccountSnapshotCreate {
	asc.mutation.SetID(s)
	return asc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (asc *AccountSnapshotCreate) SetNillableID(s *string) *AccountSnapshotCreate {
	if s != nil {
		asc.SetID(*s)
	}
	return asc
}

// SetAccount sets the "account" edge to the Account entity.
func (asc *AccountSnapshotCreate) SetAccount(a *Account) *AccountSnapshotCreate {
	return asc.SetAccountID(a.ID)
}

// Mutation returns the AccountSnapshotMutation object of the builder.
func (asc *AccountSnapshotCreate) Mutation() *AccountSnapshotMutation {
	return asc.mutation
}

// Save creates the AccountSnapshot in the database.
func (asc *AccountSnapshotCreate) Save(ctx context.Context) (*AccountSnapshot, error) {
	asc.defaults()
	return withHooks(ctx, asc.sqlSave, asc.mutation, asc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AccountSnapshotCreate) SaveX(ctx context.Context) *AccountSnapshot {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AccountSnapshotCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AccountSnapshotCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *AccountSnapshotCreate) defaults() {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		v := accountsnapshot.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		v := accountsnapshot.DefaultUpdatedAt()
		asc.mutation.SetUpdatedAt(v)
	}
	if _, ok := asc.mutation.ID(); !ok {
		v := accountsnapshot.DefaultID()
		asc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AccountSnapshotCreate) check() error {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "AccountSnapshot.created_at"`)}
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "AccountSnapshot.updated_at"`)}
	}
	if _, ok := asc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "AccountSnapshot.type"`)}
	}
	if v, ok := asc.mutation.GetType(); ok {
		if err := accountsnapshot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.type": %w`, err)}
		}
	}
	if _, ok := asc.mutation.LastBattleTime(); !ok {
		return &ValidationError{Name: "last_battle_time", err: errors.New(`db: missing required field "AccountSnapshot.last_battle_time"`)}
	}
	if _, ok := asc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`db: missing required field "AccountSnapshot.account_id"`)}
	}
	if v, ok := asc.mutation.AccountID(); ok {
		if err := accountsnapshot.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.account_id": %w`, err)}
		}
	}
	if _, ok := asc.mutation.ReferenceID(); !ok {
		return &ValidationError{Name: "reference_id", err: errors.New(`db: missing required field "AccountSnapshot.reference_id"`)}
	}
	if v, ok := asc.mutation.ReferenceID(); ok {
		if err := accountsnapshot.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.reference_id": %w`, err)}
		}
	}
	if _, ok := asc.mutation.RatingBattles(); !ok {
		return &ValidationError{Name: "rating_battles", err: errors.New(`db: missing required field "AccountSnapshot.rating_battles"`)}
	}
	if _, ok := asc.mutation.RatingFrame(); !ok {
		return &ValidationError{Name: "rating_frame", err: errors.New(`db: missing required field "AccountSnapshot.rating_frame"`)}
	}
	if _, ok := asc.mutation.RegularBattles(); !ok {
		return &ValidationError{Name: "regular_battles", err: errors.New(`db: missing required field "AccountSnapshot.regular_battles"`)}
	}
	if _, ok := asc.mutation.RegularFrame(); !ok {
		return &ValidationError{Name: "regular_frame", err: errors.New(`db: missing required field "AccountSnapshot.regular_frame"`)}
	}
	if _, ok := asc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`db: missing required edge "AccountSnapshot.account"`)}
	}
	return nil
}

func (asc *AccountSnapshotCreate) sqlSave(ctx context.Context) (*AccountSnapshot, error) {
	if err := asc.check(); err != nil {
		return nil, err
	}
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected AccountSnapshot.ID type: %T", _spec.ID.Value)
		}
	}
	asc.mutation.id = &_node.ID
	asc.mutation.done = true
	return _node, nil
}

func (asc *AccountSnapshotCreate) createSpec() (*AccountSnapshot, *sqlgraph.CreateSpec) {
	var (
		_node = &AccountSnapshot{config: asc.config}
		_spec = sqlgraph.NewCreateSpec(accountsnapshot.Table, sqlgraph.NewFieldSpec(accountsnapshot.FieldID, field.TypeString))
	)
	if id, ok := asc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.SetField(accountsnapshot.FieldCreatedAt, field.TypeInt, value)
		_node.CreatedAt = value
	}
	if value, ok := asc.mutation.UpdatedAt(); ok {
		_spec.SetField(accountsnapshot.FieldUpdatedAt, field.TypeInt, value)
		_node.UpdatedAt = value
	}
	if value, ok := asc.mutation.GetType(); ok {
		_spec.SetField(accountsnapshot.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := asc.mutation.LastBattleTime(); ok {
		_spec.SetField(accountsnapshot.FieldLastBattleTime, field.TypeInt, value)
		_node.LastBattleTime = value
	}
	if value, ok := asc.mutation.ReferenceID(); ok {
		_spec.SetField(accountsnapshot.FieldReferenceID, field.TypeString, value)
		_node.ReferenceID = value
	}
	if value, ok := asc.mutation.RatingBattles(); ok {
		_spec.SetField(accountsnapshot.FieldRatingBattles, field.TypeInt, value)
		_node.RatingBattles = value
	}
	if value, ok := asc.mutation.RatingFrame(); ok {
		_spec.SetField(accountsnapshot.FieldRatingFrame, field.TypeJSON, value)
		_node.RatingFrame = value
	}
	if value, ok := asc.mutation.RegularBattles(); ok {
		_spec.SetField(accountsnapshot.FieldRegularBattles, field.TypeInt, value)
		_node.RegularBattles = value
	}
	if value, ok := asc.mutation.RegularFrame(); ok {
		_spec.SetField(accountsnapshot.FieldRegularFrame, field.TypeJSON, value)
		_node.RegularFrame = value
	}
	if nodes := asc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accountsnapshot.AccountTable,
			Columns: []string{accountsnapshot.AccountColumn},
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

// AccountSnapshotCreateBulk is the builder for creating many AccountSnapshot entities in bulk.
type AccountSnapshotCreateBulk struct {
	config
	err      error
	builders []*AccountSnapshotCreate
}

// Save creates the AccountSnapshot entities in the database.
func (ascb *AccountSnapshotCreateBulk) Save(ctx context.Context) ([]*AccountSnapshot, error) {
	if ascb.err != nil {
		return nil, ascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AccountSnapshot, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountSnapshotMutation)
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
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AccountSnapshotCreateBulk) SaveX(ctx context.Context) []*AccountSnapshot {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AccountSnapshotCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AccountSnapshotCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
