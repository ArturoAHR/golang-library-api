package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookPage struct {
	Id           uuid.UUID      `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	Content      string         `gorm:"column:content" json:"content"`
	PageNumber   int            `gorm:"column:page_number" json:"pageNumber"`
	BookFormatId uuid.UUID      `gorm:"column:book_format_id;type:uuid" json:"bookFormatId"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`

	BookFormat BookFormat `gorm:"foreignKey:BookFormatId" json:"bookFormat,omitempty"`
}

func (BookPage) TableName() string {
	return "book_page"
}
