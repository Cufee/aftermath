// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/moderationrequest"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/models"
)

// ModerationRequestCreate is the builder for creating a ModerationRequest entity.
type ModerationRequestCreate struct {
	config
	mutation *ModerationRequestMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mrc *ModerationRequestCreate) SetCreatedAt(t time.Time) *ModerationRequestCreate {
	mrc.mutation.SetCreatedAt(t)
	return mrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableCreatedAt(t *time.Time) *ModerationRequestCreate {
	if t != nil {
		mrc.SetCreatedAt(*t)
	}
	return mrc
}

// SetUpdatedAt sets the "updated_at" field.
func (mrc *ModerationRequestCreate) SetUpdatedAt(t time.Time) *ModerationRequestCreate {
	mrc.mutation.SetUpdatedAt(t)
	return mrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableUpdatedAt(t *time.Time) *ModerationRequestCreate {
	if t != nil {
		mrc.SetUpdatedAt(*t)
	}
	return mrc
}

// SetModeratorID sets the "moderator_id" field.
func (mrc *ModerationRequestCreate) SetModeratorID(s string) *ModerationRequestCreate {
	mrc.mutation.SetModeratorID(s)
	return mrc
}

// SetNillableModeratorID sets the "moderator_id" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableModeratorID(s *string) *ModerationRequestCreate {
	if s != nil {
		mrc.SetModeratorID(*s)
	}
	return mrc
}

// SetModeratorComment sets the "moderator_comment" field.
func (mrc *ModerationRequestCreate) SetModeratorComment(s string) *ModerationRequestCreate {
	mrc.mutation.SetModeratorComment(s)
	return mrc
}

// SetNillableModeratorComment sets the "moderator_comment" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableModeratorComment(s *string) *ModerationRequestCreate {
	if s != nil {
		mrc.SetModeratorComment(*s)
	}
	return mrc
}

// SetContext sets the "context" field.
func (mrc *ModerationRequestCreate) SetContext(s string) *ModerationRequestCreate {
	mrc.mutation.SetContext(s)
	return mrc
}

// SetNillableContext sets the "context" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableContext(s *string) *ModerationRequestCreate {
	if s != nil {
		mrc.SetContext(*s)
	}
	return mrc
}

// SetReferenceID sets the "reference_id" field.
func (mrc *ModerationRequestCreate) SetReferenceID(s string) *ModerationRequestCreate {
	mrc.mutation.SetReferenceID(s)
	return mrc
}

// SetRequestorID sets the "requestor_id" field.
func (mrc *ModerationRequestCreate) SetRequestorID(s string) *ModerationRequestCreate {
	mrc.mutation.SetRequestorID(s)
	return mrc
}

// SetActionReason sets the "action_reason" field.
func (mrc *ModerationRequestCreate) SetActionReason(s string) *ModerationRequestCreate {
	mrc.mutation.SetActionReason(s)
	return mrc
}

// SetNillableActionReason sets the "action_reason" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableActionReason(s *string) *ModerationRequestCreate {
	if s != nil {
		mrc.SetActionReason(*s)
	}
	return mrc
}

// SetActionStatus sets the "action_status" field.
func (mrc *ModerationRequestCreate) SetActionStatus(ms models.ModerationStatus) *ModerationRequestCreate {
	mrc.mutation.SetActionStatus(ms)
	return mrc
}

// SetData sets the "data" field.
func (mrc *ModerationRequestCreate) SetData(m map[string]interface{}) *ModerationRequestCreate {
	mrc.mutation.SetData(m)
	return mrc
}

// SetID sets the "id" field.
func (mrc *ModerationRequestCreate) SetID(s string) *ModerationRequestCreate {
	mrc.mutation.SetID(s)
	return mrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mrc *ModerationRequestCreate) SetNillableID(s *string) *ModerationRequestCreate {
	if s != nil {
		mrc.SetID(*s)
	}
	return mrc
}

// SetModerator sets the "moderator" edge to the User entity.
func (mrc *ModerationRequestCreate) SetModerator(u *User) *ModerationRequestCreate {
	return mrc.SetModeratorID(u.ID)
}

// SetRequestor sets the "requestor" edge to the User entity.
func (mrc *ModerationRequestCreate) SetRequestor(u *User) *ModerationRequestCreate {
	return mrc.SetRequestorID(u.ID)
}

// Mutation returns the ModerationRequestMutation object of the builder.
func (mrc *ModerationRequestCreate) Mutation() *ModerationRequestMutation {
	return mrc.mutation
}

// Save creates the ModerationRequest in the database.
func (mrc *ModerationRequestCreate) Save(ctx context.Context) (*ModerationRequest, error) {
	mrc.defaults()
	return withHooks(ctx, mrc.sqlSave, mrc.mutation, mrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mrc *ModerationRequestCreate) SaveX(ctx context.Context) *ModerationRequest {
	v, err := mrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrc *ModerationRequestCreate) Exec(ctx context.Context) error {
	_, err := mrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrc *ModerationRequestCreate) ExecX(ctx context.Context) {
	if err := mrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mrc *ModerationRequestCreate) defaults() {
	if _, ok := mrc.mutation.CreatedAt(); !ok {
		v := moderationrequest.DefaultCreatedAt()
		mrc.mutation.SetCreatedAt(v)
	}
	if _, ok := mrc.mutation.UpdatedAt(); !ok {
		v := moderationrequest.DefaultUpdatedAt()
		mrc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mrc.mutation.ID(); !ok {
		v := moderationrequest.DefaultID()
		mrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mrc *ModerationRequestCreate) check() error {
	if _, ok := mrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "ModerationRequest.created_at"`)}
	}
	if _, ok := mrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "ModerationRequest.updated_at"`)}
	}
	if _, ok := mrc.mutation.ReferenceID(); !ok {
		return &ValidationError{Name: "reference_id", err: errors.New(`db: missing required field "ModerationRequest.reference_id"`)}
	}
	if v, ok := mrc.mutation.ReferenceID(); ok {
		if err := moderationrequest.ReferenceIDValidator(v); err != nil {
			return &ValidationError{Name: "reference_id", err: fmt.Errorf(`db: validator failed for field "ModerationRequest.reference_id": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.RequestorID(); !ok {
		return &ValidationError{Name: "requestor_id", err: errors.New(`db: missing required field "ModerationRequest.requestor_id"`)}
	}
	if v, ok := mrc.mutation.RequestorID(); ok {
		if err := moderationrequest.RequestorIDValidator(v); err != nil {
			return &ValidationError{Name: "requestor_id", err: fmt.Errorf(`db: validator failed for field "ModerationRequest.requestor_id": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.ActionStatus(); !ok {
		return &ValidationError{Name: "action_status", err: errors.New(`db: missing required field "ModerationRequest.action_status"`)}
	}
	if v, ok := mrc.mutation.ActionStatus(); ok {
		if err := moderationrequest.ActionStatusValidator(v); err != nil {
			return &ValidationError{Name: "action_status", err: fmt.Errorf(`db: validator failed for field "ModerationRequest.action_status": %w`, err)}
		}
	}
	if _, ok := mrc.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`db: missing required field "ModerationRequest.data"`)}
	}
	if _, ok := mrc.mutation.RequestorID(); !ok {
		return &ValidationError{Name: "requestor", err: errors.New(`db: missing required edge "ModerationRequest.requestor"`)}
	}
	return nil
}

func (mrc *ModerationRequestCreate) sqlSave(ctx context.Context) (*ModerationRequest, error) {
	if err := mrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ModerationRequest.ID type: %T", _spec.ID.Value)
		}
	}
	mrc.mutation.id = &_node.ID
	mrc.mutation.done = true
	return _node, nil
}

func (mrc *ModerationRequestCreate) createSpec() (*ModerationRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &ModerationRequest{config: mrc.config}
		_spec = sqlgraph.NewCreateSpec(moderationrequest.Table, sqlgraph.NewFieldSpec(moderationrequest.FieldID, field.TypeString))
	)
	if id, ok := mrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mrc.mutation.CreatedAt(); ok {
		_spec.SetField(moderationrequest.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mrc.mutation.UpdatedAt(); ok {
		_spec.SetField(moderationrequest.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mrc.mutation.ModeratorComment(); ok {
		_spec.SetField(moderationrequest.FieldModeratorComment, field.TypeString, value)
		_node.ModeratorComment = value
	}
	if value, ok := mrc.mutation.Context(); ok {
		_spec.SetField(moderationrequest.FieldContext, field.TypeString, value)
		_node.Context = value
	}
	if value, ok := mrc.mutation.ReferenceID(); ok {
		_spec.SetField(moderationrequest.FieldReferenceID, field.TypeString, value)
		_node.ReferenceID = value
	}
	if value, ok := mrc.mutation.ActionReason(); ok {
		_spec.SetField(moderationrequest.FieldActionReason, field.TypeString, value)
		_node.ActionReason = value
	}
	if value, ok := mrc.mutation.ActionStatus(); ok {
		_spec.SetField(moderationrequest.FieldActionStatus, field.TypeEnum, value)
		_node.ActionStatus = value
	}
	if value, ok := mrc.mutation.Data(); ok {
		_spec.SetField(moderationrequest.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if nodes := mrc.mutation.ModeratorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moderationrequest.ModeratorTable,
			Columns: []string{moderationrequest.ModeratorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ModeratorID = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mrc.mutation.RequestorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   moderationrequest.RequestorTable,
			Columns: []string{moderationrequest.RequestorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RequestorID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ModerationRequestCreateBulk is the builder for creating many ModerationRequest entities in bulk.
type ModerationRequestCreateBulk struct {
	config
	err      error
	builders []*ModerationRequestCreate
}

// Save creates the ModerationRequest entities in the database.
func (mrcb *ModerationRequestCreateBulk) Save(ctx context.Context) ([]*ModerationRequest, error) {
	if mrcb.err != nil {
		return nil, mrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mrcb.builders))
	nodes := make([]*ModerationRequest, len(mrcb.builders))
	mutators := make([]Mutator, len(mrcb.builders))
	for i := range mrcb.builders {
		func(i int, root context.Context) {
			builder := mrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModerationRequestMutation)
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
					_, err = mutators[i+1].Mutate(root, mrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mrcb *ModerationRequestCreateBulk) SaveX(ctx context.Context) []*ModerationRequest {
	v, err := mrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mrcb *ModerationRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := mrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mrcb *ModerationRequestCreateBulk) ExecX(ctx context.Context) {
	if err := mrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
