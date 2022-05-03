package resolvers

import (
	"context"

	"graphql-go-template/internal/models"
)

// Category resolvers
type categoryResolver struct{ *Resolver }

func (r *categoryResolver) ID(ctx context.Context, obj *models.Category) (string, error) {
	return obj.ID.String(), nil
}
