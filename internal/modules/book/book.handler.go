package bookmodule

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	subRouter := mux.NewRouter()
	controller := bookController{}

	subRouter.HandleFunc("/book", controller.getBooks).Methods("GET")
	subRouter.HandleFunc("/book/{bookId}", controller.getBook).Methods("GET")
	subRouter.HandleFunc("/book/{bookId}/format/{formatId}", controller.getBookPage).Methods("GET")

	subRouter.ServeHTTP(writer, request)
}
