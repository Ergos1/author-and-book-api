package core

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
)

type AuthorWithBooks struct {
	*author.Author
	Books []*book.Book
}

func (s *Service) GetAuthorById(ctx context.Context, id int64) (*AuthorWithBooks, error) {
	author, err := s.authorService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	books, err := s.bookService.GetByAuthorID(ctx, id)
	if err != nil {
		return nil, err
	}

	authorWithBooks := &AuthorWithBooks{
		Author: author,
		Books:  books,
	}

	return authorWithBooks, nil
}
