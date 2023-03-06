//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/tuoitrevohoc/app-template/api/internal/config"
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap"
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap/data"
	"github.com/tuoitrevohoc/app-template/api/internal/bootstrap/data/migrations"
	"github.com/tuoitrevohoc/app-template/api/pkg/resolvers"
	"github.com/tuoitrevohoc/app-template/api/pkg/logger"
)

func CreateServer() (*Server, error) {
	wire.Build(
		config.NewConfig,
		migrations.AllMigrations,
		data.NewMigrator,
		bootstrap.NewSchema,
		bootstrap.NewEntClient,
		resolvers.NewResolver,
		logger.NewLogger,
		logger.NewMiddleWare,
		NewServer,
	)

	return nil, nil
}
