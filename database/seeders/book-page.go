package seeder

import (
	mockdata "github.com/ArturoAHR/golang-library-api/mocks/data"
	"gorm.io/gorm"
)

func SeedBookPages(db *gorm.DB) error {
	bookPages, err := mockdata.GetBookPages()
	if err != nil {
		return err
	}

	return seedRecords(db, bookPages)
}
