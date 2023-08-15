package mockdata

import (
	"time"

	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	parsers "github.com/ArturoAHR/golang-library-api/internal/utils"
	"github.com/google/uuid"
)

var Books = []dbmodels.Book{
	{
		Id:              uuid.MustParse("3084da87-182f-4da4-8d3d-948b35cc6f6c"),
		Name:            "Viaje al Centro de la Tierra",
		Author:          "Julio Verne",
		Pages:           5,
		Isbn:            "978-8447326327",
		PublicationDate: parsers.ParseDate("02-01-2006", "25-11-1864"),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	},
	{
		Id:              uuid.MustParse("d3b9f6f4-c8ac-41d2-9d42-80417bab5961"),
		Name:            "Matilda",
		Author:          "Roald, Dahl",
		Pages:           5,
		Isbn:            "978-0140327595",
		PublicationDate: parsers.ParseDate("28-09-1989", "25-11-1864"),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}, {
		Id:              uuid.MustParse("8af3b633-86c8-4688-a99c-65e94c7eadc8"),
		Name:            "La Odisea",
		Author:          "Homero",
		Pages:           5,
		Isbn:            "978-1539427698",
		PublicationDate: parsers.ParseDate("09-10-2016", "25-11-1864"),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	},
}
