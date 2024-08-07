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
		field.Enum("type").
			GoType(models.TaskType("")),
		field.String("reference_id").NotEmpty(),
		field.Strings("targets"),
		//
		field.Enum("status").
			GoType(models.TaskStatus("")),
		field.Time("scheduled_after"),
		field.Time("last_run"),
		field.Int("tries_left").Default(0),
		//
		field.JSON("logs", []models.TaskLog{}),
		field.JSON("data", map[string]string{}),
	)
}

// Edges of the CronTask.
func (CronTask) Edges() []ent.Edge {
	return nil
}

func (CronTask) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id"),
		index.Fields("reference_id"),
		index.Fields("status", "last_run"),
		index.Fields("status", "created_at"),
		index.Fields("status", "scheduled_after"),
	}
}
