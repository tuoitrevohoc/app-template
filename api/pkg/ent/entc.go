//go:build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
    ex, err := entgql.NewExtension(
        // Tell Ent to generate a GraphQL schema for
        // the Ent schema in a file named ent.graphql.
        entgql.WithSchemaGenerator(),
        entgql.WithSchemaPath("./schema/graphql/ent.graphql"),
		entgql.WithConfigPath("./gqlgen.yml"),
    )
    
    if err != nil {
        log.Fatalf("creating entgql extension: %v", err)
    }
    
    opts := []entc.Option{
        entc.Extensions(ex),
    }

    config := gen.Config{
        Target: "./pkg/ent",
        Package: "github.com/tuoitrevohoc/app-template/api/pkg/ent",
    }
    
    if err := entc.Generate("./schema/entity", &config, opts...); err != nil {
        log.Fatalf("running ent codegen: %v", err)
    }
}
