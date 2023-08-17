package bookmoduletypes

import dbmodels "github.com/ArturoAHR/golang-library-api/database/models"

type GetBookByIdResponseDto struct {
	Book dbmodels.Book `json:"book"`
}
