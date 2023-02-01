package bootstrap

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/tuoitrevohoc/app-template/api/app/resolvers"
	"github.com/tuoitrevohoc/app-template/api/ent"
)

func NewSchema(resolver *resolvers.Resolver) graphql.ExecutableSchema {
	return ent.NewExecutableSchema(ent.Config{
		Resolvers: resolver,
	})
}
