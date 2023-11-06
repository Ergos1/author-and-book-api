package core

import (
	"context"
)

type UpdateAuthorRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *Service) UpdateAuthor(ctx context.Context, request UpdateAuthorRequest) error {
	author := buildAuthorFromUpdateRequest(request)
	return s.authorService.Update(ctx, author.ID, &author)
}
