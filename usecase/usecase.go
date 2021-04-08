package usecase

import (
	"manga-server/domain/dto"
	"manga-server/domain/models"
	"manga-server/repository/mysql_repository"
)

type usecase struct {
	Repository   mysql_repository.Repository
}

func NewUseCaseImpl(repository mysql_repository.Repository) UseCase {
	return &usecase{
		Repository:   repository,
	}
}

type UseCase interface {
	GetManga(page uint, pageSize uint) ([]dto.DBGetManga, *uint, error)
	AddManga(manga models.Manga) (*models.Manga, error)
}

func (u usecase) GetManga(page uint, pageSize uint) ([]dto.DBGetManga, *uint, error) {
	count, data, err := u.Repository.GetManga(page, pageSize)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}

func (u usecase) AddManga(manga models.Manga) (*models.Manga, error) {

	//AnimeGenreSlug := request.Slug
	//if AnimeGenreSlug == nil {
	//	a := slug.Make(request.Name)
	//	AnimeGenreSlug = &a
	//}

	data, err := u.Repository.AddManga(manga)

	if err != nil {
		return nil, err
	}
	return data, nil
}