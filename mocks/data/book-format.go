package mockdata

import (
	"time"

	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	"github.com/google/uuid"
)

var BookFormats = []dbmodels.BookFormat{
	{
		Id:         uuid.MustParse("bd7be8f1-a3f0-4471-9d2f-1962ea00ff08"),
		BookId:     uuid.MustParse("3084da87-182f-4da4-8d3d-948b35cc6f6c"),
		FormatId:   uuid.MustParse("460832c8-fece-4f17-a6f8-482b10eec265"),
		ProviderId: uuid.MustParse("48454a04-ac64-4d34-9f90-264af31ab387"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
	{
		Id:         uuid.MustParse("1cb7f509-a94f-4d71-b010-da1288501dd7"),
		BookId:     uuid.MustParse("d3b9f6f4-c8ac-41d2-9d42-80417bab5961"),
		FormatId:   uuid.MustParse("93f4731b-2298-406f-9115-3b84728382e0"),
		ProviderId: uuid.MustParse("48454a04-ac64-4d34-9f90-264af31ab387"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, {
		Id:         uuid.MustParse("54593594-98c9-4632-8ff6-79c14b058944"),
		BookId:     uuid.MustParse("8af3b633-86c8-4688-a99c-65e94c7eadc8"),
		FormatId:   uuid.MustParse("460832c8-fece-4f17-a6f8-482b10eec265"),
		ProviderId: uuid.MustParse("a506771c-9eed-4a1a-b823-52db1f6c6ca6"),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	},
}
