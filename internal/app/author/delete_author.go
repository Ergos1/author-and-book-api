package author

import "context"

func (s *AuthorService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
