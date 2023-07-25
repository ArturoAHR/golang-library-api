package main

import (
	"fmt"
	"os"
	"time"

	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	dsn := os.Getenv("GORM_DSN")
	tmpDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = tmpDB
}

func main() {

	book := dbmodels.Book{
		Id:              uuid.New(),
		Name:            "Test Book",
		Author:          "Test Author",
		Pages:           200,
		Isbn:            "1234567890",
		PublicationDate: time.Now(),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	DB.Create(&book)

	fmt.Printf("Book: %+v\n", book)

	// Query and print
	var queriedBook dbmodels.Book
	DB.Preload("BookFormats.Format").Preload("BookFormats.Provider").Preload("BookFormats.BookPages").Find(&queriedBook, book.Id)
	fmt.Printf("Book: %+v\n", queriedBook)

}
