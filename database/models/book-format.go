package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookFormat struct {
	Id         uuid.UUID      `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	BookId     uuid.UUID      `gorm:"column:book_id;type:uuid" json:"bookId"`
	FormatId   uuid.UUID      `gorm:"column:format_id;type:uuid" json:"formatId"`
	ProviderId uuid.UUID      `gorm:"column:provider_id;type:uuid" json:"providerId"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`

	Book     Book     `gorm:"foreignKey:BookId" json:"book,omitempty"`
	Format   Format   `gorm:"foreignKey:FormatId" json:"format,omitempty"`
	Provider Provider `gorm:"foreignKey:ProviderId" json:"provider,omitempty"`

	BookPages []BookPage `json:"bookPages,omitempty"`
}

func (BookFormat) TableName() string {
	return "book_format"
}
