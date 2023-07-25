package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Provider struct {
	Id        uuid.UUID      `gorm:"column:id;primaryKey;type:uuid"`
	Name      string         `gorm:"column:name"`
	Email     string         `gorm:"column:email"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	BookFormats []BookFormat
}

func (Provider) TableName() string {
	return "provider"
}
