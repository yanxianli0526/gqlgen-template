package resolvers

import (
	"context"

	orm "graphql-go-template/internal/database"
	"graphql-go-template/internal/errors"
	"graphql-go-template/internal/gql/generated"

	"github.com/99designs/gqlgen/graphql"
)

type contextKey string

var (
	UserIDCtxKey = contextKey("userID")
)

type Resolver struct {
	ORM *orm.ORM
}

func NewRootResolvers(orm *orm.ORM) generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{
			ORM: orm, // pass in the ORM instance in the resolvers to be used
		},
	}

	// Schema Directive
	c.Directives.IsAuthenticated = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ctxUserID := ctx.Value(UserIDCtxKey)
		if ctxUserID == nil {
			return nil, errors.UnAuthorizedError
		}
		return next(ctx)
	}

	return c
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Menu() generated.MenuResolver {
	return &menuResolver{r}
}
func (r *Resolver) User() generated.UserResolver {
	return &userResolver{r}
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }
