package bookmodule

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	subRouter := mux.NewRouter()

	subRouter.HandleFunc("/book", getBooks).Methods("GET")
	subRouter.HandleFunc("/book/{bookId}", getBook).Methods("GET")
	subRouter.HandleFunc("/book/{bookId}/format/{formatId}", getBookPage).Methods("GET")

	subRouter.ServeHTTP(writer, request)
}
