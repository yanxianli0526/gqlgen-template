package resolvers

import (
	"context"
	orm "graphql-go-template/internal/database"

	"graphql-go-template/internal/models"
)

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	users, err := orm.GetUsers(r.ORM.DB)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// Menu resolvers
type userResolver struct{ *Resolver }

func (r *userResolver) ID(ctx context.Context, obj *models.User) (string, error) {
	return obj.ID.String(), nil
}
