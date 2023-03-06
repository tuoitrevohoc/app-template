// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap"
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap/data"
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap/data/migrations"
	"github.com/tuoitrevohoc/app-template/api/internal/config"
	"github.com/tuoitrevohoc/app-template/api/pkg/logger"
	"github.com/tuoitrevohoc/app-template/api/pkg/resolvers"
)

// Injectors from wire.go:

func CreateServer() (*Server, error) {
	configurations, err := config.NewConfig()
	if err != nil {
		return nil, err
	}
	zapLogger, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	v := migrations.AllMigrations()
	migrator := data.NewMigrator(zapLogger, v)
	client, err := bootstrap.NewEntClient(configurations, migrator)
	if err != nil {
		return nil, err
	}
	resolver := resolvers.NewResolver(client)
	executableSchema := bootstrap.NewSchema(resolver)
	middleWare := logger.NewMiddleWare(zapLogger)
	server := NewServer(executableSchema, configurations, middleWare)
	return server, nil
}
