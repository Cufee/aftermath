// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/accountsnapshot"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/stats/frame"
)

// AccountSnapshotUpdate is the builder for updating AccountSnapshot entities.
type AccountSnapshotUpdate struct {
	config
	hooks    []Hook
	mutation *AccountSnapshotMutation
}

// Where appends a list predicates to the AccountSnapshotUpdate builder.
func (asu *AccountSnapshotUpdate) Where(ps ...predicate.AccountSnapshot) *AccountSnapshotUpdate {
	asu.mutation.Where(ps...)
	return asu
}

// SetUpdatedAt sets the "updated_at" field.
func (asu *AccountSnapshotUpdate) SetUpdatedAt(i int64) *AccountSnapshotUpdate {
	asu.mutation.ResetUpdatedAt()
	asu.mutation.SetUpdatedAt(i)
	return asu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (asu *AccountSnapshotUpdate) AddUpdatedAt(i int64) *AccountSnapshotUpdate {
	asu.mutation.AddUpdatedAt(i)
	return asu
}

// SetType sets the "type" field.
func (asu *AccountSnapshotUpdate) SetType(mt models.SnapshotType) *AccountSnapshotUpdate {
	asu.mutation.SetType(mt)
	return asu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableType(mt *models.SnapshotType) *AccountSnapshotUpdate {
	if mt != nil {
		asu.SetType(*mt)
	}
	return asu
}

// SetLastBattleTime sets the "last_battle_time" field.
func (asu *AccountSnapshotUpdate) SetLastBattleTime(i int64) *AccountSnapshotUpdate {
	asu.mutation.ResetLastBattleTime()
	asu.mutation.SetLastBattleTime(i)
	return asu
}

// SetNillableLastBattleTime sets the "last_battle_time" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableLastBattleTime(i *int64) *AccountSnapshotUpdate {
	if i != nil {
		asu.SetLastBattleTime(*i)
	}
	return asu
}

// AddLastBattleTime adds i to the "last_battle_time" field.
func (asu *AccountSnapshotUpdate) AddLastBattleTime(i int64) *AccountSnapshotUpdate {
	asu.mutation.AddLastBattleTime(i)
	return asu
}

// SetReferenceID sets the "reference_id" field.
func (asu *AccountSnapshotUpdate) SetReferenceID(s string) *AccountSnapshotUpdate {
	asu.mutation.SetReferenceID(s)
	return asu
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableReferenceID(s *string) *AccountSnapshotUpdate {
	if s != nil {
		asu.SetReferenceID(*s)
	}
	return asu
}

// SetRatingBattles sets the "rating_battles" field.
func (asu *AccountSnapshotUpdate) SetRatingBattles(i int) *AccountSnapshotUpdate {
	asu.mutation.ResetRatingBattles()
	asu.mutation.SetRatingBattles(i)
	return asu
}

// SetNillableRatingBattles sets the "rating_battles" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableRatingBattles(i *int) *AccountSnapshotUpdate {
	if i != nil {
		asu.SetRatingBattles(*i)
	}
	return asu
}

// AddRatingBattles adds i to the "rating_battles" field.
func (asu *AccountSnapshotUpdate) AddRatingBattles(i int) *AccountSnapshotUpdate {
	asu.mutation.AddRatingBattles(i)
	return asu
}

// SetRatingFrame sets the "rating_frame" field.
func (asu *AccountSnapshotUpdate) SetRatingFrame(ff frame.StatsFrame) *AccountSnapshotUpdate {
	asu.mutation.SetRatingFrame(ff)
	return asu
}

// SetNillableRatingFrame sets the "rating_frame" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableRatingFrame(ff *frame.StatsFrame) *AccountSnapshotUpdate {
	if ff != nil {
		asu.SetRatingFrame(*ff)
	}
	return asu
}

// SetRegularBattles sets the "regular_battles" field.
func (asu *AccountSnapshotUpdate) SetRegularBattles(i int) *AccountSnapshotUpdate {
	asu.mutation.ResetRegularBattles()
	asu.mutation.SetRegularBattles(i)
	return asu
}

// SetNillableRegularBattles sets the "regular_battles" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableRegularBattles(i *int) *AccountSnapshotUpdate {
	if i != nil {
		asu.SetRegularBattles(*i)
	}
	return asu
}

// AddRegularBattles adds i to the "regular_battles" field.
func (asu *AccountSnapshotUpdate) AddRegularBattles(i int) *AccountSnapshotUpdate {
	asu.mutation.AddRegularBattles(i)
	return asu
}

// SetRegularFrame sets the "regular_frame" field.
func (asu *AccountSnapshotUpdate) SetRegularFrame(ff frame.StatsFrame) *AccountSnapshotUpdate {
	asu.mutation.SetRegularFrame(ff)
	return asu
}

// SetNillableRegularFrame sets the "regular_frame" field if the given value is not nil.
func (asu *AccountSnapshotUpdate) SetNillableRegularFrame(ff *frame.StatsFrame) *AccountSnapshotUpdate {
	if ff != nil {
		asu.SetRegularFrame(*ff)
	}
	return asu
}

// Mutation returns the AccountSnapshotMutation object of the builder.
func (asu *AccountSnapshotUpdate) Mutation() *AccountSnapshotMutation {
	return asu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (asu *AccountSnapshotUpdate) Save(ctx context.Context) (int, error) {
	asu.defaults()
	return withHooks(ctx, asu.sqlSave, asu.mutation, asu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asu *AccountSnapshotUpdate) SaveX(ctx context.Context) int {
	affected, err := asu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (asu *AccountSnapshotUpdate) Exec(ctx context.Context) error {
	_, err := asu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asu *AccountSnapshotUpdate) ExecX(ctx context.Context) {
	if err := asu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asu *AccountSnapshotUpdate) defaults() {
	if _, ok := asu.mutation.UpdatedAt(); !ok {
		v := accountsnapshot.UpdateDefaultUpdatedAt()
		asu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asu *AccountSnapshotUpdate) check() error {
	if v, ok := asu.mutation.GetType(); ok {
		if err := accountsnapshot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.type": %w`, err)}
		}
	}
	if v, ok := asu.mutation.ReferenceID(); ok {
		if err := accountsnapshot.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.reference_id": %w`, err)}
		}
	}
	if _, ok := asu.mutation.AccountID(); asu.mutation.AccountCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "AccountSnapshot.account"`)
	}
	return nil
}

func (asu *AccountSnapshotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := asu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(accountsnapshot.Table, accountsnapshot.Columns, sqlgraph.NewFieldSpec(accountsnapshot.FieldID, field.TypeString))
	if ps := asu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asu.mutation.UpdatedAt(); ok {
		_spec.SetField(accountsnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(accountsnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.GetType(); ok {
		_spec.SetField(accountsnapshot.FieldType, field.TypeEnum, value)
	}
	if value, ok := asu.mutation.LastBattleTime(); ok {
		_spec.SetField(accountsnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.AddedLastBattleTime(); ok {
		_spec.AddField(accountsnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asu.mutation.ReferenceID(); ok {
		_spec.SetField(accountsnapshot.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := asu.mutation.RatingBattles(); ok {
		_spec.SetField(accountsnapshot.FieldRatingBattles, field.TypeInt, value)
	}
	if value, ok := asu.mutation.AddedRatingBattles(); ok {
		_spec.AddField(accountsnapshot.FieldRatingBattles, field.TypeInt, value)
	}
	if value, ok := asu.mutation.RatingFrame(); ok {
		_spec.SetField(accountsnapshot.FieldRatingFrame, field.TypeJSON, value)
	}
	if value, ok := asu.mutation.RegularBattles(); ok {
		_spec.SetField(accountsnapshot.FieldRegularBattles, field.TypeInt, value)
	}
	if value, ok := asu.mutation.AddedRegularBattles(); ok {
		_spec.AddField(accountsnapshot.FieldRegularBattles, field.TypeInt, value)
	}
	if value, ok := asu.mutation.RegularFrame(); ok {
		_spec.SetField(accountsnapshot.FieldRegularFrame, field.TypeJSON, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, asu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountsnapshot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	asu.mutation.done = true
	return n, nil
}

// AccountSnapshotUpdateOne is the builder for updating a single AccountSnapshot entity.
type AccountSnapshotUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountSnapshotMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (asuo *AccountSnapshotUpdateOne) SetUpdatedAt(i int64) *AccountSnapshotUpdateOne {
	asuo.mutation.ResetUpdatedAt()
	asuo.mutation.SetUpdatedAt(i)
	return asuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (asuo *AccountSnapshotUpdateOne) AddUpdatedAt(i int64) *AccountSnapshotUpdateOne {
	asuo.mutation.AddUpdatedAt(i)
	return asuo
}

// SetType sets the "type" field.
func (asuo *AccountSnapshotUpdateOne) SetType(mt models.SnapshotType) *AccountSnapshotUpdateOne {
	asuo.mutation.SetType(mt)
	return asuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableType(mt *models.SnapshotType) *AccountSnapshotUpdateOne {
	if mt != nil {
		asuo.SetType(*mt)
	}
	return asuo
}

// SetLastBattleTime sets the "last_battle_time" field.
func (asuo *AccountSnapshotUpdateOne) SetLastBattleTime(i int64) *AccountSnapshotUpdateOne {
	asuo.mutation.ResetLastBattleTime()
	asuo.mutation.SetLastBattleTime(i)
	return asuo
}

// SetNillableLastBattleTime sets the "last_battle_time" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableLastBattleTime(i *int64) *AccountSnapshotUpdateOne {
	if i != nil {
		asuo.SetLastBattleTime(*i)
	}
	return asuo
}

// AddLastBattleTime adds i to the "last_battle_time" field.
func (asuo *AccountSnapshotUpdateOne) AddLastBattleTime(i int64) *AccountSnapshotUpdateOne {
	asuo.mutation.AddLastBattleTime(i)
	return asuo
}

// SetReferenceID sets the "reference_id" field.
func (asuo *AccountSnapshotUpdateOne) SetReferenceID(s string) *AccountSnapshotUpdateOne {
	asuo.mutation.SetReferenceID(s)
	return asuo
}

// SetNillableReferenceID sets the "reference_id" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableReferenceID(s *string) *AccountSnapshotUpdateOne {
	if s != nil {
		asuo.SetReferenceID(*s)
	}
	return asuo
}

// SetRatingBattles sets the "rating_battles" field.
func (asuo *AccountSnapshotUpdateOne) SetRatingBattles(i int) *AccountSnapshotUpdateOne {
	asuo.mutation.ResetRatingBattles()
	asuo.mutation.SetRatingBattles(i)
	return asuo
}

// SetNillableRatingBattles sets the "rating_battles" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableRatingBattles(i *int) *AccountSnapshotUpdateOne {
	if i != nil {
		asuo.SetRatingBattles(*i)
	}
	return asuo
}

// AddRatingBattles adds i to the "rating_battles" field.
func (asuo *AccountSnapshotUpdateOne) AddRatingBattles(i int) *AccountSnapshotUpdateOne {
	asuo.mutation.AddRatingBattles(i)
	return asuo
}

// SetRatingFrame sets the "rating_frame" field.
func (asuo *AccountSnapshotUpdateOne) SetRatingFrame(ff frame.StatsFrame) *AccountSnapshotUpdateOne {
	asuo.mutation.SetRatingFrame(ff)
	return asuo
}

// SetNillableRatingFrame sets the "rating_frame" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableRatingFrame(ff *frame.StatsFrame) *AccountSnapshotUpdateOne {
	if ff != nil {
		asuo.SetRatingFrame(*ff)
	}
	return asuo
}

// SetRegularBattles sets the "regular_battles" field.
func (asuo *AccountSnapshotUpdateOne) SetRegularBattles(i int) *AccountSnapshotUpdateOne {
	asuo.mutation.ResetRegularBattles()
	asuo.mutation.SetRegularBattles(i)
	return asuo
}

// SetNillableRegularBattles sets the "regular_battles" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableRegularBattles(i *int) *AccountSnapshotUpdateOne {
	if i != nil {
		asuo.SetRegularBattles(*i)
	}
	return asuo
}

// AddRegularBattles adds i to the "regular_battles" field.
func (asuo *AccountSnapshotUpdateOne) AddRegularBattles(i int) *AccountSnapshotUpdateOne {
	asuo.mutation.AddRegularBattles(i)
	return asuo
}

// SetRegularFrame sets the "regular_frame" field.
func (asuo *AccountSnapshotUpdateOne) SetRegularFrame(ff frame.StatsFrame) *AccountSnapshotUpdateOne {
	asuo.mutation.SetRegularFrame(ff)
	return asuo
}

// SetNillableRegularFrame sets the "regular_frame" field if the given value is not nil.
func (asuo *AccountSnapshotUpdateOne) SetNillableRegularFrame(ff *frame.StatsFrame) *AccountSnapshotUpdateOne {
	if ff != nil {
		asuo.SetRegularFrame(*ff)
	}
	return asuo
}

// Mutation returns the AccountSnapshotMutation object of the builder.
func (asuo *AccountSnapshotUpdateOne) Mutation() *AccountSnapshotMutation {
	return asuo.mutation
}

// Where appends a list predicates to the AccountSnapshotUpdate builder.
func (asuo *AccountSnapshotUpdateOne) Where(ps ...predicate.AccountSnapshot) *AccountSnapshotUpdateOne {
	asuo.mutation.Where(ps...)
	return asuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asuo *AccountSnapshotUpdateOne) Select(field string, fields ...string) *AccountSnapshotUpdateOne {
	asuo.fields = append([]string{field}, fields...)
	return asuo
}

// Save executes the query and returns the updated AccountSnapshot entity.
func (asuo *AccountSnapshotUpdateOne) Save(ctx context.Context) (*AccountSnapshot, error) {
	asuo.defaults()
	return withHooks(ctx, asuo.sqlSave, asuo.mutation, asuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (asuo *AccountSnapshotUpdateOne) SaveX(ctx context.Context) *AccountSnapshot {
	node, err := asuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asuo *AccountSnapshotUpdateOne) Exec(ctx context.Context) error {
	_, err := asuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asuo *AccountSnapshotUpdateOne) ExecX(ctx context.Context) {
	if err := asuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asuo *AccountSnapshotUpdateOne) defaults() {
	if _, ok := asuo.mutation.UpdatedAt(); !ok {
		v := accountsnapshot.UpdateDefaultUpdatedAt()
		asuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asuo *AccountSnapshotUpdateOne) check() error {
	if v, ok := asuo.mutation.GetType(); ok {
		if err := accountsnapshot.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.type": %w`, err)}
		}
	}
	if v, ok := asuo.mutation.ReferenceID(); ok {
		if err := accountsnapshot.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "AccountSnapshot.reference_id": %w`, err)}
		}
	}
	if _, ok := asuo.mutation.AccountID(); asuo.mutation.AccountCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "AccountSnapshot.account"`)
	}
	return nil
}

func (asuo *AccountSnapshotUpdateOne) sqlSave(ctx context.Context) (_node *AccountSnapshot, err error) {
	if err := asuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(accountsnapshot.Table, accountsnapshot.Columns, sqlgraph.NewFieldSpec(accountsnapshot.FieldID, field.TypeString))
	id, ok := asuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AccountSnapshot.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountsnapshot.FieldID)
		for _, f := range fields {
			if !accountsnapshot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != accountsnapshot.FieldID {
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
		_spec.SetField(accountsnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(accountsnapshot.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.GetType(); ok {
		_spec.SetField(accountsnapshot.FieldType, field.TypeEnum, value)
	}
	if value, ok := asuo.mutation.LastBattleTime(); ok {
		_spec.SetField(accountsnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.AddedLastBattleTime(); ok {
		_spec.AddField(accountsnapshot.FieldLastBattleTime, field.TypeInt64, value)
	}
	if value, ok := asuo.mutation.ReferenceID(); ok {
		_spec.SetField(accountsnapshot.FieldReferenceID, field.TypeString, value)
	}
	if value, ok := asuo.mutation.RatingBattles(); ok {
		_spec.SetField(accountsnapshot.FieldRatingBattles, field.TypeInt, value)
	}
	if value, ok := asuo.mutation.AddedRatingBattles(); ok {
		_spec.AddField(accountsnapshot.FieldRatingBattles, field.TypeInt, value)
	}
	if value, ok := asuo.mutation.RatingFrame(); ok {
		_spec.SetField(accountsnapshot.FieldRatingFrame, field.TypeJSON, value)
	}
	if value, ok := asuo.mutation.RegularBattles(); ok {
		_spec.SetField(accountsnapshot.FieldRegularBattles, field.TypeInt, value)
	}
	if value, ok := asuo.mutation.AddedRegularBattles(); ok {
		_spec.AddField(accountsnapshot.FieldRegularBattles, field.TypeInt, value)
	}
	if value, ok := asuo.mutation.RegularFrame(); ok {
		_spec.SetField(accountsnapshot.FieldRegularFrame, field.TypeJSON, value)
	}
	_node = &AccountSnapshot{config: asuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{accountsnapshot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	asuo.mutation.done = true
	return _node, nil
}