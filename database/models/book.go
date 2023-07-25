package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	Id              uuid.UUID      `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	Name            string         `gorm:"column:name" json:"name"`
	Author          string         `gorm:"column:author" json:"author"`
	Pages           int            `gorm:"column:pages" json:"pages"`
	Isbn            string         `gorm:"column:isbn" json:"isbn"`
	PublicationDate time.Time      `gorm:"column:publication_date" json:"publicationDate"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`

	BookFormats []BookFormat `json:"bookFormats,omitempty"`
}

func (Book) TableName() string {
	return "book"
}
