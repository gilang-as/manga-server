package schema

import (
	"errors"
	"github.com/graphql-go/graphql"
	"manga-server/domain/models"
	graphql2 "manga-server/pkg/graphql"
	"manga-server/usecase"
	"time"
)

// Manga Type
var mangaType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Manga",
		Description: "Manga data",
		Fields: graphql.Fields{
			"id":      &graphql.Field{Type: graphql.ID},
			"title":      &graphql.Field{Type: graphql.String},
			"original_title":       &graphql.Field{Type: graphql.String},
			"english_title": &graphql.Field{Type: graphql.String},
			"status": &graphql.Field{Type: graphql.String},
			"volumes": &graphql.Field{Type: graphql.Int},
			"chapters": &graphql.Field{Type: graphql.Int},
			"publishing": &graphql.Field{Type: graphql.Boolean},
			"published_from": &graphql.Field{Type: graphql.DateTime},
			"published_to": &graphql.Field{Type: graphql.DateTime},
			"synopsis": &graphql.Field{Type: graphql.String},
			"image_url": &graphql.Field{Type: graphql.String},
			"created_at": &graphql.Field{Type: graphql.DateTime},
		},
	},
)

// Manga Query
func GetManga(usecase usecase.UseCase) *graphql.Field {
	return graphql2.Paginated(&graphql2.PaginatedField{
		Name: "Languages",
		Type: mangaType,
		Args: nil,
		DataAndCountResolve: func(p graphql.ResolveParams, page graphql2.Page) (interface{}, *int, error) {
			data, total, err := usecase.GetManga(page.Skip, page.Limit)
			if err != nil {
				return nil, nil, err
			}
			return data, total, nil
		},
	})
}


func AddManga(usecase usecase.UseCase) *graphql.Field {
	return &graphql.Field{
		Type: mangaType,
		Args: graphql.FieldConfigArgument{
			"title":      &graphql.ArgumentConfig{Type: graphql.String},
			"original_title":       &graphql.ArgumentConfig{Type: graphql.String},
			"english_title": &graphql.ArgumentConfig{Type: graphql.String},
			"status": &graphql.ArgumentConfig{Type: graphql.String},
			"volumes": &graphql.ArgumentConfig{Type: graphql.Int},
			"chapters": &graphql.ArgumentConfig{Type: graphql.Int},
			"publishing": &graphql.ArgumentConfig{Type: graphql.Boolean},
			"published_from": &graphql.ArgumentConfig{Type: graphql.DateTime},
			"published_to": &graphql.ArgumentConfig{Type: graphql.DateTime},
			"synopsis": &graphql.ArgumentConfig{Type: graphql.String},
			"image_url": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			title, ok := p.Args["title"].(string)
			originalTitle, ok := p.Args["original_title"].(string)
			englishTitle, ok := p.Args["english_title"].(string)
			status, ok := p.Args["status"].(string)
			volumes, ok := p.Args["volumes"].(int)
			chapters, ok := p.Args["chapters"].(int)
			publishing, ok := p.Args["publishing"].(bool)
			publishedFrom, ok := p.Args["published_from"].(time.Time)
			publishedTo, ok := p.Args["published_to"].(time.Time)
			synopsis, ok := p.Args["synopsis"].(string)
			imageUrl, ok := p.Args["image_url"].(string)

			if ok {
				value := models.Manga{
					Title:         title,
					OriginalTitle: originalTitle,
					EnglishTitle:  englishTitle,
					Status:        status,
					Volumes:       uint(volumes),
					Chapters:      uint(chapters),
					Publishing:    publishing,
					PublishedFrom: publishedFrom,
					PublishedTo:   publishedTo,
					Synopsis:      synopsis,
					ImageUrl:      imageUrl,
				}
				data, err := usecase.AddManga(value)
				if err != nil {
					return nil, err
				}
				return data, nil
			}
			return nil, errors.New("Error")
		},
		Description: "mangaType",
	}
}