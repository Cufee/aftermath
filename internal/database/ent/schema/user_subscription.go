package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserSubscription holds the schema definition for the UserSubscription entity.
type UserSubscription struct {
	ent.Schema
}

// Fields of the UserSubscription.
func (UserSubscription) Fields() []ent.Field {
	return append(defaultFields,
		field.Enum("type").
			GoType(models.SubscriptionType("")),
		field.Int64("expires_at"),
		//
		field.String("user_id").NotEmpty().Immutable(),
		field.String("permissions").NotEmpty(),
		field.String("reference_id").NotEmpty(),
	)
}

// Edges of the UserSubscription.
func (UserSubscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("subscriptions").Field("user_id").Required().Immutable().Unique(),
	}
}

func (UserSubscription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("type", "user_id"),
		index.Fields("expires_at"),
		index.Fields("expires_at", "user_id"),
	}
}
