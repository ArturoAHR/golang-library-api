package seeder

import (
	mockdata "github.com/ArturoAHR/golang-library-api/mocks/data"
	"gorm.io/gorm"
)

func SeedProviders(db *gorm.DB) error {
	return seedRecords(db, mockdata.Providers)
}
