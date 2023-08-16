package bookmodule

import (
	"net/http"

	bookmoduletypes "github.com/ArturoAHR/golang-library-api/internal/modules/book/types"
	"github.com/ArturoAHR/golang-library-api/internal/utils"
)

type BookController struct {
	service BookService
}

func (c *BookController) getBooks(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()

	query := bookmoduletypes.GetBooksQueryDto{
		Page:     queryParams.Get("page"),
		PageSize: queryParams.Get("pageSize"),
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

func (c *BookController) getBook(writer http.ResponseWriter, request *http.Request) {

}

func (c *BookController) getBookPage(writer http.ResponseWriter, request *http.Request) {

}
