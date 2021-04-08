package mysql_repository

import (
	"github.com/jinzhu/gorm"
	"manga-server/domain/dto"
	"manga-server/domain/models"
	"manga-server/pkg/mysql"
	"math"
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
	GetManga(page int, pageSize int) (*int,[]dto.DBGetManga, error)
	AddManga(data models.Manga) (*models.Manga, error)
}

func (m MysqlRepository) GetManga(page int, pageSize int) (*int,[]dto.DBGetManga, error) {
	var data []dto.DBGetManga
	var count int
	err := m.db.Model(&models.Manga{}).Count(&count).Error
	err = m.db.Scopes(mysql.Paginate(uint(page), uint(pageSize))).Order("created_at DESC").Find(&data).Error
	if err != nil {
		return nil, nil, err
	}
	counts := math.Ceil(float64(count)/float64(pageSize))
	count = int(counts)
	return &count, data, nil
}

func (m MysqlRepository) AddManga(data models.Manga) (*models.Manga, error)  {
	err := m.db.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

