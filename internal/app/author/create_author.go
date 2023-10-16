package author

import (
	"context"
)

func (s *AuthorService) Create(ctx context.Context, author *Author) (int64, error) {
	authorRow := author.MapToModel()
	return s.repo.Create(ctx, &authorRow)
}
