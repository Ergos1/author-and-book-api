package core

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
)

func (s *Service) GetBookById(ctx context.Context, id int64) (*book.Book, error) {
	book, err := s.bookService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return book, nil
}
