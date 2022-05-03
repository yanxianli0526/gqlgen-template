package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"graphql-go-template/internal/gql/generated"
	gqlmodels "graphql-go-template/internal/gql/models"
	"graphql-go-template/internal/models"
)

type Resolver struct{}

func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
	panic("not implemented")
}

func (r *mutationResolver) CreateMenu(ctx context.Context, input *gqlmodels.MenuInput) (bool, error) {
	panic("not implemented")
}

func (r *mutationResolver) UpdateMenu(ctx context.Context, input *gqlmodels.MenuInput, id string) (bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Menus(ctx context.Context) ([]*models.Menu, error) {
	panic("not implemented")
}

func (r *queryResolver) Menu(ctx context.Context, id string) (*models.Menu, error) {
	panic("not implemented")
}

// Menu returns generated.MenuResolver implementation.
func (r *Resolver) Menu() generated.MenuResolver { return &menuResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type menuResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
