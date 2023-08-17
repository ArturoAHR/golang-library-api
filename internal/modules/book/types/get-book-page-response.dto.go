package bookmoduletypes

import dbmodels "github.com/ArturoAHR/golang-library-api/database/models"

type GetBookPageResponseDto struct {
	BookPage dbmodels.BookPage `json:"bookPage"`
}
