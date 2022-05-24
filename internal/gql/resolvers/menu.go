package resolvers

import (
	"context"
	"fmt"
	orm "graphql-go-template/internal/database"

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

	// fmt.Println(r)
	// fmt.Println(&r)

	// fmt.Println("a", ctx)

	menus, err := orm.GetMenus(r.ORM.DB)
	if err != nil {
		return nil, err
	}

	return menus, nil
}

func (r *queryResolver) Menu(ctx context.Context, menuIdStr string) (*models.Menu, error) {
	// rezCtx := graphql.GetResolverContext(ctx)
	// satisfies := []string{"unit"}
	// qaq1 := graphql.CollectFieldsCtx(ctx, satisfies)
	qaq := graphql.CollectFieldsCtx(ctx, nil)
	// aa := graphql.GetRequestContext(ctx).RawQuery
	// bb := graphql.GetOperationContext(ctx).Variables
	// cc := graphql.GetFieldContext(ctx)
	// graphql.GetRequestContext(ctx)
	// fmt.Println("aa", aa)
	// fmt.Println("bb", bb)
	// fmt.Println("cc", cc)
	// fmt.Println("cc", qaq[1].Arguments[0])
	// fmt.Println("rezCtx", rezCtx)
	// fmt.Println("qaq", qaq)
	// fmt.Println("qaq1", qaq1)

	menuId, err := uuid.Parse(menuIdStr)
	if err != nil {
		return nil, err
	}

	menu, err := orm.GetMenuById(r.ORM.DB, menuId)
	if err != nil {
		return nil, err
	}
	fmt.Println("qaq[1].Arguments[0].Value.String()", qaq[1].Arguments[0].Value.String())

	menu.ItemName = qaq[1].Arguments[0].Value.String()
	if qaq[1].Arguments[0].Value.String() == "JPY" {

		menu.Price = menu.Price * 4.32
	}
	return menu, nil
}

// Menu resolvers
type menuResolver struct{ *Resolver }

func (r *menuResolver) ID(ctx context.Context, obj *models.Menu) (string, error) {
	return obj.ID.String(), nil
}
