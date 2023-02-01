package resolvers

import "github.com/tuoitrevohoc/app-template/api/ent"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Client *ent.Client
}

func NewResolver(client *ent.Client) *Resolver {
	return &Resolver{
		Client: client,
	}
}
