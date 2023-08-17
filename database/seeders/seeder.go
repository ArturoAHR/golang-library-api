package seeder

import (
	"gorm.io/gorm"
)

func seedRecords[T any](db *gorm.DB, records []T) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, record := range records {
			if err := tx.Create(&record).Error; err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
