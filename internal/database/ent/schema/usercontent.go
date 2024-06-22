package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

// UserContent holds the schema definition for the UserContent entity.
type UserContent struct {
	ent.Schema
}

// Fields of the UserContent.
func (UserContent) Fields() []ent.Field {
	return append(defaultFields,
		field.Enum("type").
			GoType(models.UserContentType("")),
		field.String("user_id").Immutable(),
		field.String("reference_id"),
		//
		field.Any("value"),
		field.JSON("metadata", map[string]any{}),
	)
}

// Edges of the UserContent.
func (UserContent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("content").Field("user_id").Required().Immutable().Unique(),
	}
}

func (UserContent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
		index.Fields("type", "user_id"),
		index.Fields("reference_id"),
		index.Fields("type", "reference_id"),
	}
}
