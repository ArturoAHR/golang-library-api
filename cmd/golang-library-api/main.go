package main

import (
	"net/http"

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
	router := mux.NewRouter()
	router.PathPrefix("/book").HandlerFunc(bookmodule.Handler)

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
