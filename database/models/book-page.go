package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookPage struct {
	Id           uuid.UUID      `gorm:"column:id;primaryKey;type:uuid"`
	Content      string         `gorm:"column:content"`
	PageNumber   int            `gorm:"column:page_number"`
	BookFormatId uuid.UUID      `gorm:"column:book_format_id;type:uuid"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at"`

	BookFormat BookFormat `gorm:"foreignKey:BookFormatId"`
}

func (BookPage) TableName() string {
	return "book_page"
}
