package resolvers

import (
	"context"

	"graphql-go-template/internal/models"
)

func (r *queryResolver) Menus(ctx context.Context) ([]*models.Menu, error) {
	return nil, nil
}

func (r *queryResolver) Menu(ctx context.Context, patientIdStr string) (*models.Menu, error) {
	return nil, nil
}

// Menu resolvers
type menuResolver struct{ *Resolver }

func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
	return obj.ID.String(), nil
}
