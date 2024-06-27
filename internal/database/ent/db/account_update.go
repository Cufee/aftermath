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
	"github.com/cufee/aftermath/internal/database/ent/db/account"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/clan"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/vehiclesnapshot"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks     []Hook
	mutation  *AccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetLastBattleTime sets the "last_battle_time" field.
func (au *AccountUpdate) SetLastBattleTime(t time.Time) *AccountUpdate {
	au.mutation.SetLastBattleTime(t)
	return au
}

// SetNillableLastBattleTime sets the "last_battle_time" field if the given value is not nil.
func (au *AccountUpdate) SetNillableLastBattleTime(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetLastBattleTime(*t)
	}
	return au
}

// SetAccountCreatedAt sets the "account_created_at" field.
func (au *AccountUpdate) SetAccountCreatedAt(t time.Time) *AccountUpdate {
	au.mutation.SetAccountCreatedAt(t)
	return au
}

// SetNillableAccountCreatedAt sets the "account_created_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAccountCreatedAt(t *time.Time) *AccountUpdate {
	if t != nil {
		au.SetAccountCreatedAt(*t)
	}
	return au
}

// SetRealm sets the "realm" field.
func (au *AccountUpdate) SetRealm(s string) *AccountUpdate {
	au.mutation.SetRealm(s)
	return au
}

// SetNillableRealm sets the "realm" field if the given value is not nil.
func (au *AccountUpdate) SetNillableRealm(s *string) *AccountUpdate {
	if s != nil {
		au.SetRealm(*s)
	}
	return au
}

// SetNickname sets the "nickname" field.
func (au *AccountUpdate) SetNickname(s string) *AccountUpdate {
	au.mutation.SetNickname(s)
	return au
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (au *AccountUpdate) SetNillableNickname(s *string) *AccountUpdate {
	if s != nil {
		au.SetNickname(*s)
	}
	return au
}

// SetPrivate sets the "private" field.
func (au *AccountUpdate) SetPrivate(b bool) *AccountUpdate {
	au.mutation.SetPrivate(b)
	return au
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePrivate(b *bool) *AccountUpdate {
	if b != nil {
		au.SetPrivate(*b)
	}
	return au
}

// SetClanID sets the "clan_id" field.
func (au *AccountUpdate) SetClanID(s string) *AccountUpdate {
	au.mutation.SetClanID(s)
	return au
}

// SetNillableClanID sets the "clan_id" field if the given value is not nil.
func (au *AccountUpdate) SetNillableClanID(s *string) *AccountUpdate {
	if s != nil {
		au.SetClanID(*s)
	}
	return au
}

// ClearClanID clears the value of the "clan_id" field.
func (au *AccountUpdate) ClearClanID() *AccountUpdate {
	au.mutation.ClearClanID()
	return au
}

// SetClan sets the "clan" edge to the Clan entity.
func (au *AccountUpdate) SetClan(c *Clan) *AccountUpdate {
	return au.SetClanID(c.ID)
}

// AddSnapshotIDs adds the "snapshots" edge to the AccountSnapshot entity by IDs.
func (au *AccountUpdate) AddSnapshotIDs(ids ...string) *AccountUpdate {
	au.mutation.AddSnapshotIDs(ids...)
	return au
}

// AddSnapshots adds the "snapshots" edges to the AccountSnapshot entity.
func (au *AccountUpdate) AddSnapshots(a ...*AccountSnapshot) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddSnapshotIDs(ids...)
}

// AddVehicleSnapshotIDs adds the "vehicle_snapshots" edge to the VehicleSnapshot entity by IDs.
func (au *AccountUpdate) AddVehicleSnapshotIDs(ids ...string) *AccountUpdate {
	au.mutation.AddVehicleSnapshotIDs(ids...)
	return au
}

// AddVehicleSnapshots adds the "vehicle_snapshots" edges to the VehicleSnapshot entity.
func (au *AccountUpdate) AddVehicleSnapshots(v ...*VehicleSnapshot) *AccountUpdate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.AddVehicleSnapshotIDs(ids...)
}

// AddAchievementSnapshotIDs adds the "achievement_snapshots" edge to the AchievementsSnapshot entity by IDs.
func (au *AccountUpdate) AddAchievementSnapshotIDs(ids ...string) *AccountUpdate {
	au.mutation.AddAchievementSnapshotIDs(ids...)
	return au
}

// AddAchievementSnapshots adds the "achievement_snapshots" edges to the AchievementsSnapshot entity.
func (au *AccountUpdate) AddAchievementSnapshots(a ...*AchievementsSnapshot) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAchievementSnapshotIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// ClearClan clears the "clan" edge to the Clan entity.
func (au *AccountUpdate) ClearClan() *AccountUpdate {
	au.mutation.ClearClan()
	return au
}

// ClearSnapshots clears all "snapshots" edges to the AccountSnapshot entity.
func (au *AccountUpdate) ClearSnapshots() *AccountUpdate {
	au.mutation.ClearSnapshots()
	return au
}

// RemoveSnapshotIDs removes the "snapshots" edge to AccountSnapshot entities by IDs.
func (au *AccountUpdate) RemoveSnapshotIDs(ids ...string) *AccountUpdate {
	au.mutation.RemoveSnapshotIDs(ids...)
	return au
}

// RemoveSnapshots removes "snapshots" edges to AccountSnapshot entities.
func (au *AccountUpdate) RemoveSnapshots(a ...*AccountSnapshot) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveSnapshotIDs(ids...)
}

// ClearVehicleSnapshots clears all "vehicle_snapshots" edges to the VehicleSnapshot entity.
func (au *AccountUpdate) ClearVehicleSnapshots() *AccountUpdate {
	au.mutation.ClearVehicleSnapshots()
	return au
}

// RemoveVehicleSnapshotIDs removes the "vehicle_snapshots" edge to VehicleSnapshot entities by IDs.
func (au *AccountUpdate) RemoveVehicleSnapshotIDs(ids ...string) *AccountUpdate {
	au.mutation.RemoveVehicleSnapshotIDs(ids...)
	return au
}

// RemoveVehicleSnapshots removes "vehicle_snapshots" edges to VehicleSnapshot entities.
func (au *AccountUpdate) RemoveVehicleSnapshots(v ...*VehicleSnapshot) *AccountUpdate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return au.RemoveVehicleSnapshotIDs(ids...)
}

// ClearAchievementSnapshots clears all "achievement_snapshots" edges to the AchievementsSnapshot entity.
func (au *AccountUpdate) ClearAchievementSnapshots() *AccountUpdate {
	au.mutation.ClearAchievementSnapshots()
	return au
}

// RemoveAchievementSnapshotIDs removes the "achievement_snapshots" edge to AchievementsSnapshot entities by IDs.
func (au *AccountUpdate) RemoveAchievementSnapshotIDs(ids ...string) *AccountUpdate {
	au.mutation.RemoveAchievementSnapshotIDs(ids...)
	return au
}

// RemoveAchievementSnapshots removes "achievement_snapshots" edges to AchievementsSnapshot entities.
func (au *AccountUpdate) RemoveAchievementSnapshots(a ...*AchievementsSnapshot) *AccountUpdate {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAchievementSnapshotIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Realm(); ok {
		if err := account.RealmValidator(v); err != nil {
			return &ValidationError{Name: "realm", err: fmt.Errorf(`db: validator failed for field "Account.realm": %w`, err)}
		}
	}
	if v, ok := au.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`db: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (au *AccountUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccountUpdate {
	au.modifiers = append(au.modifiers, modifiers...)
	return au
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeString))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.LastBattleTime(); ok {
		_spec.SetField(account.FieldLastBattleTime, field.TypeTime, value)
	}
	if value, ok := au.mutation.AccountCreatedAt(); ok {
		_spec.SetField(account.FieldAccountCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Realm(); ok {
		_spec.SetField(account.FieldRealm, field.TypeString, value)
	}
	if value, ok := au.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
	}
	if value, ok := au.mutation.Private(); ok {
		_spec.SetField(account.FieldPrivate, field.TypeBool, value)
	}
	if au.mutation.ClanCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ClanIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.SnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedSnapshotsIDs(); len(nodes) > 0 && !au.mutation.SnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.SnapshotsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.VehicleSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedVehicleSnapshotsIDs(); len(nodes) > 0 && !au.mutation.VehicleSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.VehicleSnapshotsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AchievementSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAchievementSnapshotsIDs(); len(nodes) > 0 && !au.mutation.AchievementSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AchievementSnapshotsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(au.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AccountMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetLastBattleTime sets the "last_battle_time" field.
func (auo *AccountUpdateOne) SetLastBattleTime(t time.Time) *AccountUpdateOne {
	auo.mutation.SetLastBattleTime(t)
	return auo
}

// SetNillableLastBattleTime sets the "last_battle_time" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableLastBattleTime(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetLastBattleTime(*t)
	}
	return auo
}

// SetAccountCreatedAt sets the "account_created_at" field.
func (auo *AccountUpdateOne) SetAccountCreatedAt(t time.Time) *AccountUpdateOne {
	auo.mutation.SetAccountCreatedAt(t)
	return auo
}

// SetNillableAccountCreatedAt sets the "account_created_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAccountCreatedAt(t *time.Time) *AccountUpdateOne {
	if t != nil {
		auo.SetAccountCreatedAt(*t)
	}
	return auo
}

// SetRealm sets the "realm" field.
func (auo *AccountUpdateOne) SetRealm(s string) *AccountUpdateOne {
	auo.mutation.SetRealm(s)
	return auo
}

// SetNillableRealm sets the "realm" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableRealm(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetRealm(*s)
	}
	return auo
}

// SetNickname sets the "nickname" field.
func (auo *AccountUpdateOne) SetNickname(s string) *AccountUpdateOne {
	auo.mutation.SetNickname(s)
	return auo
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableNickname(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetNickname(*s)
	}
	return auo
}

// SetPrivate sets the "private" field.
func (auo *AccountUpdateOne) SetPrivate(b bool) *AccountUpdateOne {
	auo.mutation.SetPrivate(b)
	return auo
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePrivate(b *bool) *AccountUpdateOne {
	if b != nil {
		auo.SetPrivate(*b)
	}
	return auo
}

// SetClanID sets the "clan_id" field.
func (auo *AccountUpdateOne) SetClanID(s string) *AccountUpdateOne {
	auo.mutation.SetClanID(s)
	return auo
}

// SetNillableClanID sets the "clan_id" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableClanID(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetClanID(*s)
	}
	return auo
}

// ClearClanID clears the value of the "clan_id" field.
func (auo *AccountUpdateOne) ClearClanID() *AccountUpdateOne {
	auo.mutation.ClearClanID()
	return auo
}

// SetClan sets the "clan" edge to the Clan entity.
func (auo *AccountUpdateOne) SetClan(c *Clan) *AccountUpdateOne {
	return auo.SetClanID(c.ID)
}

// AddSnapshotIDs adds the "snapshots" edge to the AccountSnapshot entity by IDs.
func (auo *AccountUpdateOne) AddSnapshotIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.AddSnapshotIDs(ids...)
	return auo
}

// AddSnapshots adds the "snapshots" edges to the AccountSnapshot entity.
func (auo *AccountUpdateOne) AddSnapshots(a ...*AccountSnapshot) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddSnapshotIDs(ids...)
}

// AddVehicleSnapshotIDs adds the "vehicle_snapshots" edge to the VehicleSnapshot entity by IDs.
func (auo *AccountUpdateOne) AddVehicleSnapshotIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.AddVehicleSnapshotIDs(ids...)
	return auo
}

// AddVehicleSnapshots adds the "vehicle_snapshots" edges to the VehicleSnapshot entity.
func (auo *AccountUpdateOne) AddVehicleSnapshots(v ...*VehicleSnapshot) *AccountUpdateOne {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.AddVehicleSnapshotIDs(ids...)
}

// AddAchievementSnapshotIDs adds the "achievement_snapshots" edge to the AchievementsSnapshot entity by IDs.
func (auo *AccountUpdateOne) AddAchievementSnapshotIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.AddAchievementSnapshotIDs(ids...)
	return auo
}

// AddAchievementSnapshots adds the "achievement_snapshots" edges to the AchievementsSnapshot entity.
func (auo *AccountUpdateOne) AddAchievementSnapshots(a ...*AchievementsSnapshot) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAchievementSnapshotIDs(ids...)
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// ClearClan clears the "clan" edge to the Clan entity.
func (auo *AccountUpdateOne) ClearClan() *AccountUpdateOne {
	auo.mutation.ClearClan()
	return auo
}

// ClearSnapshots clears all "snapshots" edges to the AccountSnapshot entity.
func (auo *AccountUpdateOne) ClearSnapshots() *AccountUpdateOne {
	auo.mutation.ClearSnapshots()
	return auo
}

// RemoveSnapshotIDs removes the "snapshots" edge to AccountSnapshot entities by IDs.
func (auo *AccountUpdateOne) RemoveSnapshotIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.RemoveSnapshotIDs(ids...)
	return auo
}

// RemoveSnapshots removes "snapshots" edges to AccountSnapshot entities.
func (auo *AccountUpdateOne) RemoveSnapshots(a ...*AccountSnapshot) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveSnapshotIDs(ids...)
}

// ClearVehicleSnapshots clears all "vehicle_snapshots" edges to the VehicleSnapshot entity.
func (auo *AccountUpdateOne) ClearVehicleSnapshots() *AccountUpdateOne {
	auo.mutation.ClearVehicleSnapshots()
	return auo
}

// RemoveVehicleSnapshotIDs removes the "vehicle_snapshots" edge to VehicleSnapshot entities by IDs.
func (auo *AccountUpdateOne) RemoveVehicleSnapshotIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.RemoveVehicleSnapshotIDs(ids...)
	return auo
}

// RemoveVehicleSnapshots removes "vehicle_snapshots" edges to VehicleSnapshot entities.
func (auo *AccountUpdateOne) RemoveVehicleSnapshots(v ...*VehicleSnapshot) *AccountUpdateOne {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return auo.RemoveVehicleSnapshotIDs(ids...)
}

// ClearAchievementSnapshots clears all "achievement_snapshots" edges to the AchievementsSnapshot entity.
func (auo *AccountUpdateOne) ClearAchievementSnapshots() *AccountUpdateOne {
	auo.mutation.ClearAchievementSnapshots()
	return auo
}

// RemoveAchievementSnapshotIDs removes the "achievement_snapshots" edge to AchievementsSnapshot entities by IDs.
func (auo *AccountUpdateOne) RemoveAchievementSnapshotIDs(ids ...string) *AccountUpdateOne {
	auo.mutation.RemoveAchievementSnapshotIDs(ids...)
	return auo
}

// RemoveAchievementSnapshots removes "achievement_snapshots" edges to AchievementsSnapshot entities.
func (auo *AccountUpdateOne) RemoveAchievementSnapshots(a ...*AchievementsSnapshot) *AccountUpdateOne {
	ids := make([]string, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAchievementSnapshotIDs(ids...)
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Realm(); ok {
		if err := account.RealmValidator(v); err != nil {
			return &ValidationError{Name: "realm", err: fmt.Errorf(`db: validator failed for field "Account.realm": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Nickname(); ok {
		if err := account.NicknameValidator(v); err != nil {
			return &ValidationError{Name: "nickname", err: fmt.Errorf(`db: validator failed for field "Account.nickname": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (auo *AccountUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AccountUpdateOne {
	auo.modifiers = append(auo.modifiers, modifiers...)
	return auo
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeString))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.LastBattleTime(); ok {
		_spec.SetField(account.FieldLastBattleTime, field.TypeTime, value)
	}
	if value, ok := auo.mutation.AccountCreatedAt(); ok {
		_spec.SetField(account.FieldAccountCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Realm(); ok {
		_spec.SetField(account.FieldRealm, field.TypeString, value)
	}
	if value, ok := auo.mutation.Nickname(); ok {
		_spec.SetField(account.FieldNickname, field.TypeString, value)
	}
	if value, ok := auo.mutation.Private(); ok {
		_spec.SetField(account.FieldPrivate, field.TypeBool, value)
	}
	if auo.mutation.ClanCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ClanIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.SnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedSnapshotsIDs(); len(nodes) > 0 && !auo.mutation.SnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.SnapshotsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.VehicleSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedVehicleSnapshotsIDs(); len(nodes) > 0 && !auo.mutation.VehicleSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.VehicleSnapshotsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AchievementSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAchievementSnapshotsIDs(); len(nodes) > 0 && !auo.mutation.AchievementSnapshotsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AchievementSnapshotsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(auo.modifiers...)
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
