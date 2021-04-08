package mysql

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        string     `gorm:"type:varchar(255);primary_key; not null;" json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
}

func (base *Model) BeforeCreate(scope *gorm.Scope) error {
	uuid,_ := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}