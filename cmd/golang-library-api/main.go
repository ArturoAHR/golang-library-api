package main

import (
	"fmt"
	"net/http"
	"os"

	dbconnection "github.com/ArturoAHR/golang-library-api/internal/database"
	bookmodule "github.com/ArturoAHR/golang-library-api/internal/modules/book"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("no .env file found")
	}

	dbconnection.GetInstance()
}

func main() {
	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	router := mux.NewRouter()
	router.PathPrefix("/book").HandlerFunc(bookmodule.Handler)

	http.Handle("/", router)
	http.ListenAndServe(appPort, nil)
}
