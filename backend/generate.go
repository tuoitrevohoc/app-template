package main

//go:generate go run -mod=mod ./pkg/ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen
//go:generate go run -mod=mod github.com/google/wire/cmd/wire ./app
