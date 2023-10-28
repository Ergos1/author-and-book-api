package book

import "context"

func (s *BookService) Create(ctx context.Context, book *Book) (int64, error) {
	bookRow := book.MapToModel()
	return s.repo.Create(ctx, &bookRow)
}
