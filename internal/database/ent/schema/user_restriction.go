package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

type UserRestriction struct {
	ent.Schema
}

func (UserRestriction) Fields() []ent.Field {
	return append(defaultFields,
		field.Time("expires_at"),
		field.Enum("type").
			GoType(models.UserRestrictionType("")),
		field.String("user_id").NotEmpty().Immutable(),
		field.String("restriction").NotEmpty(),
		field.String("public_reason").NotEmpty(),
		field.String("moderator_comment"),
		field.JSON("events", []models.RestrictionUpdate{}),
	)
} // Edges of the UserSubscription.
func (UserRestriction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("restrictions").Field("user_id").Required().Immutable().Unique(),
	}
}

func (UserRestriction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("user_id"),
		index.Fields("expires_at", "user_id"),
	}
}
