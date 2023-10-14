package core

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
)

type AuthorService interface {
	GetByID(ctx context.Context, id int64) (*author.Author, error)
	Create(ctx context.Context, author *author.Author) (int64, error)
	Update(ctx context.Context, id int64, author *author.Author) error
	Delete(ctx context.Context, id int64) error
}

type BookService interface {
	GetByID(ctx context.Context, id int64) (*book.Book, error)
	Create(ctx context.Context, book *book.Book) (int64, error)
	// Update(ctx context.Context, id int64, book *book.Book) error
	// Delete(ctx context.Context, id int64) error
	GetByAuthorID(ctx context.Context, authorID int64) ([]*book.Book, error)
}

type Service struct {
	authorService AuthorService
	bookService   BookService
}

func NewService(as AuthorService, bs BookService) *Service {
	return &Service{
		authorService: as,
		bookService:   bs,
	}
}
