package core

import (
	"context"
)

func (s *Service) DeleteAuthorById(ctx context.Context, id int64) error {
	err := s.authorService.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
