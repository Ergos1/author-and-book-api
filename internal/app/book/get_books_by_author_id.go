package book

import (
	"context"
)

func (s *BookService) GetByAuthorID(ctx context.Context, id int64) ([]*Book, error) {
	bookRows, err := s.repo.GetByAuthorId(ctx, id)
	if err != nil {
		return nil, err
	}

	books := MapFromModels(bookRows)
	return books, err
}
