package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/tuoitrevohoc/app-template/api"
	"github.com/tuoitrevohoc/app-template/api/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Invoices(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.InvoiceConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns api.QueryResolver implementation.
func (r *Resolver) Query() api.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
