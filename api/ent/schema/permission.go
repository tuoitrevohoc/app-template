package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invoice holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Name of the permission"),
		field.String("description"),
	}
}

// Edges of the Invoice.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions"),
	}
}

func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate()),
		entgql.RelayConnection(),
	}
}
