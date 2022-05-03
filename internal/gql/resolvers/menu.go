package resolvers

import (
	"context"
	orm "graphql-go-template/internal/database"

	gqlmodels "graphql-go-template/internal/gql/models"
	"graphql-go-template/internal/models"

	"github.com/google/uuid"
)

func (r *mutationResolver) CreateMenu(ctx context.Context, input *gqlmodels.MenuInput) (bool, error) {

	menu := models.Menu{
		ID:            uuid.New(),
		ItemName:      input.ItemName,
		Price:         input.Price,
		IsStopSelling: input.IsStopSelling,
	}

	err := orm.CreateMenu(r.ORM.DB, &menu)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateMenu(ctx context.Context, input *gqlmodels.MenuInput, id string) (bool, error) {
	return true, nil
}

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
