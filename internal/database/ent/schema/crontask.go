package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/cufee/aftermath/internal/database/models"
)

// CronTask holds the schema definition for the CronTask entity.
type CronTask struct {
	ent.Schema
}

// Fields of the CronTask.
func (CronTask) Fields() []ent.Field {
	return append(defaultFields,
		field.String("type").NotEmpty(),
		field.String("reference_id").NotEmpty(),
		field.Strings("targets"),
		//
		field.String("status").NotEmpty(),
		field.Int("scheduled_after"),
		field.Int("last_run"),
		//
		field.JSON("logs", []models.TaskLog{}),
		field.JSON("data", map[string]any{}),
	)
}

// Edges of the CronTask.
func (CronTask) Edges() []ent.Edge {
	return nil
}

func (CronTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("reference_id"),
		index.Fields("status", "last_run"),
		index.Fields("status", "created_at"),
		index.Fields("status", "scheduled_after"),
	}
}
