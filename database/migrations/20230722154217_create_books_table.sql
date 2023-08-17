-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE book (
  id UUID NOT NULL DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL,
  author VARCHAR NOT NULL,
  pages INT NOT NULL DEFAULT 0,
  isbn TEXT NOT NULL,
  publication_date DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE book;
-- +goose StatementEnd