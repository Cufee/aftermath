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
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/clan"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(i int64) *AccountCreate {
	ac.mutation.SetCreatedAt(i)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(i *int64) *AccountCreate {
	if i != nil {
		ac.SetCreatedAt(*i)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(i int64) *AccountCreate {
	ac.mutation.SetUpdatedAt(i)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(i *int64) *AccountCreate {
	if i != nil {
		ac.SetUpdatedAt(*i)
	}
	return ac
}

// SetLastBattleTime sets the "last_battle_time" field.
func (ac *AccountCreate) SetLastBattleTime(i int64) *AccountCreate {
	ac.mutation.SetLastBattleTime(i)
	return ac
}

// SetAccountCreatedAt sets the "account_created_at" field.
func (ac *AccountCreate) SetAccountCreatedAt(i int64) *AccountCreate {
	ac.mutation.SetAccountCreatedAt(i)
	return ac
}

// SetRealm sets the "realm" field.
func (ac *AccountCreate) SetRealm(s string) *AccountCreate {
	ac.mutation.SetRealm(s)
	return ac
}

// SetNickname sets the "nickname" field.
func (ac *AccountCreate) SetNickname(s string) *AccountCreate {
	ac.mutation.SetNickname(s)
	return ac
}

// SetPrivate sets the "private" field.
func (ac *AccountCreate) SetPrivate(b bool) *AccountCreate {
	ac.mutation.SetPrivate(b)
	return ac
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (ac *AccountCreate) SetNillablePrivate(b *bool) *AccountCreate {
	if b != nil {
		ac.SetPrivate(*b)
	}
	return ac
}

// SetClanID sets the "clan_id" field.
func (ac *AccountCreate) SetClanID(s string) *AccountCreate {
	ac.mutation.SetClanID(s)
	return ac
}

// SetNillableClanID sets the "clan_id" field if the given value is not nil.
func (ac *AccountCreate) SetNillableClanID(s *string) *AccountCreate {
	if s != nil {
		ac.SetClanID(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(s string) *AccountCreate {
	ac.mutation.SetID(s)
	return ac
}

// SetClan sets the "clan" edge to the Clan entity.
func (ac *AccountCreate) SetClan(c *Clan) *AccountCreate {
	return ac.SetClanID(c.ID)
}

// AddSnapshotIDs adds the "snapshots" edge to the AccountSnapshot entity by IDs.
func (ac *AccountCreate) AddSnapshotIDs(ids ...string) *AccountCreate {
	ac.mutation.AddSnapshotIDs(ids...)
	return ac
}

// AddSnapshots adds the "snapshots" edges to the AccountSnapshot entity.
func (ac *AccountCreate) AddSnapshots(a ...*AccountSnapshot) *AccountCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddSnapshotIDs(ids...)
}

// AddVehicleSnapshotIDs adds the "vehicle_snapshots" edge to the VehicleSnapshot entity by IDs.
func (ac *AccountCreate) AddVehicleSnapshotIDs(ids ...string) *AccountCreate {
	ac.mutation.AddVehicleSnapshotIDs(ids...)
	return ac
}

// AddVehicleSnapshots adds the "vehicle_snapshots" edges to the VehicleSnapshot entity.
func (ac *AccountCreate) AddVehicleSnapshots(v ...*VehicleSnapshot) *AccountCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ac.AddVehicleSnapshotIDs(ids...)
}

// AddAchievementSnapshotIDs adds the "achievement_snapshots" edge to the AchievementsSnapshot entity by IDs.
func (ac *AccountCreate) AddAchievementSnapshotIDs(ids ...string) *AccountCreate {
	ac.mutation.AddAchievementSnapshotIDs(ids...)
	return ac
}

// AddAchievementSnapshots adds the "achievement_snapshots" edges to the AchievementsSnapshot entity.
func (ac *AccountCreate) AddAchievementSnapshots(a ...*AchievementsSnapshot) *AccountCreate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ac.AddAchievementSnapshotIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Private(); !ok {
		v := account.DefaultPrivate
		ac.mutation.SetPrivate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "Account.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "Account.updated_at"`)}
	}
	if _, ok := ac.mutation.LastBattleTime(); !ok {
		return &ValidationError{Name: "last_battle_time", err: errors.New(`db: missing required field "Account.last_battle_time"`)}
	}
	if _, ok := ac.mutation.AccountCreatedAt(); !ok {
		return &ValidationError{Name: "account_created_at", err: errors.New(`db: missing required field "Account.account_created_at"`)}
	}
	if _, ok := ac.mutation.Realm(); !ok {
		return &ValidationError{Name: "realm", err: errors.New(`db: missing required field "Account.realm"`)}
	}
	if v, ok := ac.mutation.Realm(); ok {
		if err := account.RealmValidator(v); err != nil {
			return &ValidationError{Name: "realm", err: fmt.Errorf(`db: validator failed for field "Account.realm": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`db: missing required field "Account.nickname"`)}
	}
	if v, ok := ac.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`db: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Private(); !ok {
		return &ValidationError{Name: "private", err: errors.New(`db: missing required field "Account.private"`)}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Account.ID type: %T", _spec.ID.Value)
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeString))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.LastBattleTime(); ok {
		_spec.SetField(account.FieldLastBattleTime, field.TypeInt64, value)
		_node.LastBattleTime = value
	}
	if value, ok := ac.mutation.AccountCreatedAt(); ok {
		_spec.SetField(account.FieldAccountCreatedAt, field.TypeInt64, value)
		_node.AccountCreatedAt = value
	}
	if value, ok := ac.mutation.Realm(); ok {
		_spec.SetField(account.FieldRealm, field.TypeString, value)
		_node.Realm = value
	}
	if value, ok := ac.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := ac.mutation.Private(); ok {
		_spec.SetField(account.FieldPrivate, field.TypeBool, value)
		_node.Private = value
	}
	if nodes := ac.mutation.ClanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   account.ClanTable,
			Columns: []string{account.ClanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(clan.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClanID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.SnapshotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.SnapshotsTable,
			Columns: []string{account.SnapshotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accountsnapshot.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.VehicleSnapshotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.VehicleSnapshotsTable,
			Columns: []string{account.VehicleSnapshotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(vehiclesnapshot.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.AchievementSnapshotsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   account.AchievementSnapshotsTable,
			Columns: []string{account.AchievementSnapshotsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(achievementssnapshot.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}