package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

var defaultFields = []ent.Field{
	field.String("id").
		Unique().
		Immutable().
		DefaultFunc(cuid.New),
	field.Int("created_at").
		Immutable().
		DefaultFunc(timeNow),
	field.Int("updated_at").
		DefaultFunc(timeNow).
		UpdateDefault(timeNow),
}

func timeNow() int { return int(time.Now().Unix()) }
