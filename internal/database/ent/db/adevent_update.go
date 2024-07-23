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
	"github.com/cufee/aftermath/internal/database/ent/db/adevent"
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
)

// AdEventUpdate is the builder for updating AdEvent entities.
type AdEventUpdate struct {
	config
	hooks     []Hook
	mutation  *AdEventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AdEventUpdate builder.
func (aeu *AdEventUpdate) Where(ps ...predicate.AdEvent) *AdEventUpdate {
	aeu.mutation.Where(ps...)
	return aeu
}

// SetUpdatedAt sets the "updated_at" field.
func (aeu *AdEventUpdate) SetUpdatedAt(t time.Time) *AdEventUpdate {
	aeu.mutation.SetUpdatedAt(t)
	return aeu
}

// SetUserID sets the "user_id" field.
func (aeu *AdEventUpdate) SetUserID(s string) *AdEventUpdate {
	aeu.mutation.SetUserID(s)
	return aeu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (aeu *AdEventUpdate) SetNillableUserID(s *string) *AdEventUpdate {
	if s != nil {
		aeu.SetUserID(*s)
	}
	return aeu
}

// SetGuildID sets the "guild_id" field.
func (aeu *AdEventUpdate) SetGuildID(s string) *AdEventUpdate {
	aeu.mutation.SetGuildID(s)
	return aeu
}

// SetNillableGuildID sets the "guild_id" field if the given value is not nil.
func (aeu *AdEventUpdate) SetNillableGuildID(s *string) *AdEventUpdate {
	if s != nil {
		aeu.SetGuildID(*s)
	}
	return aeu
}

// SetChannelID sets the "channel_id" field.
func (aeu *AdEventUpdate) SetChannelID(s string) *AdEventUpdate {
	aeu.mutation.SetChannelID(s)
	return aeu
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (aeu *AdEventUpdate) SetNillableChannelID(s *string) *AdEventUpdate {
	if s != nil {
		aeu.SetChannelID(*s)
	}
	return aeu
}

// SetLocale sets the "locale" field.
func (aeu *AdEventUpdate) SetLocale(s string) *AdEventUpdate {
	aeu.mutation.SetLocale(s)
	return aeu
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (aeu *AdEventUpdate) SetNillableLocale(s *string) *AdEventUpdate {
	if s != nil {
		aeu.SetLocale(*s)
	}
	return aeu
}

// SetMessageID sets the "message_id" field.
func (aeu *AdEventUpdate) SetMessageID(s string) *AdEventUpdate {
	aeu.mutation.SetMessageID(s)
	return aeu
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (aeu *AdEventUpdate) SetNillableMessageID(s *string) *AdEventUpdate {
	if s != nil {
		aeu.SetMessageID(*s)
	}
	return aeu
}

// SetMetadata sets the "metadata" field.
func (aeu *AdEventUpdate) SetMetadata(m map[string]interface{}) *AdEventUpdate {
	aeu.mutation.SetMetadata(m)
	return aeu
}

// ClearMetadata clears the value of the "metadata" field.
func (aeu *AdEventUpdate) ClearMetadata() *AdEventUpdate {
	aeu.mutation.ClearMetadata()
	return aeu
}

// Mutation returns the AdEventMutation object of the builder.
func (aeu *AdEventUpdate) Mutation() *AdEventMutation {
	return aeu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aeu *AdEventUpdate) Save(ctx context.Context) (int, error) {
	aeu.defaults()
	return withHooks(ctx, aeu.sqlSave, aeu.mutation, aeu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeu *AdEventUpdate) SaveX(ctx context.Context) int {
	affected, err := aeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aeu *AdEventUpdate) Exec(ctx context.Context) error {
	_, err := aeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeu *AdEventUpdate) ExecX(ctx context.Context) {
	if err := aeu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aeu *AdEventUpdate) defaults() {
	if _, ok := aeu.mutation.UpdatedAt(); !ok {
		v := adevent.UpdateDefaultUpdatedAt()
		aeu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aeu *AdEventUpdate) check() error {
	if v, ok := aeu.mutation.UserID(); ok {
		if err := adevent.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`db: validator failed for field "AdEvent.user_id": %w`, err)}
		}
	}
	if v, ok := aeu.mutation.ChannelID(); ok {
		if err := adevent.ChannelIDValidator(v); err != nil {
			return &ValidationError{Name: "channel_id", err: fmt.Errorf(`db: validator failed for field "AdEvent.channel_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aeu *AdEventUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdEventUpdate {
	aeu.modifiers = append(aeu.modifiers, modifiers...)
	return aeu
}

func (aeu *AdEventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := aeu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(adevent.Table, adevent.Columns, sqlgraph.NewFieldSpec(adevent.FieldID, field.TypeString))
	if ps := aeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeu.mutation.UpdatedAt(); ok {
		_spec.SetField(adevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aeu.mutation.UserID(); ok {
		_spec.SetField(adevent.FieldUserID, field.TypeString, value)
	}
	if value, ok := aeu.mutation.GuildID(); ok {
		_spec.SetField(adevent.FieldGuildID, field.TypeString, value)
	}
	if value, ok := aeu.mutation.ChannelID(); ok {
		_spec.SetField(adevent.FieldChannelID, field.TypeString, value)
	}
	if value, ok := aeu.mutation.Locale(); ok {
		_spec.SetField(adevent.FieldLocale, field.TypeString, value)
	}
	if value, ok := aeu.mutation.MessageID(); ok {
		_spec.SetField(adevent.FieldMessageID, field.TypeString, value)
	}
	if value, ok := aeu.mutation.Metadata(); ok {
		_spec.SetField(adevent.FieldMetadata, field.TypeJSON, value)
	}
	if aeu.mutation.MetadataCleared() {
		_spec.ClearField(adevent.FieldMetadata, field.TypeJSON)
	}
	_spec.AddModifiers(aeu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, aeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	aeu.mutation.done = true
	return n, nil
}

// AdEventUpdateOne is the builder for updating a single AdEvent entity.
type AdEventUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AdEventMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (aeuo *AdEventUpdateOne) SetUpdatedAt(t time.Time) *AdEventUpdateOne {
	aeuo.mutation.SetUpdatedAt(t)
	return aeuo
}

// SetUserID sets the "user_id" field.
func (aeuo *AdEventUpdateOne) SetUserID(s string) *AdEventUpdateOne {
	aeuo.mutation.SetUserID(s)
	return aeuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (aeuo *AdEventUpdateOne) SetNillableUserID(s *string) *AdEventUpdateOne {
	if s != nil {
		aeuo.SetUserID(*s)
	}
	return aeuo
}

// SetGuildID sets the "guild_id" field.
func (aeuo *AdEventUpdateOne) SetGuildID(s string) *AdEventUpdateOne {
	aeuo.mutation.SetGuildID(s)
	return aeuo
}

// SetNillableGuildID sets the "guild_id" field if the given value is not nil.
func (aeuo *AdEventUpdateOne) SetNillableGuildID(s *string) *AdEventUpdateOne {
	if s != nil {
		aeuo.SetGuildID(*s)
	}
	return aeuo
}

// SetChannelID sets the "channel_id" field.
func (aeuo *AdEventUpdateOne) SetChannelID(s string) *AdEventUpdateOne {
	aeuo.mutation.SetChannelID(s)
	return aeuo
}

// SetNillableChannelID sets the "channel_id" field if the given value is not nil.
func (aeuo *AdEventUpdateOne) SetNillableChannelID(s *string) *AdEventUpdateOne {
	if s != nil {
		aeuo.SetChannelID(*s)
	}
	return aeuo
}

// SetLocale sets the "locale" field.
func (aeuo *AdEventUpdateOne) SetLocale(s string) *AdEventUpdateOne {
	aeuo.mutation.SetLocale(s)
	return aeuo
}

// SetNillableLocale sets the "locale" field if the given value is not nil.
func (aeuo *AdEventUpdateOne) SetNillableLocale(s *string) *AdEventUpdateOne {
	if s != nil {
		aeuo.SetLocale(*s)
	}
	return aeuo
}

// SetMessageID sets the "message_id" field.
func (aeuo *AdEventUpdateOne) SetMessageID(s string) *AdEventUpdateOne {
	aeuo.mutation.SetMessageID(s)
	return aeuo
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (aeuo *AdEventUpdateOne) SetNillableMessageID(s *string) *AdEventUpdateOne {
	if s != nil {
		aeuo.SetMessageID(*s)
	}
	return aeuo
}

// SetMetadata sets the "metadata" field.
func (aeuo *AdEventUpdateOne) SetMetadata(m map[string]interface{}) *AdEventUpdateOne {
	aeuo.mutation.SetMetadata(m)
	return aeuo
}

// ClearMetadata clears the value of the "metadata" field.
func (aeuo *AdEventUpdateOne) ClearMetadata() *AdEventUpdateOne {
	aeuo.mutation.ClearMetadata()
	return aeuo
}

// Mutation returns the AdEventMutation object of the builder.
func (aeuo *AdEventUpdateOne) Mutation() *AdEventMutation {
	return aeuo.mutation
}

// Where appends a list predicates to the AdEventUpdate builder.
func (aeuo *AdEventUpdateOne) Where(ps ...predicate.AdEvent) *AdEventUpdateOne {
	aeuo.mutation.Where(ps...)
	return aeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aeuo *AdEventUpdateOne) Select(field string, fields ...string) *AdEventUpdateOne {
	aeuo.fields = append([]string{field}, fields...)
	return aeuo
}

// Save executes the query and returns the updated AdEvent entity.
func (aeuo *AdEventUpdateOne) Save(ctx context.Context) (*AdEvent, error) {
	aeuo.defaults()
	return withHooks(ctx, aeuo.sqlSave, aeuo.mutation, aeuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (aeuo *AdEventUpdateOne) SaveX(ctx context.Context) *AdEvent {
	node, err := aeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aeuo *AdEventUpdateOne) Exec(ctx context.Context) error {
	_, err := aeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aeuo *AdEventUpdateOne) ExecX(ctx context.Context) {
	if err := aeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aeuo *AdEventUpdateOne) defaults() {
	if _, ok := aeuo.mutation.UpdatedAt(); !ok {
		v := adevent.UpdateDefaultUpdatedAt()
		aeuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (aeuo *AdEventUpdateOne) check() error {
	if v, ok := aeuo.mutation.UserID(); ok {
		if err := adevent.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`db: validator failed for field "AdEvent.user_id": %w`, err)}
		}
	}
	if v, ok := aeuo.mutation.ChannelID(); ok {
		if err := adevent.ChannelIDValidator(v); err != nil {
			return &ValidationError{Name: "channel_id", err: fmt.Errorf(`db: validator failed for field "AdEvent.channel_id": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aeuo *AdEventUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AdEventUpdateOne {
	aeuo.modifiers = append(aeuo.modifiers, modifiers...)
	return aeuo
}

func (aeuo *AdEventUpdateOne) sqlSave(ctx context.Context) (_node *AdEvent, err error) {
	if err := aeuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(adevent.Table, adevent.Columns, sqlgraph.NewFieldSpec(adevent.FieldID, field.TypeString))
	id, ok := aeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AdEvent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adevent.FieldID)
		for _, f := range fields {
			if !adevent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != adevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(adevent.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := aeuo.mutation.UserID(); ok {
		_spec.SetField(adevent.FieldUserID, field.TypeString, value)
	}
	if value, ok := aeuo.mutation.GuildID(); ok {
		_spec.SetField(adevent.FieldGuildID, field.TypeString, value)
	}
	if value, ok := aeuo.mutation.ChannelID(); ok {
		_spec.SetField(adevent.FieldChannelID, field.TypeString, value)
	}
	if value, ok := aeuo.mutation.Locale(); ok {
		_spec.SetField(adevent.FieldLocale, field.TypeString, value)
	}
	if value, ok := aeuo.mutation.MessageID(); ok {
		_spec.SetField(adevent.FieldMessageID, field.TypeString, value)
	}
	if value, ok := aeuo.mutation.Metadata(); ok {
		_spec.SetField(adevent.FieldMetadata, field.TypeJSON, value)
	}
	if aeuo.mutation.MetadataCleared() {
		_spec.ClearField(adevent.FieldMetadata, field.TypeJSON)
	}
	_spec.AddModifiers(aeuo.modifiers...)
	_node = &AdEvent{config: aeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adevent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	aeuo.mutation.done = true
	return _node, nil
}