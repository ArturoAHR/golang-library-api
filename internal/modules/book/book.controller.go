package bookmodule

import "net/http"

type BookController struct {
	service BookService
}

func (b *BookController) getBooks(writer http.ResponseWriter, request *http.Request) {

}

func (b *BookController) getBook(writer http.ResponseWriter, request *http.Request) {

}

func (b *BookController) getBookPage(writer http.ResponseWriter, request *http.Request) {

}
