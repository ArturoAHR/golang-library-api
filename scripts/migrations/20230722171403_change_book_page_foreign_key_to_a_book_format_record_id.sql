-- +goose Up
-- +goose StatementBegin
ALTER TABLE book_page 
DROP CONSTRAINT fk_book_to_book_page;

-- Rename the column
ALTER TABLE book_page 
RENAME COLUMN book_id TO book_format_id;

-- Add the new foreign key constraint
ALTER TABLE book_page 
ADD CONSTRAINT fk_book_format_to_book_page 
FOREIGN KEY(book_format_id) REFERENCES book_format(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE book_page 
DROP CONSTRAINT fk_book_format_to_book_page;

-- Rename the column
ALTER TABLE book_page 
RENAME COLUMN book_format_id TO book_id;

-- Add the new foreign key constraint
ALTER TABLE book_page 
ADD CONSTRAINT fk_book_to_book_page 
FOREIGN KEY(book_id) REFERENCES book(id);
-- +goose StatementEnd
