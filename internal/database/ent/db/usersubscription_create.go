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
	"github.com/cufee/aftermath/internal/database/ent/db/usersubscription"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserSubscriptionCreate is the builder for creating a UserSubscription entity.
type UserSubscriptionCreate struct {
	config
	mutation *UserSubscriptionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserSubscriptionCreate) SetCreatedAt(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableCreatedAt(t *time.Time) *UserSubscriptionCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UserSubscriptionCreate) SetUpdatedAt(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetUpdatedAt(t)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableUpdatedAt(t *time.Time) *UserSubscriptionCreate {
	if t != nil {
		usc.SetUpdatedAt(*t)
	}
	return usc
}

// SetType sets the "type" field.
func (usc *UserSubscriptionCreate) SetType(mt models.SubscriptionType) *UserSubscriptionCreate {
	usc.mutation.SetType(mt)
	return usc
}

// SetExpiresAt sets the "expires_at" field.
func (usc *UserSubscriptionCreate) SetExpiresAt(t time.Time) *UserSubscriptionCreate {
	usc.mutation.SetExpiresAt(t)
	return usc
}

// SetUserID sets the "user_id" field.
func (usc *UserSubscriptionCreate) SetUserID(s string) *UserSubscriptionCreate {
	usc.mutation.SetUserID(s)
	return usc
}

// SetPermissions sets the "permissions" field.
func (usc *UserSubscriptionCreate) SetPermissions(s string) *UserSubscriptionCreate {
	usc.mutation.SetPermissions(s)
	return usc
}

// SetReferenceID sets the "reference_id" field.
func (usc *UserSubscriptionCreate) SetReferenceID(s string) *UserSubscriptionCreate {
	usc.mutation.SetReferenceID(s)
	return usc
}

// SetID sets the "id" field.
func (usc *UserSubscriptionCreate) SetID(s string) *UserSubscriptionCreate {
	usc.mutation.SetID(s)
	return usc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (usc *UserSubscriptionCreate) SetNillableID(s *string) *UserSubscriptionCreate {
	if s != nil {
		usc.SetID(*s)
	}
	return usc
}

// SetUser sets the "user" edge to the User entity.
func (usc *UserSubscriptionCreate) SetUser(u *User) *UserSubscriptionCreate {
	return usc.SetUserID(u.ID)
}

// Mutation returns the UserSubscriptionMutation object of the builder.
func (usc *UserSubscriptionCreate) Mutation() *UserSubscriptionMutation {
	return usc.mutation
}

// Save creates the UserSubscription in the database.
func (usc *UserSubscriptionCreate) Save(ctx context.Context) (*UserSubscription, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSubscriptionCreate) SaveX(ctx context.Context) *UserSubscription {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSubscriptionCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSubscriptionCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSubscriptionCreate) defaults() {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		v := usersubscription.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		v := usersubscription.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
	if _, ok := usc.mutation.ID(); !ok {
		v := usersubscription.DefaultID()
		usc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSubscriptionCreate) check() error {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "UserSubscription.created_at"`)}
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "UserSubscription.updated_at"`)}
	}
	if _, ok := usc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`db: missing required field "UserSubscription.type"`)}
	}
	if v, ok := usc.mutation.GetType(); ok {
		if err := usersubscription.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "UserSubscription.type": %w`, err)}
		}
	}
	if _, ok := usc.mutation.ExpiresAt(); !ok {
		return &ValidationError{Name: "expires_at", err: errors.New(`db: missing required field "UserSubscription.expires_at"`)}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`db: missing required field "UserSubscription.user_id"`)}
	}
	if v, ok := usc.mutation.UserID(); ok {
		if err := usersubscription.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`db: validator failed for field "UserSubscription.user_id": %w`, err)}
		}
	}
	if _, ok := usc.mutation.Permissions(); !ok {
		return &ValidationError{Name: "permissions", err: errors.New(`db: missing required field "UserSubscription.permissions"`)}
	}
	if v, ok := usc.mutation.Permissions(); ok {
		if err := usersubscription.PermissionsValidator(v); err != nil {
			return &ValidationError{Name: "permissions", err: fmt.Errorf(`db: validator failed for field "UserSubscription.permissions": %w`, err)}
		}
	}
	if _, ok := usc.mutation.ReferenceID(); !ok {
		return &ValidationError{Name: "reference_id", err: errors.New(`db: missing required field "UserSubscription.reference_id"`)}
	}
	if v, ok := usc.mutation.ReferenceID(); ok {
		if err := usersubscription.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "UserSubscription.reference_id": %w`, err)}
		}
	}
	if _, ok := usc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`db: missing required edge "UserSubscription.user"`)}
	}
	return nil
}

func (usc *UserSubscriptionCreate) sqlSave(ctx context.Context) (*UserSubscription, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserSubscription.ID type: %T", _spec.ID.Value)
		}
	}
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSubscriptionCreate) createSpec() (*UserSubscription, *sqlgraph.CreateSpec) {
	var (
		_node = &UserSubscription{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(usersubscription.Table, sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeString))
	)
	if id, ok := usc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.SetField(usersubscription.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.SetField(usersubscription.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := usc.mutation.GetType(); ok {
		_spec.SetField(usersubscription.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := usc.mutation.ExpiresAt(); ok {
		_spec.SetField(usersubscription.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := usc.mutation.Permissions(); ok {
		_spec.SetField(usersubscription.FieldPermissions, field.TypeString, value)
		_node.Permissions = value
	}
	if value, ok := usc.mutation.ReferenceID(); ok {
		_spec.SetField(usersubscription.FieldReferenceID, field.TypeString, value)
		_node.ReferenceID = value
	}
	if nodes := usc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usersubscription.UserTable,
			Columns: []string{usersubscription.UserColumn},
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

// UserSubscriptionCreateBulk is the builder for creating many UserSubscription entities in bulk.
type UserSubscriptionCreateBulk struct {
	config
	err      error
	builders []*UserSubscriptionCreate
}

// Save creates the UserSubscription entities in the database.
func (uscb *UserSubscriptionCreateBulk) Save(ctx context.Context) ([]*UserSubscription, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserSubscription, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSubscriptionMutation)
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
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSubscriptionCreateBulk) SaveX(ctx context.Context) []*UserSubscription {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSubscriptionCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSubscriptionCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
