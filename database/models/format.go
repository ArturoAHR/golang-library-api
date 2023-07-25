package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Format struct {
	Id        uuid.UUID      `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`

	BookFormats []BookFormat `json:"bookFormats,omitempty"`
}

func (Format) TableName() string {
	return "format"
}
