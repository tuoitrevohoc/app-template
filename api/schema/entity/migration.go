package entity

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Migration struct {
	ent.Schema
}

// Fields of the Migration.
func (Migration) Fields() []ent.Field {
	return []ent.Field{
		field.String("migration").Unique(),
		field.Time("execution_at").Default(time.Now),
	}
}
