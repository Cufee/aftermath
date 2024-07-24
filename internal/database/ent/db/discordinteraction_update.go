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
	"github.com/cufee/aftermath/internal/database/ent/db/discordinteraction"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/models"
)

// DiscordInteractionUpdate is the builder for updating DiscordInteraction entities.
type DiscordInteractionUpdate struct {
	config
	hooks     []Hook
	mutation  *DiscordInteractionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DiscordInteractionUpdate builder.
func (diu *DiscordInteractionUpdate) Where(ps ...predicate.DiscordInteraction) *DiscordInteractionUpdate {
	diu.mutation.Where(ps...)
	return diu
}

// SetUpdatedAt sets the "updated_at" field.
func (diu *DiscordInteractionUpdate) SetUpdatedAt(t time.Time) *DiscordInteractionUpdate {
	diu.mutation.SetUpdatedAt(t)
	return diu
}

// SetType sets the "type" field.
func (diu *DiscordInteractionUpdate) SetType(mit models.DiscordInteractionType) *DiscordInteractionUpdate {
	diu.mutation.SetType(mit)
	return diu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (diu *DiscordInteractionUpdate) SetNillableType(mit *models.DiscordInteractionType) *DiscordInteractionUpdate {
	if mit != nil {
		diu.SetType(*mit)
	}
	return diu
}

// SetLocale sets the "locale" field.
func (diu *DiscordInteractionUpdate) SetLocale(s string) *DiscordInteractionUpdate {
	diu.mutation.SetLocale(s)
	return diu
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (diu *DiscordInteractionUpdate) SetNillableLocale(s *string) *DiscordInteractionUpdate {
	if s != nil {
		diu.SetLocale(*s)
	}
	return diu
}

// SetMetadata sets the "metadata" field.
func (diu *DiscordInteractionUpdate) SetMetadata(m map[string]interface{}) *DiscordInteractionUpdate {
	diu.mutation.SetMetadata(m)
	return diu
}

// Mutation returns the DiscordInteractionMutation object of the builder.
func (diu *DiscordInteractionUpdate) Mutation() *DiscordInteractionMutation {
	return diu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (diu *DiscordInteractionUpdate) Save(ctx context.Context) (int, error) {
	diu.defaults()
	return withHooks(ctx, diu.sqlSave, diu.mutation, diu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (diu *DiscordInteractionUpdate) SaveX(ctx context.Context) int {
	affected, err := diu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (diu *DiscordInteractionUpdate) Exec(ctx context.Context) error {
	_, err := diu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diu *DiscordInteractionUpdate) ExecX(ctx context.Context) {
	if err := diu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diu *DiscordInteractionUpdate) defaults() {
	if _, ok := diu.mutation.UpdatedAt(); !ok {
		v := discordinteraction.UpdateDefaultUpdatedAt()
		diu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (diu *DiscordInteractionUpdate) check() error {
	if v, ok := diu.mutation.GetType(); ok {
		if err := discordinteraction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "DiscordInteraction.type": %w`, err)}
		}
	}
	if _, ok := diu.mutation.UserID(); diu.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "DiscordInteraction.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (diu *DiscordInteractionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DiscordInteractionUpdate {
	diu.modifiers = append(diu.modifiers, modifiers...)
	return diu
}

func (diu *DiscordInteractionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := diu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(discordinteraction.Table, discordinteraction.Columns, sqlgraph.NewFieldSpec(discordinteraction.FieldID, field.TypeString))
	if ps := diu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diu.mutation.UpdatedAt(); ok {
		_spec.SetField(discordinteraction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := diu.mutation.GetType(); ok {
		_spec.SetField(discordinteraction.FieldType, field.TypeEnum, value)
	}
	if value, ok := diu.mutation.Locale(); ok {
		_spec.SetField(discordinteraction.FieldLocale, field.TypeString, value)
	}
	if value, ok := diu.mutation.Metadata(); ok {
		_spec.SetField(discordinteraction.FieldMetadata, field.TypeJSON, value)
	}
	_spec.AddModifiers(diu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, diu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordinteraction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	diu.mutation.done = true
	return n, nil
}

// DiscordInteractionUpdateOne is the builder for updating a single DiscordInteraction entity.
type DiscordInteractionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DiscordInteractionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (diuo *DiscordInteractionUpdateOne) SetUpdatedAt(t time.Time) *DiscordInteractionUpdateOne {
	diuo.mutation.SetUpdatedAt(t)
	return diuo
}

// SetType sets the "type" field.
func (diuo *DiscordInteractionUpdateOne) SetType(mit models.DiscordInteractionType) *DiscordInteractionUpdateOne {
	diuo.mutation.SetType(mit)
	return diuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (diuo *DiscordInteractionUpdateOne) SetNillableType(mit *models.DiscordInteractionType) *DiscordInteractionUpdateOne {
	if mit != nil {
		diuo.SetType(*mit)
	}
	return diuo
}

// SetLocale sets the "locale" field.
func (diuo *DiscordInteractionUpdateOne) SetLocale(s string) *DiscordInteractionUpdateOne {
	diuo.mutation.SetLocale(s)
	return diuo
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (diuo *DiscordInteractionUpdateOne) SetNillableLocale(s *string) *DiscordInteractionUpdateOne {
	if s != nil {
		diuo.SetLocale(*s)
	}
	return diuo
}

// SetMetadata sets the "metadata" field.
func (diuo *DiscordInteractionUpdateOne) SetMetadata(m map[string]interface{}) *DiscordInteractionUpdateOne {
	diuo.mutation.SetMetadata(m)
	return diuo
}

// Mutation returns the DiscordInteractionMutation object of the builder.
func (diuo *DiscordInteractionUpdateOne) Mutation() *DiscordInteractionMutation {
	return diuo.mutation
}

// Where appends a list predicates to the DiscordInteractionUpdate builder.
func (diuo *DiscordInteractionUpdateOne) Where(ps ...predicate.DiscordInteraction) *DiscordInteractionUpdateOne {
	diuo.mutation.Where(ps...)
	return diuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (diuo *DiscordInteractionUpdateOne) Select(field string, fields ...string) *DiscordInteractionUpdateOne {
	diuo.fields = append([]string{field}, fields...)
	return diuo
}

// Save executes the query and returns the updated DiscordInteraction entity.
func (diuo *DiscordInteractionUpdateOne) Save(ctx context.Context) (*DiscordInteraction, error) {
	diuo.defaults()
	return withHooks(ctx, diuo.sqlSave, diuo.mutation, diuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (diuo *DiscordInteractionUpdateOne) SaveX(ctx context.Context) *DiscordInteraction {
	node, err := diuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (diuo *DiscordInteractionUpdateOne) Exec(ctx context.Context) error {
	_, err := diuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (diuo *DiscordInteractionUpdateOne) ExecX(ctx context.Context) {
	if err := diuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (diuo *DiscordInteractionUpdateOne) defaults() {
	if _, ok := diuo.mutation.UpdatedAt(); !ok {
		v := discordinteraction.UpdateDefaultUpdatedAt()
		diuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (diuo *DiscordInteractionUpdateOne) check() error {
	if v, ok := diuo.mutation.GetType(); ok {
		if err := discordinteraction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "DiscordInteraction.type": %w`, err)}
		}
	}
	if _, ok := diuo.mutation.UserID(); diuo.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "DiscordInteraction.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (diuo *DiscordInteractionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DiscordInteractionUpdateOne {
	diuo.modifiers = append(diuo.modifiers, modifiers...)
	return diuo
}

func (diuo *DiscordInteractionUpdateOne) sqlSave(ctx context.Context) (_node *DiscordInteraction, err error) {
	if err := diuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(discordinteraction.Table, discordinteraction.Columns, sqlgraph.NewFieldSpec(discordinteraction.FieldID, field.TypeString))
	id, ok := diuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "DiscordInteraction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := diuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, discordinteraction.FieldID)
		for _, f := range fields {
			if !discordinteraction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != discordinteraction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := diuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := diuo.mutation.UpdatedAt(); ok {
		_spec.SetField(discordinteraction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := diuo.mutation.GetType(); ok {
		_spec.SetField(discordinteraction.FieldType, field.TypeEnum, value)
	}
	if value, ok := diuo.mutation.Locale(); ok {
		_spec.SetField(discordinteraction.FieldLocale, field.TypeString, value)
	}
	if value, ok := diuo.mutation.Metadata(); ok {
		_spec.SetField(discordinteraction.FieldMetadata, field.TypeJSON, value)
	}
	_spec.AddModifiers(diuo.modifiers...)
	_node = &DiscordInteraction{config: diuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, diuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{discordinteraction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	diuo.mutation.done = true
	return _node, nil
}
