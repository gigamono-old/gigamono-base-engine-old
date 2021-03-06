package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/gigamono/gigamono-base-engine/internal/graphql/generated"
	"github.com/gigamono/gigamono-base-engine/internal/graphql/model"
)

func (r *queryResolver) GetSessionUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
