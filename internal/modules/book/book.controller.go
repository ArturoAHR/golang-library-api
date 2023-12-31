package bookmodule

import (
	"net/http"
	"strconv"

	bookmoduletypes "github.com/ArturoAHR/golang-library-api/internal/modules/book/types"
	"github.com/ArturoAHR/golang-library-api/internal/utils"
	"github.com/gorilla/mux"
)

type BookController struct {
	service BookService
}

// GET /book
//
// Retrieves a paginated list of books.
func (c *BookController) getBooks(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()

	query := bookmoduletypes.GetBooksQueryDto{
		Page:     queryParams.Get("page"),
		PageSize: queryParams.Get("pageSize"),
	}

	if query.Page == "" {
		query.Page = "1"
	}

	if query.PageSize == "" {
		query.PageSize = "10"
	}

	err := utils.ValidateStruct(query)
	if err != nil {
		utils.SendJSONError(writer, err, http.StatusBadRequest)
		return
	}

	filter := bookmoduletypes.GetBooksFilter{
		Page:     utils.MustAtoi(query.Page),
		PageSize: utils.MustAtoi(query.PageSize),
	}

	books, metadata, err := c.service.GetBooks(filter)
	if err != nil {
		utils.SendJSONError(writer, err, http.StatusInternalServerError)
		return
	}

	response := bookmoduletypes.GetBooksResponseDto{
		Books:    books,
		Metadata: metadata,
	}

	utils.SendJSONResponse(writer, response, http.StatusOK)
}

// GET /book/:bookId
//
// Retrieves a book by its id, with its tied entities except for pages.
func (c *BookController) getBook(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]

	book, err := c.service.GetBookById(bookId)
	if err != nil {
		utils.SendJSONError(writer, err, http.StatusInternalServerError)
		return
	}

	response := bookmoduletypes.GetBookByIdResponseDto{
		Book: book,
	}

	utils.SendJSONResponse(writer, response, http.StatusOK)
}

// GET /book/version/:versionId/page/:pageNumber
//
// Gets a book page from a book format id and a page number, with its tied entities.
func (c *BookController) getBookPage(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	bookFormatId := vars["versionId"]
	pageNumberString := vars["pageNumber"]

	if pageNumberString == "" {
		pageNumberString = "1"
	}

	pageNumber, err := strconv.Atoi(pageNumberString)
	if err != nil {
		utils.SendJSONError(writer, err, http.StatusInternalServerError)
		return
	}

	bookPage, err := c.service.GetBookPage(bookFormatId, pageNumber)
	if err != nil {
		utils.SendJSONError(writer, err, http.StatusInternalServerError)
		return
	}

	response := bookmoduletypes.GetBookPageResponseDto{
		BookPage: bookPage,
	}

	utils.SendJSONResponse(writer, response, http.StatusOK)
}
