package graphql_handler

import (
	"manga-server/repository/mysql_repository/schema"
	"manga-server/usecase"

	"github.com/graphql-go/graphql"
)

func newQuery(usecase usecase.UseCase) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getManga": schema.GetManga(usecase),
		},
	})
}

func newMutation(usecase usecase.UseCase) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"addManga": schema.AddManga(usecase),
		},
	})
}