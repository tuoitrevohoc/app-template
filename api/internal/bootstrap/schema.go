package bootstrap

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/tuoitrevohoc/app-template/api/pkg/graph"
	"github.com/tuoitrevohoc/app-template/api/pkg/resolvers"
)

func NewSchema(resolver *resolvers.Resolver) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(graph.Config{
		Resolvers: resolver,
	})
}
