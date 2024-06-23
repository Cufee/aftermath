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
	field.Int64("created_at").
		Immutable().
		DefaultFunc(timeNow),
	field.Int64("updated_at").
		DefaultFunc(timeNow).
		UpdateDefault(timeNow),
}

func timeNow() int64 { return time.Now().Unix() }
