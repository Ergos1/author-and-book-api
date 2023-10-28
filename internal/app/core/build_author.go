package core

import "gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"

func buildAuthorFromCreateRequest(request CreateAuthorRequest) author.Author {
	author := author.Author{
		ID:   request.ID,
		Name: request.Name,
	}

	return author
}

func buildAuthorFromUpdateRequest(request UpdateAuthorRequest) author.Author {
	author := author.Author{
		ID:   request.ID,
		Name: request.Name,
	}

	return author
}
