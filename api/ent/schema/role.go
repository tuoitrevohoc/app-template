package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invoice holds the schema definition for the Invoice entity.
type Role struct {
	ent.Schema
}

// Fields of the Invoice.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Name of the role"),
		field.String("description"),
	}
}

// Edges of the Invoice.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("role"),
		edge.To("permissions", Permission.Type).
			Comment("List of permissions of this role"),
	}
}

func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate()),
		entgql.RelayConnection(),
	}
}
