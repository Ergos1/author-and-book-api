package core

import (
	"context"
	"fmt"
)

type UpdateAuthorRequest struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (s *Service) UpdateAuthor(ctx context.Context, request UpdateAuthorRequest) error {
	author := buildAuthorFromUpdateRequest(request)
	fmt.Println("HI")
	return s.authorService.Update(ctx, author.ID, &author)
}
