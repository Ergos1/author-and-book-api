package core

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
)

func (s *Service) GetAuthorById(ctx context.Context, id int64) (*author.Author, error) {
	author, err := s.authorService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return author, nil
}
