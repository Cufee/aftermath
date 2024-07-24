// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/userrestriction"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserRestrictionCreate is the builder for creating a UserRestriction entity.
type UserRestrictionCreate struct {
	config
	mutation *UserRestrictionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (urc *UserRestrictionCreate) SetCreatedAt(t time.Time) *UserRestrictionCreate {
	urc.mutation.SetCreatedAt(t)
	return urc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (urc *UserRestrictionCreate) SetNillableCreatedAt(t *time.Time) *UserRestrictionCreate {
	if t != nil {
		urc.SetCreatedAt(*t)
	}
	return urc
}

// SetUpdatedAt sets the "updated_at" field.
func (urc *UserRestrictionCreate) SetUpdatedAt(t time.Time) *UserRestrictionCreate {
	urc.mutation.SetUpdatedAt(t)
	return urc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (urc *UserRestrictionCreate) SetNillableUpdatedAt(t *time.Time) *UserRestrictionCreate {
	if t != nil {
		urc.SetUpdatedAt(*t)
	}
	return urc
}

// SetExpiresAt sets the "expires_at" field.
func (urc *UserRestrictionCreate) SetExpiresAt(t time.Time) *UserRestrictionCreate {
	urc.mutation.SetExpiresAt(t)
	return urc
}

// SetType sets the "type" field.
func (urc *UserRestrictionCreate) SetType(mrt models.UserRestrictionType) *UserRestrictionCreate {
	urc.mutation.SetType(mrt)
	return urc
}

// SetUserID sets the "user_id" field.
func (urc *UserRestrictionCreate) SetUserID(s string) *UserRestrictionCreate {
	urc.mutation.SetUserID(s)
	return urc
}

// SetRestriction sets the "restriction" field.
func (urc *UserRestrictionCreate) SetRestriction(s string) *UserRestrictionCreate {
	urc.mutation.SetRestriction(s)
	return urc
}

// SetPublicReason sets the "public_reason" field.
func (urc *UserRestrictionCreate) SetPublicReason(s string) *UserRestrictionCreate {
	urc.mutation.SetPublicReason(s)
	return urc
}

// SetModeratorComment sets the "moderator_comment" field.
func (urc *UserRestrictionCreate) SetModeratorComment(s string) *UserRestrictionCreate {
	urc.mutation.SetModeratorComment(s)
	return urc
}

// SetEvents sets the "events" field.
func (urc *UserRestrictionCreate) SetEvents(mu []models.RestrictionUpdate) *UserRestrictionCreate {
	urc.mutation.SetEvents(mu)
	return urc
}

// SetID sets the "id" field.
func (urc *UserRestrictionCreate) SetID(s string) *UserRestrictionCreate {
	urc.mutation.SetID(s)
	return urc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (urc *UserRestrictionCreate) SetNillableID(s *string) *UserRestrictionCreate {
	if s != nil {
		urc.SetID(*s)
	}
	return urc
}

// SetUser sets the "user" edge to the User entity.
func (urc *UserRestrictionCreate) SetUser(u *User) *UserRestrictionCreate {
	return urc.SetUserID(u.ID)
}

// Mutation returns the UserRestrictionMutation object of the builder.
func (urc *UserRestrictionCreate) Mutation() *UserRestrictionMutation {
	return urc.mutation
}

// Save creates the UserRestriction in the database.
func (urc *UserRestrictionCreate) Save(ctx context.Context) (*UserRestriction, error) {
	urc.defaults()
	return withHooks(ctx, urc.sqlSave, urc.mutation, urc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (urc *UserRestrictionCreate) SaveX(ctx context.Context) *UserRestriction {
	v, err := urc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urc *UserRestrictionCreate) Exec(ctx context.Context) error {
	_, err := urc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urc *UserRestrictionCreate) ExecX(ctx context.Context) {
	if err := urc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (urc *UserRestrictionCreate) defaults() {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		v := userrestriction.DefaultCreatedAt()
		urc.mutation.SetCreatedAt(v)
	}
	if _, ok := urc.mutation.UpdatedAt(); !ok {
		v := userrestriction.DefaultUpdatedAt()
		urc.mutation.SetUpdatedAt(v)
	}
	if _, ok := urc.mutation.ID(); !ok {
		v := userrestriction.DefaultID()
		urc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (urc *UserRestrictionCreate) check() error {
	if _, ok := urc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "UserRestriction.created_at"`)}
	}
	if _, ok := urc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "UserRestriction.updated_at"`)}
	}
	if _, ok := urc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`db: missing required field "UserRestriction.expires_at"`)}
	}
	if _, ok := urc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "UserRestriction.type"`)}
	}
	if v, ok := urc.mutation.GetType(); ok {
		if err := userrestriction.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserRestriction.type": %w`, err)}
		}
	}
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`db: missing required field "UserRestriction.user_id"`)}
	}
	if v, ok := urc.mutation.UserID(); ok {
		if err := userrestriction.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`db: validator failed for field "UserRestriction.user_id": %w`, err)}
		}
	}
	if _, ok := urc.mutation.Restriction(); !ok {
		return &ValidationError{Name: "restriction", err: errors.New(`db: missing required field "UserRestriction.restriction"`)}
	}
	if v, ok := urc.mutation.Restriction(); ok {
		if err := userrestriction.RestrictionValidator(v); err != nil {
			return &ValidationError{Name: "restriction", err: fmt.Errorf(`db: validator failed for field "UserRestriction.restriction": %w`, err)}
		}
	}
	if _, ok := urc.mutation.PublicReason(); !ok {
		return &ValidationError{Name: "public_reason", err: errors.New(`db: missing required field "UserRestriction.public_reason"`)}
	}
	if v, ok := urc.mutation.PublicReason(); ok {
		if err := userrestriction.PublicReasonValidator(v); err != nil {
			return &ValidationError{Name: "public_reason", err: fmt.Errorf(`db: validator failed for field "UserRestriction.public_reason": %w`, err)}
		}
	}
	if _, ok := urc.mutation.ModeratorComment(); !ok {
		return &ValidationError{Name: "moderator_comment", err: errors.New(`db: missing required field "UserRestriction.moderator_comment"`)}
	}
	if _, ok := urc.mutation.Events(); !ok {
		return &ValidationError{Name: "events", err: errors.New(`db: missing required field "UserRestriction.events"`)}
	}
	if _, ok := urc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`db: missing required edge "UserRestriction.user"`)}
	}
	return nil
}

func (urc *UserRestrictionCreate) sqlSave(ctx context.Context) (*UserRestriction, error) {
	if err := urc.check(); err != nil {
		return nil, err
	}
	_node, _spec := urc.createSpec()
	if err := sqlgraph.CreateNode(ctx, urc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserRestriction.ID type: %T", _spec.ID.Value)
		}
	}
	urc.mutation.id = &_node.ID
	urc.mutation.done = true
	return _node, nil
}

func (urc *UserRestrictionCreate) createSpec() (*UserRestriction, *sqlgraph.CreateSpec) {
	var (
		_node = &UserRestriction{config: urc.config}
		_spec = sqlgraph.NewCreateSpec(userrestriction.Table, sqlgraph.NewFieldSpec(userrestriction.FieldID, field.TypeString))
	)
	if id, ok := urc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := urc.mutation.CreatedAt(); ok {
		_spec.SetField(userrestriction.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := urc.mutation.UpdatedAt(); ok {
		_spec.SetField(userrestriction.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := urc.mutation.ExpiresAt(); ok {
		_spec.SetField(userrestriction.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := urc.mutation.GetType(); ok {
		_spec.SetField(userrestriction.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := urc.mutation.Restriction(); ok {
		_spec.SetField(userrestriction.FieldRestriction, field.TypeString, value)
		_node.Restriction = value
	}
	if value, ok := urc.mutation.PublicReason(); ok {
		_spec.SetField(userrestriction.FieldPublicReason, field.TypeString, value)
		_node.PublicReason = value
	}
	if value, ok := urc.mutation.ModeratorComment(); ok {
		_spec.SetField(userrestriction.FieldModeratorComment, field.TypeString, value)
		_node.ModeratorComment = value
	}
	if value, ok := urc.mutation.Events(); ok {
		_spec.SetField(userrestriction.FieldEvents, field.TypeJSON, value)
		_node.Events = value
	}
	if nodes := urc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrestriction.UserTable,
			Columns: []string{userrestriction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserRestrictionCreateBulk is the builder for creating many UserRestriction entities in bulk.
type UserRestrictionCreateBulk struct {
	config
	err      error
	builders []*UserRestrictionCreate
}

// Save creates the UserRestriction entities in the database.
func (urcb *UserRestrictionCreateBulk) Save(ctx context.Context) ([]*UserRestriction, error) {
	if urcb.err != nil {
		return nil, urcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(urcb.builders))
	nodes := make([]*UserRestriction, len(urcb.builders))
	mutators := make([]Mutator, len(urcb.builders))
	for i := range urcb.builders {
		func(i int, root context.Context) {
			builder := urcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserRestrictionMutation)
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
					_, err = mutators[i+1].Mutate(root, urcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, urcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, urcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (urcb *UserRestrictionCreateBulk) SaveX(ctx context.Context) []*UserRestriction {
	v, err := urcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (urcb *UserRestrictionCreateBulk) Exec(ctx context.Context) error {
	_, err := urcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (urcb *UserRestrictionCreateBulk) ExecX(ctx context.Context) {
	if err := urcb.Exec(ctx); err != nil {
		panic(err)
	}
}
