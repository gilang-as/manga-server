package dto

import (
	"time"
)

type ResponseGetManga struct {
	Total uint `json:"total"`
	Manga interface{} `json:"manga"`
}

type DBGetManga struct {
	ID        string     `gorm:"type:varchar(255);primary_key; not null;" json:"id,omitempty"`
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
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
}

func (DBGetManga) TableName() string { return "manga" }