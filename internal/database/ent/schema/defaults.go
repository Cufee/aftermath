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
	field.Time("created_at").
		Immutable().
		Default(timeNow),
	field.Time("updated_at").
		Default(timeNow).
		UpdateDefault(timeNow),
}

func timeNow() time.Time { return time.Now() }
