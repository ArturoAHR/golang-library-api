package bookmodule

import (
	"net/http"

	dbconnection "github.com/ArturoAHR/golang-library-api/internal/database"
	"github.com/gorilla/mux"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	subRouter := mux.NewRouter()

	controller := BookController{
		service: BookService{
			repository: BookRepository{
				db: dbconnection.GetInstance(),
			},
		},
	}

	subRouter.HandleFunc("/book", controller.getBooks).Methods("GET")
	subRouter.HandleFunc("/book/{bookId}", controller.getBook).Methods("GET")
	subRouter.HandleFunc("/book/version/{versionId}", controller.getBookPage).Methods("GET")

	subRouter.ServeHTTP(writer, request)
}
