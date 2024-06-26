package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

type DiscordInteraction struct {
	ent.Schema
}

func (DiscordInteraction) Fields() []ent.Field {
	return append(defaultFields,
		field.String("command").NotEmpty(),
		field.String("user_id").NotEmpty().Immutable(),
		field.String("reference_id").NotEmpty(),
		field.Enum("type").
			GoType(models.DiscordInteractionType("")),
		field.String("locale"),
		field.JSON("options", models.DiscordInteractionOptions{}),
	)
}

func (DiscordInteraction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("discord_interactions").Field("user_id").Required().Immutable().Unique(),
	}
}

func (DiscordInteraction) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("command"),
		index.Fields("user_id"),
		index.Fields("user_id", "type"),
		index.Fields("reference_id"),
	}
}
