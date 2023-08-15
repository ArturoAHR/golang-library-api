package mockdata

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	"github.com/google/uuid"
)

var BaseBookPages = []dbmodels.BookPage{
	{
		Id:           uuid.MustParse("1373b0aa-2c1b-46ae-889f-3f208a225458"),
		PageNumber:   1,
		BookFormatId: uuid.MustParse("bd7be8f1-a3f0-4471-9d2f-1962ea00ff08"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("8f0d4eae-e283-460d-9fe2-0b0b9d8963bf"),
		PageNumber:   2,
		BookFormatId: uuid.MustParse("bd7be8f1-a3f0-4471-9d2f-1962ea00ff08"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("a308343a-966b-49db-86f2-1765399c9e6c"),
		PageNumber:   3,
		BookFormatId: uuid.MustParse("bd7be8f1-a3f0-4471-9d2f-1962ea00ff08"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("fb41edf4-2924-4365-8537-201b629a1fe7"),
		PageNumber:   4,
		BookFormatId: uuid.MustParse("bd7be8f1-a3f0-4471-9d2f-1962ea00ff08"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("e1100881-c47e-41bb-b666-c0d3a0cc4f6b"),
		PageNumber:   5,
		BookFormatId: uuid.MustParse("bd7be8f1-a3f0-4471-9d2f-1962ea00ff08"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},

	{
		Id:           uuid.MustParse("3efa2bcb-4d9f-4d18-a889-e9f95e823466"),
		PageNumber:   1,
		BookFormatId: uuid.MustParse("1cb7f509-a94f-4d71-b010-da1288501dd7"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("9cc26c91-8c51-439c-bef5-0a7e6d4c4f9f"),
		PageNumber:   2,
		BookFormatId: uuid.MustParse("1cb7f509-a94f-4d71-b010-da1288501dd7"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("4d92f9f9-c1b9-4bca-8f27-97e20ec74637"),
		PageNumber:   3,
		BookFormatId: uuid.MustParse("1cb7f509-a94f-4d71-b010-da1288501dd7"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("4b128190-e645-47bf-9729-456de46434ce"),
		PageNumber:   4,
		BookFormatId: uuid.MustParse("1cb7f509-a94f-4d71-b010-da1288501dd7"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("f17dbd39-7789-44e2-88e5-aae428cecba3"),
		PageNumber:   5,
		BookFormatId: uuid.MustParse("1cb7f509-a94f-4d71-b010-da1288501dd7"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},

	{
		Id:           uuid.MustParse("d2950ccd-de14-4b70-bc12-e5f71c58d4c2"),
		PageNumber:   1,
		BookFormatId: uuid.MustParse("54593594-98c9-4632-8ff6-79c14b058944"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("afbf763c-3c3f-4744-bfe1-bb776fd15c14"),
		PageNumber:   2,
		BookFormatId: uuid.MustParse("54593594-98c9-4632-8ff6-79c14b058944"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("c3797c7c-d7c7-4b82-ba15-3e4fab2d89a5"),
		PageNumber:   3,
		BookFormatId: uuid.MustParse("54593594-98c9-4632-8ff6-79c14b058944"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("9e85bb7c-3e34-4151-a82c-17c72a144138"),
		PageNumber:   4,
		BookFormatId: uuid.MustParse("54593594-98c9-4632-8ff6-79c14b058944"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, {
		Id:           uuid.MustParse("6bd17171-a7d8-4ac1-9a4d-d17167b160ba"),
		PageNumber:   5,
		BookFormatId: uuid.MustParse("54593594-98c9-4632-8ff6-79c14b058944"),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	},
}

func GetBookPages() ([]dbmodels.BookPage, error) {
	bookPages := BaseBookPages
	const bookPagesDir = "./mocks/data/book-pages"

	for i, bookPage := range BaseBookPages {
		bookFormat, err := findBookFormatById(bookPage.BookFormatId)
		if err != nil {
			return nil, err
		}

		book, err := findBookById(bookFormat.BookId)
		if err != nil {
			return nil, err
		}

		bookPagePath := filepath.Join(bookPagesDir, book.Isbn, fmt.Sprintf("page-%03d.txt", bookPage.PageNumber))

		pageContent, err := os.ReadFile(bookPagePath)
		if err != nil {
			return nil, err
		}

		bookPages[i].Content = string(pageContent)
	}

	return bookPages, nil
}
