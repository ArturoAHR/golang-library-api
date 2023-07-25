package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id              uuid.UUID      `gorm:"column:id;primaryKey;type:uuid"`
	Name            string         `gorm:"column:name"`
	Author          string         `gorm:"column:author"`
	Pages           int            `gorm:"column:pages"`
	Isbn            string         `gorm:"column:isbn"`
	PublicationDate time.Time      `gorm:"column:publication_date"`
	CreatedAt       time.Time      `gorm:"column:created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at"`

	BookFormats []BookFormat
}

func (Book) TableName() string {
	return "book"
}
