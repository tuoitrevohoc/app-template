package entity

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("leet_code_link"),
		field.String("invoiced_to"),
	}
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return nil
}

func (Invoice) Annotations() []schema.Annotation {
	return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate()),
		entgql.RelayConnection(),
	}
}
