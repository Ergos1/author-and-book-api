package book

import "context"

func (s *BookService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
