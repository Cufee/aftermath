// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/gamemap"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"golang.org/x/text/language"
)

// GameMapUpdate is the builder for updating GameMap entities.
type GameMapUpdate struct {
	config
	hooks     []Hook
	mutation  *GameMapMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GameMapUpdate builder.
func (gmu *GameMapUpdate) Where(ps ...predicate.GameMap) *GameMapUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetUpdatedAt sets the "updated_at" field.
func (gmu *GameMapUpdate) SetUpdatedAt(t time.Time) *GameMapUpdate {
	gmu.mutation.SetUpdatedAt(t)
	return gmu
}

// SetGameModes sets the "game_modes" field.
func (gmu *GameMapUpdate) SetGameModes(i []int) *GameMapUpdate {
	gmu.mutation.SetGameModes(i)
	return gmu
}

// AppendGameModes appends i to the "game_modes" field.
func (gmu *GameMapUpdate) AppendGameModes(i []int) *GameMapUpdate {
	gmu.mutation.AppendGameModes(i)
	return gmu
}

// SetSupremacyPoints sets the "supremacy_points" field.
func (gmu *GameMapUpdate) SetSupremacyPoints(i int) *GameMapUpdate {
	gmu.mutation.ResetSupremacyPoints()
	gmu.mutation.SetSupremacyPoints(i)
	return gmu
}

// SetNillableSupremacyPoints sets the "supremacy_points" field if the given value is not nil.
func (gmu *GameMapUpdate) SetNillableSupremacyPoints(i *int) *GameMapUpdate {
	if i != nil {
		gmu.SetSupremacyPoints(*i)
	}
	return gmu
}

// AddSupremacyPoints adds i to the "supremacy_points" field.
func (gmu *GameMapUpdate) AddSupremacyPoints(i int) *GameMapUpdate {
	gmu.mutation.AddSupremacyPoints(i)
	return gmu
}

// SetLocalizedNames sets the "localized_names" field.
func (gmu *GameMapUpdate) SetLocalizedNames(m map[language.Tag]string) *GameMapUpdate {
	gmu.mutation.SetLocalizedNames(m)
	return gmu
}

// Mutation returns the GameMapMutation object of the builder.
func (gmu *GameMapUpdate) Mutation() *GameMapMutation {
	return gmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GameMapUpdate) Save(ctx context.Context) (int, error) {
	gmu.defaults()
	return withHooks(ctx, gmu.sqlSave, gmu.mutation, gmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GameMapUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GameMapUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GameMapUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmu *GameMapUpdate) defaults() {
	if _, ok := gmu.mutation.UpdatedAt(); !ok {
		v := gamemap.UpdateDefaultUpdatedAt()
		gmu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gmu *GameMapUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GameMapUpdate {
	gmu.modifiers = append(gmu.modifiers, modifiers...)
	return gmu
}

func (gmu *GameMapUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(gamemap.Table, gamemap.Columns, sqlgraph.NewFieldSpec(gamemap.FieldID, field.TypeString))
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.UpdatedAt(); ok {
		_spec.SetField(gamemap.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmu.mutation.GameModes(); ok {
		_spec.SetField(gamemap.FieldGameModes, field.TypeJSON, value)
	}
	if value, ok := gmu.mutation.AppendedGameModes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, gamemap.FieldGameModes, value)
		})
	}
	if value, ok := gmu.mutation.SupremacyPoints(); ok {
		_spec.SetField(gamemap.FieldSupremacyPoints, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.AddedSupremacyPoints(); ok {
		_spec.AddField(gamemap.FieldSupremacyPoints, field.TypeInt, value)
	}
	if value, ok := gmu.mutation.LocalizedNames(); ok {
		_spec.SetField(gamemap.FieldLocalizedNames, field.TypeJSON, value)
	}
	_spec.AddModifiers(gmu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gamemap.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmu.mutation.done = true
	return n, nil
}

// GameMapUpdateOne is the builder for updating a single GameMap entity.
type GameMapUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GameMapMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (gmuo *GameMapUpdateOne) SetUpdatedAt(t time.Time) *GameMapUpdateOne {
	gmuo.mutation.SetUpdatedAt(t)
	return gmuo
}

// SetGameModes sets the "game_modes" field.
func (gmuo *GameMapUpdateOne) SetGameModes(i []int) *GameMapUpdateOne {
	gmuo.mutation.SetGameModes(i)
	return gmuo
}

// AppendGameModes appends i to the "game_modes" field.
func (gmuo *GameMapUpdateOne) AppendGameModes(i []int) *GameMapUpdateOne {
	gmuo.mutation.AppendGameModes(i)
	return gmuo
}

// SetSupremacyPoints sets the "supremacy_points" field.
func (gmuo *GameMapUpdateOne) SetSupremacyPoints(i int) *GameMapUpdateOne {
	gmuo.mutation.ResetSupremacyPoints()
	gmuo.mutation.SetSupremacyPoints(i)
	return gmuo
}

// SetNillableSupremacyPoints sets the "supremacy_points" field if the given value is not nil.
func (gmuo *GameMapUpdateOne) SetNillableSupremacyPoints(i *int) *GameMapUpdateOne {
	if i != nil {
		gmuo.SetSupremacyPoints(*i)
	}
	return gmuo
}

// AddSupremacyPoints adds i to the "supremacy_points" field.
func (gmuo *GameMapUpdateOne) AddSupremacyPoints(i int) *GameMapUpdateOne {
	gmuo.mutation.AddSupremacyPoints(i)
	return gmuo
}

// SetLocalizedNames sets the "localized_names" field.
func (gmuo *GameMapUpdateOne) SetLocalizedNames(m map[language.Tag]string) *GameMapUpdateOne {
	gmuo.mutation.SetLocalizedNames(m)
	return gmuo
}

// Mutation returns the GameMapMutation object of the builder.
func (gmuo *GameMapUpdateOne) Mutation() *GameMapMutation {
	return gmuo.mutation
}

// Where appends a list predicates to the GameMapUpdate builder.
func (gmuo *GameMapUpdateOne) Where(ps ...predicate.GameMap) *GameMapUpdateOne {
	gmuo.mutation.Where(ps...)
	return gmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GameMapUpdateOne) Select(field string, fields ...string) *GameMapUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GameMap entity.
func (gmuo *GameMapUpdateOne) Save(ctx context.Context) (*GameMap, error) {
	gmuo.defaults()
	return withHooks(ctx, gmuo.sqlSave, gmuo.mutation, gmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GameMapUpdateOne) SaveX(ctx context.Context) *GameMap {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GameMapUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GameMapUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmuo *GameMapUpdateOne) defaults() {
	if _, ok := gmuo.mutation.UpdatedAt(); !ok {
		v := gamemap.UpdateDefaultUpdatedAt()
		gmuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gmuo *GameMapUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GameMapUpdateOne {
	gmuo.modifiers = append(gmuo.modifiers, modifiers...)
	return gmuo
}

func (gmuo *GameMapUpdateOne) sqlSave(ctx context.Context) (_node *GameMap, err error) {
	_spec := sqlgraph.NewUpdateSpec(gamemap.Table, gamemap.Columns, sqlgraph.NewFieldSpec(gamemap.FieldID, field.TypeString))
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "GameMap.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gamemap.FieldID)
		for _, f := range fields {
			if !gamemap.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != gamemap.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(gamemap.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gmuo.mutation.GameModes(); ok {
		_spec.SetField(gamemap.FieldGameModes, field.TypeJSON, value)
	}
	if value, ok := gmuo.mutation.AppendedGameModes(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, gamemap.FieldGameModes, value)
		})
	}
	if value, ok := gmuo.mutation.SupremacyPoints(); ok {
		_spec.SetField(gamemap.FieldSupremacyPoints, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.AddedSupremacyPoints(); ok {
		_spec.AddField(gamemap.FieldSupremacyPoints, field.TypeInt, value)
	}
	if value, ok := gmuo.mutation.LocalizedNames(); ok {
		_spec.SetField(gamemap.FieldLocalizedNames, field.TypeJSON, value)
	}
	_spec.AddModifiers(gmuo.modifiers...)
	_node = &GameMap{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gamemap.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmuo.mutation.done = true
	return _node, nil
}