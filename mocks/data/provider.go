package mockdata

import (
	"time"

	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	"github.com/google/uuid"
)

var Providers = []dbmodels.Provider{
	{
		Id:        uuid.MustParse("48454a04-ac64-4d34-9f90-264af31ab387"),
		Name:      "Penguin Random House",
		Email:     "info@prh.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.MustParse("a506771c-9eed-4a1a-b823-52db1f6c6ca6"),
		Name:      "Dzanc Books",
		Email:     "business@dzanc.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
