package author

import "context"

func (s *AuthorService) Update(ctx context.Context, id int64, author *Author) error {
	authorRow := author.MapToModel()
	return s.repo.Update(ctx, id, &authorRow)
}
