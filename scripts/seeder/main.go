package main

import (
	"fmt"
	"log"
	"os"

	seeder "github.com/ArturoAHR/golang-library-api/database/seeders"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type SeederFunction func(*gorm.DB) error

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
	seedFunctions := []SeederFunction{
		seeder.SeedBooks,
		seeder.SeedFormats,
		seeder.SeedProviders,
		seeder.SeedBookFormats,
		seeder.SeedBookPages,
	}

	for _, seedFunction := range seedFunctions {
		err := seedFunction(DB)
		if err != nil {
			log.Fatalf("Unable to seed: %v", err)
		}
	}

	fmt.Println("Seeding procedure completed")
}
