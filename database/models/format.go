package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Format struct {
	Id        uuid.UUID      `gorm:"column:id;primaryKey;type:uuid"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	BookFormats []BookFormat
}

func (Format) TableName() string {
	return "format"
}
