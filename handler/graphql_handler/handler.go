package graphql_handler

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"manga-server/usecase"
)

func NewHandler(usecase usecase.UseCase) (*handler.Handler) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: newQuery(usecase),
			Mutation: newMutation(usecase),
		},
	)
	if err != nil {
		fmt.Println(err);
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}