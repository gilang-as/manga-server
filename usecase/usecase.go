package usecase

import (
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
	GetManga(page uint, pageSize uint) ([]models.Manga, *uint, error)
}

func (u usecase) GetManga(page uint, pageSize uint) ([]models.Manga, *uint, error) {
	count, data, err := u.Repository.GetManga(page, pageSize)
	if err != nil {
		return nil, nil, err
	}
	return data, count, nil
}