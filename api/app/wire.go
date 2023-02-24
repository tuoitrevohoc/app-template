//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/tuoitrevohoc/app-template/api/app/config"
	"github.com/tuoitrevohoc/app-template/api/app/bootstrap"
	"github.com/tuoitrevohoc/app-template/api/app/resolvers"
	"github.com/tuoitrevohoc/app-template/api/pkg/logger"
)

func CreateServer() (*Server, error) {
	wire.Build(
		config.NewConfig,
		bootstrap.NewSchema,
		bootstrap.NewEntClient,
		resolvers.NewResolver,
		logger.NewLogger,
		logger.NewMiddleWare,
		NewServer,
	)

	return nil, nil
}
