package dtos

import "gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"

type CreateBookDTO struct {
	Id       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Rating   int64  `json:"rating" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

type UpdateBookDTO struct {
	Name     string `json:"name" validate:"required"`
	Rating   int64  `json:"rating" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

type ReadBookDTO struct {
	Id       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Rating   int64  `json:"rating" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

func (d *ReadBookDTO) MapFromBook(book *book.Book) {
	d.Id = book.ID
	d.Name = book.Name
	d.Rating = book.Rating
	d.AuthorID = book.AuthorID
}

func MapFromBooks(books []*book.Book) []ReadBookDTO {
	booksDto := make([]ReadBookDTO, len(books))

	for _, book := range books {
		bookDto := &ReadBookDTO{}
		bookDto.MapFromBook(book)

		booksDto = append(booksDto, *bookDto)
	}

	return booksDto
}
