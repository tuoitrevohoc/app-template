package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tuoitrevohoc/app-template/api/ent"
)

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return r.Client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return r.Client.Noders(ctx, ids)
}

func (r *queryResolver) Invoices(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int) (*ent.InvoiceConnection, error) {
	return r.Client.Invoice.Query().Paginate(ctx, after, first, before, last)
}

// Query returns ent.QueryResolver implementation.
func (r *Resolver) Query() ent.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
