package core

import (
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
)

func buildBookFromCreateRequest(request CreateBookRequest) book.Book {
	book := book.Book{
		ID:       request.ID,
		Name:     request.Name,
		Rating:   request.Rating,
		AuthorID: request.AuthorID,
	}

	return book
}
