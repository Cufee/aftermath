// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/achievementssnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/am-wg-proxy-next/v2/types"
)

// AchievementsSnapshotUpdate is the builder for updating AchievementsSnapshot entities.
type AchievementsSnapshotUpdate struct {
	config
	hooks    []Hook
	mutation *AchievementsSnapshotMutation
}

// Where appends a list predicates to the AchievementsSnapshotUpdate builder.
func (asu *AchievementsSnapshotUpdate) Where(ps ...predicate.AchievementsSnapshot) *AchievementsSnapshotUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetUpdatedAt sets the "updated_at" field.
func (asu *AchievementsSnapshotUpdate) SetUpdatedAt(i int64) *AchievementsSnapshotUpdate {
	asu.mutation.ResetUpdatedAt()
	asu.mutation.SetUpdatedAt(i)
	return asu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (asu *AchievementsSnapshotUpdate) AddUpdatedAt(i int64) *AchievementsSnapshotUpdate {
	asu.mutation.AddUpdatedAt(i)
	return asu
}

// SetType sets the "type" field.
func (asu *AchievementsSnapshotUpdate) SetType(mt models.SnapshotType) *AchievementsSnapshotUpdate {
	asu.mutation.SetType(mt)
	return asu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (asu *AchievementsSnapshotUpdate) SetNillableType(mt *models.SnapshotType) *AchievementsSnapshotUpdate {
	if mt != nil {
		asu.SetType(*mt)
	}
	return asu
}

// SetReferenceID sets the "reference_id" field.
func (asu *AchievementsSnapshotUpdate) SetReferenceID(s string) *AchievementsSnapshotUpdate {
	asu.mutation.SetReferenceID(s)
	return asu
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (asu *AchievementsSnapshotUpdate) SetNillableReferenceID(s *string) *AchievementsSnapshotUpdate {
	if s != nil {
		asu.SetReferenceID(*s)
	}
	return asu
}

// SetBattles sets the "battles" field.
func (asu *AchievementsSnapshotUpdate) SetBattles(i int) *AchievementsSnapshotUpdate {
	asu.mutation.ResetBattles()
	asu.mutation.SetBattles(i)
	return asu
}

// SetNillableBattles sets the "battles" field if the given value is not nil.
func (asu *AchievementsSnapshotUpdate) SetNillableBattles(i *int) *AchievementsSnapshotUpdate {
	if i != nil {
		asu.SetBattles(*i)
	}
	return asu
}

// AddBattles adds i to the "battles" field.
func (asu *AchievementsSnapshotUpdate) AddBattles(i int) *AchievementsSnapshotUpdate {
	asu.mutation.AddBattles(i)
	return asu
}

// SetLastBattleTime sets the "last_battle_time" field.
func (asu *AchievementsSnapshotUpdate) SetLastBattleTime(i int64) *AchievementsSnapshotUpdate {
	asu.mutation.ResetLastBattleTime()
	asu.mutation.SetLastBattleTime(i)
	return asu
}

// SetNillableLastBattleTime sets the "last_battle_time" field if the given value is not nil.
func (asu *AchievementsSnapshotUpdate) SetNillableLastBattleTime(i *int64) *AchievementsSnapshotUpdate {
	if i != nil {
		asu.SetLastBattleTime(*i)
	}
	return asu
}

// AddLastBattleTime adds i to the "last_battle_time" field.
func (asu *AchievementsSnapshotUpdate) AddLastBattleTime(i int64) *AchievementsSnapshotUpdate {
	asu.mutation.AddLastBattleTime(i)
	return asu
}

// SetData sets the "data" field.
func (asu *AchievementsSnapshotUpdate) SetData(tf types.AchievementsFrame) *AchievementsSnapshotUpdate {
	asu.mutation.SetData(tf)
	return asu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (asu *AchievementsSnapshotUpdate) SetNillableData(tf *types.AchievementsFrame) *AchievementsSnapshotUpdate {
	if tf != nil {
		asu.SetData(*tf)
	}
	return asu
}

// Mutation returns the AchievementsSnapshotMutation object of the builder.
func (asu *AchievementsSnapshotUpdate) Mutation() *AchievementsSnapshotMutation {
	return asu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *AchievementsSnapshotUpdate) Save(ctx context.Context) (int, error) {
	asu.defaults()
	return withHooks(ctx, asu.sqlSave, asu.mutation, asu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asu *AchievementsSnapshotUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *AchievementsSnapshotUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *AchievementsSnapshotUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asu *AchievementsSnapshotUpdate) defaults() {
	if _, ok := asu.mutation.UpdatedAt(); !ok {
		v := achievementssnapshot.UpdateDefaultUpdatedAt()
		asu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asu *AchievementsSnapshotUpdate) check() error {
	if v, ok := asu.mutation.GetType(); ok {
		if err := achievementssnapshot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "AchievementsSnapshot.type": %w`, err)}
		}
	}
	if v, ok := asu.mutation.ReferenceID(); ok {
		if err := achievementssnapshot.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "AchievementsSnapshot.reference_id": %w`, err)}
		}
	}
	if _, ok := asu.mutation.AccountID(); asu.mutation.AccountCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "AchievementsSnapshot.account"`)
	}
	return nil
}

func (asu *AchievementsSnapshotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := asu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(achievementssnapshot.Table, achievementssnapshot.Columns, sqlgraph.NewFieldSpec(achievementssnapshot.FieldID, field.TypeString))
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.UpdatedAt(); ok {
		_spec.SetField(achievementssnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(achievementssnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.GetType(); ok {
		_spec.SetField(achievementssnapshot.FieldType, field.TypeEnum, value)
	}
	if value, ok := asu.mutation.ReferenceID(); ok {
		_spec.SetField(achievementssnapshot.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := asu.mutation.Battles(); ok {
		_spec.SetField(achievementssnapshot.FieldBattles, field.TypeInt, value)
	}
	if value, ok := asu.mutation.AddedBattles(); ok {
		_spec.AddField(achievementssnapshot.FieldBattles, field.TypeInt, value)
	}
	if value, ok := asu.mutation.LastBattleTime(); ok {
		_spec.SetField(achievementssnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedLastBattleTime(); ok {
		_spec.AddField(achievementssnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.Data(); ok {
		_spec.SetField(achievementssnapshot.FieldData, field.TypeJSON, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievementssnapshot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	asu.mutation.done = true
	return n, nil
}

// AchievementsSnapshotUpdateOne is the builder for updating a single AchievementsSnapshot entity.
type AchievementsSnapshotUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AchievementsSnapshotMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (asuo *AchievementsSnapshotUpdateOne) SetUpdatedAt(i int64) *AchievementsSnapshotUpdateOne {
	asuo.mutation.ResetUpdatedAt()
	asuo.mutation.SetUpdatedAt(i)
	return asuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (asuo *AchievementsSnapshotUpdateOne) AddUpdatedAt(i int64) *AchievementsSnapshotUpdateOne {
	asuo.mutation.AddUpdatedAt(i)
	return asuo
}

// SetType sets the "type" field.
func (asuo *AchievementsSnapshotUpdateOne) SetType(mt models.SnapshotType) *AchievementsSnapshotUpdateOne {
	asuo.mutation.SetType(mt)
	return asuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (asuo *AchievementsSnapshotUpdateOne) SetNillableType(mt *models.SnapshotType) *AchievementsSnapshotUpdateOne {
	if mt != nil {
		asuo.SetType(*mt)
	}
	return asuo
}

// SetReferenceID sets the "reference_id" field.
func (asuo *AchievementsSnapshotUpdateOne) SetReferenceID(s string) *AchievementsSnapshotUpdateOne {
	asuo.mutation.SetReferenceID(s)
	return asuo
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (asuo *AchievementsSnapshotUpdateOne) SetNillableReferenceID(s *string) *AchievementsSnapshotUpdateOne {
	if s != nil {
		asuo.SetReferenceID(*s)
	}
	return asuo
}

// SetBattles sets the "battles" field.
func (asuo *AchievementsSnapshotUpdateOne) SetBattles(i int) *AchievementsSnapshotUpdateOne {
	asuo.mutation.ResetBattles()
	asuo.mutation.SetBattles(i)
	return asuo
}

// SetNillableBattles sets the "battles" field if the given value is not nil.
func (asuo *AchievementsSnapshotUpdateOne) SetNillableBattles(i *int) *AchievementsSnapshotUpdateOne {
	if i != nil {
		asuo.SetBattles(*i)
	}
	return asuo
}

// AddBattles adds i to the "battles" field.
func (asuo *AchievementsSnapshotUpdateOne) AddBattles(i int) *AchievementsSnapshotUpdateOne {
	asuo.mutation.AddBattles(i)
	return asuo
}

// SetLastBattleTime sets the "last_battle_time" field.
func (asuo *AchievementsSnapshotUpdateOne) SetLastBattleTime(i int64) *AchievementsSnapshotUpdateOne {
	asuo.mutation.ResetLastBattleTime()
	asuo.mutation.SetLastBattleTime(i)
	return asuo
}

// SetNillableLastBattleTime sets the "last_battle_time" field if the given value is not nil.
func (asuo *AchievementsSnapshotUpdateOne) SetNillableLastBattleTime(i *int64) *AchievementsSnapshotUpdateOne {
	if i != nil {
		asuo.SetLastBattleTime(*i)
	}
	return asuo
}

// AddLastBattleTime adds i to the "last_battle_time" field.
func (asuo *AchievementsSnapshotUpdateOne) AddLastBattleTime(i int64) *AchievementsSnapshotUpdateOne {
	asuo.mutation.AddLastBattleTime(i)
	return asuo
}

// SetData sets the "data" field.
func (asuo *AchievementsSnapshotUpdateOne) SetData(tf types.AchievementsFrame) *AchievementsSnapshotUpdateOne {
	asuo.mutation.SetData(tf)
	return asuo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (asuo *AchievementsSnapshotUpdateOne) SetNillableData(tf *types.AchievementsFrame) *AchievementsSnapshotUpdateOne {
	if tf != nil {
		asuo.SetData(*tf)
	}
	return asuo
}

// Mutation returns the AchievementsSnapshotMutation object of the builder.
func (asuo *AchievementsSnapshotUpdateOne) Mutation() *AchievementsSnapshotMutation {
	return asuo.mutation
}

// Where appends a list predicates to the AchievementsSnapshotUpdate builder.
func (asuo *AchievementsSnapshotUpdateOne) Where(ps ...predicate.AchievementsSnapshot) *AchievementsSnapshotUpdateOne {
	asuo.mutation.Where(ps...)
	return asuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *AchievementsSnapshotUpdateOne) Select(field string, fields ...string) *AchievementsSnapshotUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated AchievementsSnapshot entity.
func (asuo *AchievementsSnapshotUpdateOne) Save(ctx context.Context) (*AchievementsSnapshot, error) {
	asuo.defaults()
	return withHooks(ctx, asuo.sqlSave, asuo.mutation, asuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *AchievementsSnapshotUpdateOne) SaveX(ctx context.Context) *AchievementsSnapshot {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *AchievementsSnapshotUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *AchievementsSnapshotUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asuo *AchievementsSnapshotUpdateOne) defaults() {
	if _, ok := asuo.mutation.UpdatedAt(); !ok {
		v := achievementssnapshot.UpdateDefaultUpdatedAt()
		asuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asuo *AchievementsSnapshotUpdateOne) check() error {
	if v, ok := asuo.mutation.GetType(); ok {
		if err := achievementssnapshot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "AchievementsSnapshot.type": %w`, err)}
		}
	}
	if v, ok := asuo.mutation.ReferenceID(); ok {
		if err := achievementssnapshot.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "AchievementsSnapshot.reference_id": %w`, err)}
		}
	}
	if _, ok := asuo.mutation.AccountID(); asuo.mutation.AccountCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "AchievementsSnapshot.account"`)
	}
	return nil
}

func (asuo *AchievementsSnapshotUpdateOne) sqlSave(ctx context.Context) (_node *AchievementsSnapshot, err error) {
	if err := asuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(achievementssnapshot.Table, achievementssnapshot.Columns, sqlgraph.NewFieldSpec(achievementssnapshot.FieldID, field.TypeString))
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AchievementsSnapshot.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, achievementssnapshot.FieldID)
		for _, f := range fields {
			if !achievementssnapshot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != achievementssnapshot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asuo.mutation.UpdatedAt(); ok {
		_spec.SetField(achievementssnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(achievementssnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.GetType(); ok {
		_spec.SetField(achievementssnapshot.FieldType, field.TypeEnum, value)
	}
	if value, ok := asuo.mutation.ReferenceID(); ok {
		_spec.SetField(achievementssnapshot.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := asuo.mutation.Battles(); ok {
		_spec.SetField(achievementssnapshot.FieldBattles, field.TypeInt, value)
	}
	if value, ok := asuo.mutation.AddedBattles(); ok {
		_spec.AddField(achievementssnapshot.FieldBattles, field.TypeInt, value)
	}
	if value, ok := asuo.mutation.LastBattleTime(); ok {
		_spec.SetField(achievementssnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedLastBattleTime(); ok {
		_spec.AddField(achievementssnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.Data(); ok {
		_spec.SetField(achievementssnapshot.FieldData, field.TypeJSON, value)
	}
	_node = &AchievementsSnapshot{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{achievementssnapshot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	asuo.mutation.done = true
	return _node, nil
}