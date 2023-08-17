# Golang Library API

A library API implemented in the Go Programming Language that exposes endpoints that allows users to receive data of books and book pages of different available formats from different providers. 

## Setup

1. Clone this repository:

   `git@github.com:ArturoAHR/golang-library-api.git`

2. Create the `.env` file based of the `.env.sample` file and fill out the environment variables as needed.

3. Run the setup Makefile script. This will install and setup the necessary tools as well as running the migrations and seeders.

   `make setup`

4. For any subsequent runs after setup use the command:

   `air`

## Endpoints

- GET /book
  
  Retrieves a paginated list of books.

- GET /book/:bookId

  Retrieves a single book record with its related entities.

- GET /book/version/:bookFormatId/page/:pageNumber

  Retrieves a single book page record with its related entities.