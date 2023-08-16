package dbconnection

import (
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var once sync.Once
var db *gorm.DB

func GetInstance() *gorm.DB {
	once.Do(func() {
		var err error
		dsn := os.Getenv("GORM_DSN")

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}
	})

	return db
}
