package core

import "context"

type CreateBookRequest struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Rating   int64  `json:"raing"`
	AuthorID int64  `json:"author_id"`
}

func (s *Service) CreateBook(ctx context.Context, request CreateBookRequest) (int64, error) {
	book := buildBookFromCreateRequest(request)
	return s.bookService.Create(ctx, &book)
}
