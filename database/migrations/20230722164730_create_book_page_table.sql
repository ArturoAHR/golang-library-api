-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE book_page (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  content VARCHAR NOT NULL,
  page_number INT NOT NULL DEFAULT 0,
  book_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_book_to_book_page
    FOREIGN KEY(book_id) 
	    REFERENCES book(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book_page;
-- +goose StatementEnd
