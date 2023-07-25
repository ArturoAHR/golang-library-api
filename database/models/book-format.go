package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookFormat struct {
	Id         uuid.UUID      `gorm:"column:id;primaryKey;type:uuid"`
	BookId     uuid.UUID      `gorm:"column:book_id;type:uuid"`
	FormatId   uuid.UUID      `gorm:"column:format_id;type:uuid"`
	ProviderId uuid.UUID      `gorm:"column:provider_id;type:uuid"`
	CreatedAt  time.Time      `gorm:"column:created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at"`

	Book     Book     `gorm:"foreignKey:BookId"`
	Format   Format   `gorm:"foreignKey:FormatId"`
	Provider Provider `gorm:"foreignKey:ProviderId"`

	BookPages []BookPage
}

func (BookFormat) TableName() string {
	return "book_format"
}
