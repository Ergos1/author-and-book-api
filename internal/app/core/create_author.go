package core

import "context"

type CreateAuthorRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *Service) CreateAuthor(ctx context.Context, request CreateAuthorRequest) (int64, error) {
	author := buildAuthorFromCreateRequest(request)
	return s.authorService.Create(ctx, &author)
}
