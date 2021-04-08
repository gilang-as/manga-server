package models

import (
	"manga-server/pkg/mysql"
	"time"
)

type Manga struct {
	mysql.Model
	Title         string    `json:"title"`
	OriginalTitle string    `json:"original_title"`
	EnglishTitle  string    `json:"english_title"`
	Status        string    `json:"status"`
	Volumes       uint      `json:"volumes"`
	Chapters      uint      `json:"chapters"`
	Publishing    bool      `json:"publishing"`
	PublishedFrom time.Time `json:"published_from"`
	PublishedTo   time.Time `json:"published_to"`
	Synopsis      string    `json:"synopsis" gorm:"type:text;"`
	ImageUrl      string    `json:"image_url"`
}

func (Manga) TableName() string { return "manga" }