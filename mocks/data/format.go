package mockdata

import (
	"time"

	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	"github.com/google/uuid"
)

var Formats = []dbmodels.Format{
	{
		Id:        uuid.MustParse("93f4731b-2298-406f-9115-3b84728382e0"),
		Name:      "html",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		Id:        uuid.MustParse("460832c8-fece-4f17-a6f8-482b10eec265"),
		Name:      "plain",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}
