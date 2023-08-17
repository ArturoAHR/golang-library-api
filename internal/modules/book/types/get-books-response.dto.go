package bookmoduletypes

import (
	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	"github.com/ArturoAHR/golang-library-api/internal/utils"
)

type GetBooksResponseDto struct {
	Metadata utils.PaginationMetadata `json:"metadata"`
	Books    []dbmodels.Book          `json:"books"`
}
