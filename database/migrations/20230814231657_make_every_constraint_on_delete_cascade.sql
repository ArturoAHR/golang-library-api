-- +goose Up
-- +goose StatementBegin

-- Book to Book Format
ALTER TABLE book_format
DROP CONSTRAINT fk_book_to_book_format;

ALTER TABLE book_format
ADD CONSTRAINT fk_book_to_book_format  
FOREIGN KEY(book_id) REFERENCES book(id)
ON DELETE CASCADE;

-- Format to Book Format
ALTER TABLE book_format
DROP CONSTRAINT fk_format_to_book_format;

ALTER TABLE book_format
ADD CONSTRAINT fk_format_to_book_format  
FOREIGN KEY(format_id) REFERENCES format(id)
ON DELETE CASCADE;

-- Provider to Book Format
ALTER TABLE book_format
DROP CONSTRAINT fk_provider_to_book_format;

ALTER TABLE book_format
ADD CONSTRAINT fk_provider_to_book_format  
FOREIGN KEY(provider_id) REFERENCES provider(id)
ON DELETE CASCADE;

-- Book Format to Book Page
ALTER TABLE book_page
DROP CONSTRAINT fk_book_format_to_book_page;

ALTER TABLE book_page
ADD CONSTRAINT fk_book_format_to_book_page  
FOREIGN KEY(book_format_id) REFERENCES book_format(id)
ON DELETE CASCADE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Book to Book Format
ALTER TABLE book_format
DROP CONSTRAINT fk_book_to_book_format;

ALTER TABLE book_format
ADD CONSTRAINT fk_book_to_book_format  
FOREIGN KEY(book_id) REFERENCES book(id);

-- Format to Book Format
ALTER TABLE book_format
DROP CONSTRAINT fk_format_to_book_format;

ALTER TABLE book_format
ADD CONSTRAINT fk_format_to_book_format  
FOREIGN KEY(format_id) REFERENCES format(id);

-- Provider to Book Format
ALTER TABLE book_format
DROP CONSTRAINT fk_provider_to_book_format;

ALTER TABLE book_format
ADD CONSTRAINT fk_provider_to_book_format  
FOREIGN KEY(provider_id) REFERENCES provider(id);

-- Book Format to Book Page
ALTER TABLE book_page
DROP CONSTRAINT fk_book_format_to_book_page;

ALTER TABLE book_page
ADD CONSTRAINT fk_book_format_to_book_page  
FOREIGN KEY(book_format_id) REFERENCES book_format(id);

-- +goose StatementEnd
