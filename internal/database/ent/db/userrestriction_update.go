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
	"github.com/cufee/aftermath/internal/database/ent/db/predicate"
	"github.com/cufee/aftermath/internal/database/ent/db/userrestriction"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserRestrictionUpdate is the builder for updating UserRestriction entities.
type UserRestrictionUpdate struct {
	config
	hooks     []Hook
	mutation  *UserRestrictionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserRestrictionUpdate builder.
func (uru *UserRestrictionUpdate) Where(ps ...predicate.UserRestriction) *UserRestrictionUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetUpdatedAt sets the "updated_at" field.
func (uru *UserRestrictionUpdate) SetUpdatedAt(t time.Time) *UserRestrictionUpdate {
	uru.mutation.SetUpdatedAt(t)
	return uru
}

// SetExpiresAt sets the "expires_at" field.
func (uru *UserRestrictionUpdate) SetExpiresAt(t time.Time) *UserRestrictionUpdate {
	uru.mutation.SetExpiresAt(t)
	return uru
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (uru *UserRestrictionUpdate) SetNillableExpiresAt(t *time.Time) *UserRestrictionUpdate {
	if t != nil {
		uru.SetExpiresAt(*t)
	}
	return uru
}

// SetType sets the "type" field.
func (uru *UserRestrictionUpdate) SetType(mrt models.UserRestrictionType) *UserRestrictionUpdate {
	uru.mutation.SetType(mrt)
	return uru
}

// SetNillableType sets the "type" field if the given value is not nil.
func (uru *UserRestrictionUpdate) SetNillableType(mrt *models.UserRestrictionType) *UserRestrictionUpdate {
	if mrt != nil {
		uru.SetType(*mrt)
	}
	return uru
}

// SetRestriction sets the "restriction" field.
func (uru *UserRestrictionUpdate) SetRestriction(s string) *UserRestrictionUpdate {
	uru.mutation.SetRestriction(s)
	return uru
}

// SetNillableRestriction sets the "restriction" field if the given value is not nil.
func (uru *UserRestrictionUpdate) SetNillableRestriction(s *string) *UserRestrictionUpdate {
	if s != nil {
		uru.SetRestriction(*s)
	}
	return uru
}

// SetPublicReason sets the "public_reason" field.
func (uru *UserRestrictionUpdate) SetPublicReason(s string) *UserRestrictionUpdate {
	uru.mutation.SetPublicReason(s)
	return uru
}

// SetNillablePublicReason sets the "public_reason" field if the given value is not nil.
func (uru *UserRestrictionUpdate) SetNillablePublicReason(s *string) *UserRestrictionUpdate {
	if s != nil {
		uru.SetPublicReason(*s)
	}
	return uru
}

// SetModeratorComment sets the "moderator_comment" field.
func (uru *UserRestrictionUpdate) SetModeratorComment(s string) *UserRestrictionUpdate {
	uru.mutation.SetModeratorComment(s)
	return uru
}

// SetNillableModeratorComment sets the "moderator_comment" field if the given value is not nil.
func (uru *UserRestrictionUpdate) SetNillableModeratorComment(s *string) *UserRestrictionUpdate {
	if s != nil {
		uru.SetModeratorComment(*s)
	}
	return uru
}

// SetEvents sets the "events" field.
func (uru *UserRestrictionUpdate) SetEvents(mu []models.RestrictionUpdate) *UserRestrictionUpdate {
	uru.mutation.SetEvents(mu)
	return uru
}

// AppendEvents appends mu to the "events" field.
func (uru *UserRestrictionUpdate) AppendEvents(mu []models.RestrictionUpdate) *UserRestrictionUpdate {
	uru.mutation.AppendEvents(mu)
	return uru
}

// Mutation returns the UserRestrictionMutation object of the builder.
func (uru *UserRestrictionUpdate) Mutation() *UserRestrictionMutation {
	return uru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserRestrictionUpdate) Save(ctx context.Context) (int, error) {
	uru.defaults()
	return withHooks(ctx, uru.sqlSave, uru.mutation, uru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserRestrictionUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserRestrictionUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserRestrictionUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uru *UserRestrictionUpdate) defaults() {
	if _, ok := uru.mutation.UpdatedAt(); !ok {
		v := userrestriction.UpdateDefaultUpdatedAt()
		uru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uru *UserRestrictionUpdate) check() error {
	if v, ok := uru.mutation.GetType(); ok {
		if err := userrestriction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserRestriction.type": %w`, err)}
		}
	}
	if v, ok := uru.mutation.Restriction(); ok {
		if err := userrestriction.RestrictionValidator(v); err != nil {
			return &ValidationError{Name: "restriction", err: fmt.Errorf(`db: validator failed for field "UserRestriction.restriction": %w`, err)}
		}
	}
	if v, ok := uru.mutation.PublicReason(); ok {
		if err := userrestriction.PublicReasonValidator(v); err != nil {
			return &ValidationError{Name: "public_reason", err: fmt.Errorf(`db: validator failed for field "UserRestriction.public_reason": %w`, err)}
		}
	}
	if _, ok := uru.mutation.UserID(); uru.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "UserRestriction.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uru *UserRestrictionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserRestrictionUpdate {
	uru.modifiers = append(uru.modifiers, modifiers...)
	return uru
}

func (uru *UserRestrictionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userrestriction.Table, userrestriction.Columns, sqlgraph.NewFieldSpec(userrestriction.FieldID, field.TypeString))
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.UpdatedAt(); ok {
		_spec.SetField(userrestriction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uru.mutation.ExpiresAt(); ok {
		_spec.SetField(userrestriction.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := uru.mutation.GetType(); ok {
		_spec.SetField(userrestriction.FieldType, field.TypeEnum, value)
	}
	if value, ok := uru.mutation.Restriction(); ok {
		_spec.SetField(userrestriction.FieldRestriction, field.TypeString, value)
	}
	if value, ok := uru.mutation.PublicReason(); ok {
		_spec.SetField(userrestriction.FieldPublicReason, field.TypeString, value)
	}
	if value, ok := uru.mutation.ModeratorComment(); ok {
		_spec.SetField(userrestriction.FieldModeratorComment, field.TypeString, value)
	}
	if value, ok := uru.mutation.Events(); ok {
		_spec.SetField(userrestriction.FieldEvents, field.TypeJSON, value)
	}
	if value, ok := uru.mutation.AppendedEvents(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, userrestriction.FieldEvents, value)
		})
	}
	_spec.AddModifiers(uru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrestriction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uru.mutation.done = true
	return n, nil
}

// UserRestrictionUpdateOne is the builder for updating a single UserRestriction entity.
type UserRestrictionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserRestrictionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (uruo *UserRestrictionUpdateOne) SetUpdatedAt(t time.Time) *UserRestrictionUpdateOne {
	uruo.mutation.SetUpdatedAt(t)
	return uruo
}

// SetExpiresAt sets the "expires_at" field.
func (uruo *UserRestrictionUpdateOne) SetExpiresAt(t time.Time) *UserRestrictionUpdateOne {
	uruo.mutation.SetExpiresAt(t)
	return uruo
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (uruo *UserRestrictionUpdateOne) SetNillableExpiresAt(t *time.Time) *UserRestrictionUpdateOne {
	if t != nil {
		uruo.SetExpiresAt(*t)
	}
	return uruo
}

// SetType sets the "type" field.
func (uruo *UserRestrictionUpdateOne) SetType(mrt models.UserRestrictionType) *UserRestrictionUpdateOne {
	uruo.mutation.SetType(mrt)
	return uruo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (uruo *UserRestrictionUpdateOne) SetNillableType(mrt *models.UserRestrictionType) *UserRestrictionUpdateOne {
	if mrt != nil {
		uruo.SetType(*mrt)
	}
	return uruo
}

// SetRestriction sets the "restriction" field.
func (uruo *UserRestrictionUpdateOne) SetRestriction(s string) *UserRestrictionUpdateOne {
	uruo.mutation.SetRestriction(s)
	return uruo
}

// SetNillableRestriction sets the "restriction" field if the given value is not nil.
func (uruo *UserRestrictionUpdateOne) SetNillableRestriction(s *string) *UserRestrictionUpdateOne {
	if s != nil {
		uruo.SetRestriction(*s)
	}
	return uruo
}

// SetPublicReason sets the "public_reason" field.
func (uruo *UserRestrictionUpdateOne) SetPublicReason(s string) *UserRestrictionUpdateOne {
	uruo.mutation.SetPublicReason(s)
	return uruo
}

// SetNillablePublicReason sets the "public_reason" field if the given value is not nil.
func (uruo *UserRestrictionUpdateOne) SetNillablePublicReason(s *string) *UserRestrictionUpdateOne {
	if s != nil {
		uruo.SetPublicReason(*s)
	}
	return uruo
}

// SetModeratorComment sets the "moderator_comment" field.
func (uruo *UserRestrictionUpdateOne) SetModeratorComment(s string) *UserRestrictionUpdateOne {
	uruo.mutation.SetModeratorComment(s)
	return uruo
}

// SetNillableModeratorComment sets the "moderator_comment" field if the given value is not nil.
func (uruo *UserRestrictionUpdateOne) SetNillableModeratorComment(s *string) *UserRestrictionUpdateOne {
	if s != nil {
		uruo.SetModeratorComment(*s)
	}
	return uruo
}

// SetEvents sets the "events" field.
func (uruo *UserRestrictionUpdateOne) SetEvents(mu []models.RestrictionUpdate) *UserRestrictionUpdateOne {
	uruo.mutation.SetEvents(mu)
	return uruo
}

// AppendEvents appends mu to the "events" field.
func (uruo *UserRestrictionUpdateOne) AppendEvents(mu []models.RestrictionUpdate) *UserRestrictionUpdateOne {
	uruo.mutation.AppendEvents(mu)
	return uruo
}

// Mutation returns the UserRestrictionMutation object of the builder.
func (uruo *UserRestrictionUpdateOne) Mutation() *UserRestrictionMutation {
	return uruo.mutation
}

// Where appends a list predicates to the UserRestrictionUpdate builder.
func (uruo *UserRestrictionUpdateOne) Where(ps ...predicate.UserRestriction) *UserRestrictionUpdateOne {
	uruo.mutation.Where(ps...)
	return uruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UserRestrictionUpdateOne) Select(field string, fields ...string) *UserRestrictionUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UserRestriction entity.
func (uruo *UserRestrictionUpdateOne) Save(ctx context.Context) (*UserRestriction, error) {
	uruo.defaults()
	return withHooks(ctx, uruo.sqlSave, uruo.mutation, uruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserRestrictionUpdateOne) SaveX(ctx context.Context) *UserRestriction {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserRestrictionUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserRestrictionUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uruo *UserRestrictionUpdateOne) defaults() {
	if _, ok := uruo.mutation.UpdatedAt(); !ok {
		v := userrestriction.UpdateDefaultUpdatedAt()
		uruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uruo *UserRestrictionUpdateOne) check() error {
	if v, ok := uruo.mutation.GetType(); ok {
		if err := userrestriction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserRestriction.type": %w`, err)}
		}
	}
	if v, ok := uruo.mutation.Restriction(); ok {
		if err := userrestriction.RestrictionValidator(v); err != nil {
			return &ValidationError{Name: "restriction", err: fmt.Errorf(`db: validator failed for field "UserRestriction.restriction": %w`, err)}
		}
	}
	if v, ok := uruo.mutation.PublicReason(); ok {
		if err := userrestriction.PublicReasonValidator(v); err != nil {
			return &ValidationError{Name: "public_reason", err: fmt.Errorf(`db: validator failed for field "UserRestriction.public_reason": %w`, err)}
		}
	}
	if _, ok := uruo.mutation.UserID(); uruo.mutation.UserCleared() && !ok {
		return errors.New(`db: clearing a required unique edge "UserRestriction.user"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uruo *UserRestrictionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserRestrictionUpdateOne {
	uruo.modifiers = append(uruo.modifiers, modifiers...)
	return uruo
}

func (uruo *UserRestrictionUpdateOne) sqlSave(ctx context.Context) (_node *UserRestriction, err error) {
	if err := uruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userrestriction.Table, userrestriction.Columns, sqlgraph.NewFieldSpec(userrestriction.FieldID, field.TypeString))
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "UserRestriction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userrestriction.FieldID)
		for _, f := range fields {
			if !userrestriction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != userrestriction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uruo.mutation.UpdatedAt(); ok {
		_spec.SetField(userrestriction.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uruo.mutation.ExpiresAt(); ok {
		_spec.SetField(userrestriction.FieldExpiresAt, field.TypeTime, value)
	}
	if value, ok := uruo.mutation.GetType(); ok {
		_spec.SetField(userrestriction.FieldType, field.TypeEnum, value)
	}
	if value, ok := uruo.mutation.Restriction(); ok {
		_spec.SetField(userrestriction.FieldRestriction, field.TypeString, value)
	}
	if value, ok := uruo.mutation.PublicReason(); ok {
		_spec.SetField(userrestriction.FieldPublicReason, field.TypeString, value)
	}
	if value, ok := uruo.mutation.ModeratorComment(); ok {
		_spec.SetField(userrestriction.FieldModeratorComment, field.TypeString, value)
	}
	if value, ok := uruo.mutation.Events(); ok {
		_spec.SetField(userrestriction.FieldEvents, field.TypeJSON, value)
	}
	if value, ok := uruo.mutation.AppendedEvents(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, userrestriction.FieldEvents, value)
		})
	}
	_spec.AddModifiers(uruo.modifiers...)
	_node = &UserRestriction{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrestriction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uruo.mutation.done = true
	return _node, nil
}
