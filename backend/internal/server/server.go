package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/tuoitrevohoc/app-template/backend/internal/config"
	"github.com/tuoitrevohoc/app-template/backend/pkg/logger"
)

type Server struct {
	schema graphql.ExecutableSchema
	config config.Configurations
	loggerMiddleWare *logger.MiddleWare
}

func (s *Server) Start() error {
	srv := handler.NewDefaultServer(s.schema)
	router := chi.NewRouter()

	router.Use(s.loggerMiddleWare.GetMiddleWare())

	router.Handle("/graphql", playground.Handler("GraphQL Playground", "/graphql/query"))
	router.Handle("/graphql/query", srv)

	log.Printf("Listening on http://localhost:%v/graphql", s.config.Port)

	return http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), router)
}

func NewServer(
	schema graphql.ExecutableSchema,
	config config.Configurations,
	loggerMiddleWare *logger.MiddleWare,
) *Server {
	return &Server{
		schema: schema,
		config: config,
		loggerMiddleWare: loggerMiddleWare,
	}
}