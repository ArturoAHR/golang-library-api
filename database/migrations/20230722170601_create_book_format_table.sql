-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE book_format (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  book_id UUID NOT NULL,
  format_id UUID NOT NULL,
  provider_id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_book_to_book_format
    FOREIGN KEY(book_id) 
	    REFERENCES book(id),
  CONSTRAINT fk_format_to_book_format
    FOREIGN KEY(format_id) 
	    REFERENCES format(id),
  CONSTRAINT fk_provider_to_book_format
    FOREIGN KEY(provider_id) 
	    REFERENCES provider(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book_format;
-- +goose StatementEnd
