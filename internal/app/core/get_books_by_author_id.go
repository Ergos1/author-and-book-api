package core

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
)

func (s *Service) GetBooksByAuthorID(ctx context.Context, authorID int64) ([]*book.Book, error) {
	books, err := s.bookService.GetByAuthorID(ctx, authorID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
