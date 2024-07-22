package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

type ModerationRequest struct {
	ent.Schema
}

func (ModerationRequest) Fields() []ent.Field {
	return append(defaultFields,
		field.String("moderator_id").Nillable().Optional(),
		field.String("moderator_comment").Optional(),
		field.String("context").Optional(),
		field.String("reference_id").NotEmpty(),
		field.String("requestor_id").NotEmpty().Immutable(),
		field.String("action_reason").Optional(),
		field.Enum("action_status").GoType(models.ModerationStatus("")),
		field.JSON("data", map[string]any{}),
	)
}

func (ModerationRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("moderator", User.Type).Ref("moderation_actions").Field("moderator_id").Unique(),
		edge.From("requestor", User.Type).Ref("moderation_requests").Field("requestor_id").Required().Immutable().Unique(),
	}
}

func (ModerationRequest) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("reference_id"),
		index.Fields("requestor_id"),
		index.Fields("moderator_id"),
		index.Fields("requestor_id", "reference_id"),
		index.Fields("requestor_id", "reference_id", "action_status"),
	}
}
