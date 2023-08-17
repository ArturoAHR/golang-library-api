package bookmodule

import (
	dbmodels "github.com/ArturoAHR/golang-library-api/database/models"
	bookmoduletypes "github.com/ArturoAHR/golang-library-api/internal/modules/book/types"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func (r *BookRepository) GetBooks(filters bookmoduletypes.GetBooksFilter) ([]dbmodels.Book, error) {
	var books []dbmodels.Book

	offset := (filters.Page - 1) * filters.PageSize
	limit := filters.PageSize

	result := r.db.Offset(offset).Limit(limit).Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}

	return books, nil
}

func (r *BookRepository) GetBookById(bookId string) (dbmodels.Book, error) {
	var book dbmodels.Book

	result := r.db.Preload("BookFormats").Preload("BookFormats.Format").Preload("BookFormats.Provider").Where("id = ?", bookId).First(&book)
	if result.Error != nil {
		return dbmodels.Book{}, result.Error
	}

	return book, nil
}

func (r *BookRepository) CountBooks(filters bookmoduletypes.GetBooksFilter) (int, error) {
	var count int64

	result := r.db.Model(&dbmodels.Book{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *BookRepository) GetBookPage(bookFormatId string, pageNumber int) (dbmodels.BookPage, error) {
	var bookPage dbmodels.BookPage

	result := r.db.Preload("BookFormat").Preload("BookFormat.Format").Preload("BookFormat.Provider").Preload("BookFormat.Book").Where("book_format_id = ? AND page_number = ?", bookFormatId, pageNumber).First(&bookPage)
	if result.Error != nil {
		return dbmodels.BookPage{}, result.Error
	}

	return bookPage, nil
}
