package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Invoice holds the schema definition for the Invoice entity.
type User struct {
	ent.Schema
}

// Fields of the Invoice.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Comment("The username"),
		field.String("password").Comment("Hashed password").Sensitive(),
	}
}

// Edges of the Invoice.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).Unique(),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
        entgql.QueryField(),
        entgql.Mutations(entgql.MutationCreate()),
		entgql.RelayConnection(),
	}
}
