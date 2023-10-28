package author

import (
	"context"
	"fmt"
)

func (s *AuthorService) Update(ctx context.Context, id int64, author *Author) error {
	authorRow := author.MapToModel()
	fmt.Println(authorRow)
	return s.repo.Update(ctx, id, &authorRow)
}
