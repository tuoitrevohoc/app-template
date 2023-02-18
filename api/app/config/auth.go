package config

import "github.com/go-pkgz/auth"

func NewAuthService() *auth.Service {

	options := auth.Opts {
	}

	service := auth.NewService(options)
	return service
}