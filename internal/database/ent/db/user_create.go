// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cufee/aftermath/internal/database/ent/db/discordinteraction"
	"github.com/cufee/aftermath/internal/database/ent/db/moderationrequest"
	"github.com/cufee/aftermath/internal/database/ent/db/session"
	"github.com/cufee/aftermath/internal/database/ent/db/user"
	"github.com/cufee/aftermath/internal/database/ent/db/userconnection"
	"github.com/cufee/aftermath/internal/database/ent/db/usercontent"
	"github.com/cufee/aftermath/internal/database/ent/db/userrestriction"
	"github.com/cufee/aftermath/internal/database/ent/db/usersubscription"
	"github.com/cufee/aftermath/internal/database/ent/db/widgetsettings"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(t time.Time) *UserCreate {
	uc.mutation.SetCreatedAt(t)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetCreatedAt(*t)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(t time.Time) *UserCreate {
	uc.mutation.SetUpdatedAt(t)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(t *time.Time) *UserCreate {
	if t != nil {
		uc.SetUpdatedAt(*t)
	}
	return uc
}

// SetUsername sets the "username" field.
func (uc *UserCreate) SetUsername(s string) *UserCreate {
	uc.mutation.SetUsername(s)
	return uc
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uc *UserCreate) SetNillableUsername(s *string) *UserCreate {
	if s != nil {
		uc.SetUsername(*s)
	}
	return uc
}

// SetPermissions sets the "permissions" field.
func (uc *UserCreate) SetPermissions(s string) *UserCreate {
	uc.mutation.SetPermissions(s)
	return uc
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (uc *UserCreate) SetNillablePermissions(s *string) *UserCreate {
	if s != nil {
		uc.SetPermissions(*s)
	}
	return uc
}

// SetFeatureFlags sets the "feature_flags" field.
func (uc *UserCreate) SetFeatureFlags(s []string) *UserCreate {
	uc.mutation.SetFeatureFlags(s)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(s string) *UserCreate {
	uc.mutation.SetID(s)
	return uc
}

// AddDiscordInteractionIDs adds the "discord_interactions" edge to the DiscordInteraction entity by IDs.
func (uc *UserCreate) AddDiscordInteractionIDs(ids ...string) *UserCreate {
	uc.mutation.AddDiscordInteractionIDs(ids...)
	return uc
}

// AddDiscordInteractions adds the "discord_interactions" edges to the DiscordInteraction entity.
func (uc *UserCreate) AddDiscordInteractions(d ...*DiscordInteraction) *UserCreate {
	ids := make([]string, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uc.AddDiscordInteractionIDs(ids...)
}

// AddSubscriptionIDs adds the "subscriptions" edge to the UserSubscription entity by IDs.
func (uc *UserCreate) AddSubscriptionIDs(ids ...string) *UserCreate {
	uc.mutation.AddSubscriptionIDs(ids...)
	return uc
}

// AddSubscriptions adds the "subscriptions" edges to the UserSubscription entity.
func (uc *UserCreate) AddSubscriptions(u ...*UserSubscription) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddSubscriptionIDs(ids...)
}

// AddConnectionIDs adds the "connections" edge to the UserConnection entity by IDs.
func (uc *UserCreate) AddConnectionIDs(ids ...string) *UserCreate {
	uc.mutation.AddConnectionIDs(ids...)
	return uc
}

// AddConnections adds the "connections" edges to the UserConnection entity.
func (uc *UserCreate) AddConnections(u ...*UserConnection) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddConnectionIDs(ids...)
}

// AddWidgetIDs adds the "widgets" edge to the WidgetSettings entity by IDs.
func (uc *UserCreate) AddWidgetIDs(ids ...string) *UserCreate {
	uc.mutation.AddWidgetIDs(ids...)
	return uc
}

// AddWidgets adds the "widgets" edges to the WidgetSettings entity.
func (uc *UserCreate) AddWidgets(w ...*WidgetSettings) *UserCreate {
	ids := make([]string, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uc.AddWidgetIDs(ids...)
}

// AddContentIDs adds the "content" edge to the UserContent entity by IDs.
func (uc *UserCreate) AddContentIDs(ids ...string) *UserCreate {
	uc.mutation.AddContentIDs(ids...)
	return uc
}

// AddContent adds the "content" edges to the UserContent entity.
func (uc *UserCreate) AddContent(u ...*UserContent) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddContentIDs(ids...)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (uc *UserCreate) AddSessionIDs(ids ...string) *UserCreate {
	uc.mutation.AddSessionIDs(ids...)
	return uc
}

// AddSessions adds the "sessions" edges to the Session entity.
func (uc *UserCreate) AddSessions(s ...*Session) *UserCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddSessionIDs(ids...)
}

// AddModerationRequestIDs adds the "moderation_requests" edge to the ModerationRequest entity by IDs.
func (uc *UserCreate) AddModerationRequestIDs(ids ...string) *UserCreate {
	uc.mutation.AddModerationRequestIDs(ids...)
	return uc
}

// AddModerationRequests adds the "moderation_requests" edges to the ModerationRequest entity.
func (uc *UserCreate) AddModerationRequests(m ...*ModerationRequest) *UserCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddModerationRequestIDs(ids...)
}

// AddModerationActionIDs adds the "moderation_actions" edge to the ModerationRequest entity by IDs.
func (uc *UserCreate) AddModerationActionIDs(ids ...string) *UserCreate {
	uc.mutation.AddModerationActionIDs(ids...)
	return uc
}

// AddModerationActions adds the "moderation_actions" edges to the ModerationRequest entity.
func (uc *UserCreate) AddModerationActions(m ...*ModerationRequest) *UserCreate {
	ids := make([]string, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uc.AddModerationActionIDs(ids...)
}

// AddRestrictionIDs adds the "restrictions" edge to the UserRestriction entity by IDs.
func (uc *UserCreate) AddRestrictionIDs(ids ...string) *UserCreate {
	uc.mutation.AddRestrictionIDs(ids...)
	return uc
}

// AddRestrictions adds the "restrictions" edges to the UserRestriction entity.
func (uc *UserCreate) AddRestrictions(u ...*UserRestriction) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddRestrictionIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	uc.defaults()
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt()
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt()
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.Username(); !ok {
		v := user.DefaultUsername
		uc.mutation.SetUsername(v)
	}
	if _, ok := uc.mutation.Permissions(); !ok {
		v := user.DefaultPermissions
		uc.mutation.SetPermissions(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`db: missing required field "User.username"`)}
	}
	if _, ok := uc.mutation.Permissions(); !ok {
		return &ValidationError{Name: "permissions", err: errors.New(`db: missing required field "User.permissions"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected User.ID type: %T", _spec.ID.Value)
		}
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeString))
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := uc.mutation.Permissions(); ok {
		_spec.SetField(user.FieldPermissions, field.TypeString, value)
		_node.Permissions = value
	}
	if value, ok := uc.mutation.FeatureFlags(); ok {
		_spec.SetField(user.FieldFeatureFlags, field.TypeJSON, value)
		_node.FeatureFlags = value
	}
	if nodes := uc.mutation.DiscordInteractionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.DiscordInteractionsTable,
			Columns: []string{user.DiscordInteractionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(discordinteraction.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.SubscriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SubscriptionsTable,
			Columns: []string{user.SubscriptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersubscription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ConnectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ConnectionsTable,
			Columns: []string{user.ConnectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userconnection.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.WidgetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.WidgetsTable,
			Columns: []string{user.WidgetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(widgetsettings.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ContentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ContentTable,
			Columns: []string{user.ContentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usercontent.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SessionsTable,
			Columns: []string{user.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(session.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ModerationRequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModerationRequestsTable,
			Columns: []string{user.ModerationRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moderationrequest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.ModerationActionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ModerationActionsTable,
			Columns: []string{user.ModerationActionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(moderationrequest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.mutation.RestrictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.RestrictionsTable,
			Columns: []string{user.RestrictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userrestriction.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
