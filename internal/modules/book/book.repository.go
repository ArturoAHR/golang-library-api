package bookmodule

import "gorm.io/gorm"

type BookRepository struct {
	db *gorm.DB
}
