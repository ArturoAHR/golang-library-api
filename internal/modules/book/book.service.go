package bookmodule

import (
	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	bookmoduletypes "github.com/ArturoAHR/golang-library-api/internal/modules/book/types"
	"github.com/ArturoAHR/golang-library-api/internal/utils"
)

type BookService struct {
	repository BookRepository
}

func (s *BookService) GetBooks(filter bookmoduletypes.GetBooksFilter) ([]dbmodels.Book, utils.PaginationMetadata, error) {
	books, err := s.repository.GetBooks(filter)
	if err != nil {
		return nil, utils.PaginationMetadata{}, err
	}

	count, err := s.repository.CountBooks(filter)
	if err != nil {
		return nil, utils.PaginationMetadata{}, err
	}

	metadata := utils.PaginationMetadata{
		TotalItems: count,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalPages: count / filter.PageSize,
	}

	return books, metadata, nil
}
