package dbmodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Provider struct {
	Id        uuid.UUID      `gorm:"column:id;primaryKey;type:uuid" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Email     string         `gorm:"column:email" json:"email"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`

	BookFormats []BookFormat `json:"bookFormats,omitempty"`
}

func (Provider) TableName() string {
	return "provider"
}
