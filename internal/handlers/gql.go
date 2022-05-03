package handlers

import (
	"graphql-go-template/internal/auth"
	orm "graphql-go-template/internal/database"
	gql "graphql-go-template/internal/gql/generated"
	"graphql-go-template/internal/gql/resolvers"

	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// GraphqlHandler defines the GQLGen GraphQL server handler
func GraphqlHandler(orm *orm.ORM) gin.HandlerFunc {

	h := auth.Middleware(
		handler.GraphQL(gql.NewExecutableSchema(resolvers.NewRootResolvers(orm))),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
