package main

import (
	dbconnection "github.com/ArturoAHR/golang-library-api/internal/database"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("no .env file found")
	}

	dbconnection.GetInstance()
}

func main() {

}
