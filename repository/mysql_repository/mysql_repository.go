package mysql_repository

import (
	"github.com/jinzhu/gorm"
	"manga-server/domain/models"
	"manga-server/pkg/mysql"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) Repository {
	return &MysqlRepository{
		db: db,
	}
}

type Repository interface {
	GetManga(page uint, pageSize uint) (*uint,[]models.Manga, error)
}

func (m MysqlRepository) GetManga(page uint, pageSize uint) (*uint,[]models.Manga, error) {
	var data []models.Manga
	var count uint
	err := m.db.Model(&models.Manga{}).Count(&count).Error
	err = m.db.Scopes(mysql.Paginate(page, pageSize)).Find(&data).Error
	if err != nil {
		return nil, nil, err
	}
	return &count, data, nil
}

