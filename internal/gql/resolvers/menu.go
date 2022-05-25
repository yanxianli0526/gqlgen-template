package resolvers

import (
	"context"
	orm "graphql-go-template/internal/database"
	"regexp"
	"strings"

	gqlmodels "graphql-go-template/internal/gql/models"
	"graphql-go-template/internal/models"

	"github.com/99designs/gqlgen/graphql"
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

func (r *mutationResolver) UpdateMenu(ctx context.Context, input *gqlmodels.MenuInput, menuIdStr string) (bool, error) {

	id, err := uuid.Parse(menuIdStr)
	if err != nil {
		return false, err
	}

	menu := models.Menu{
		ID:            id,
		ItemName:      input.ItemName,
		Price:         input.Price,
		IsStopSelling: input.IsStopSelling,
	}

	err = orm.UpdateMenu(r.ORM.DB, &menu)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) DeleteMenu(ctx context.Context, menuIdStr string) (bool, error) {

	id, err := uuid.Parse(menuIdStr)
	if err != nil {
		return false, err
	}

	err = orm.DeleteMenu(r.ORM.DB, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *queryResolver) Menus(ctx context.Context) ([]*models.Menu, error) {
	collectFieldsCtx := graphql.CollectFieldsCtx(ctx, nil)

	// 可以用這個減低db的負擔 只針對需要的欄位去進行select(其實可做不可不做)
	var allName []string
	for i := range collectFieldsCtx {
		dbColumnName := ToSnakeCase(collectFieldsCtx[i].Name)
		allName = append(allName, dbColumnName)
	}

	menus, err := orm.GetMenus(r.ORM.DB, allName)
	if err != nil {
		return nil, err
	}

	for i := range menus {
		if collectFieldsCtx[1].Arguments[0].Value.String() == "JPY" {
			menus[i].Price = menus[i].Price * 4.32
		}
	}

	return menus, nil
}

func (r *queryResolver) Menu(ctx context.Context, menuIdStr string) (*models.Menu, error) {
	collectFieldsCtx := graphql.CollectFieldsCtx(ctx, nil)

	menuId, err := uuid.Parse(menuIdStr)
	if err != nil {
		return nil, err
	}

	menu, err := orm.GetMenuById(r.ORM.DB, menuId)
	if err != nil {
		return nil, err
	}
	if collectFieldsCtx[1].Arguments[0].Value.String() == "JPY" {
		menu.Price = menu.Price * 4.32
	}
	return menu, nil
}

// Menu resolvers
type menuResolver struct{ *Resolver }

func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
	return obj.ID.String(), nil
}

var matchFirstCap = regexp.MustCompile("([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "_${1}")
	return strings.ToLower(snake)
}
