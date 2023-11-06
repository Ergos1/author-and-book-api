package dtos

import "gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"

type CreateAuthorDTO struct {
	Id   int64  `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateAuthorDTO struct {
	Name string `json:"name" validate:"required"`
}

type ReadAuthorDTO struct {
	Id    int64         `json:"id" validate:"required"`
	Name  string        `json:"name" validate:"required"`
	Books []ReadBookDTO `json:"books"`
}

func (d *ReadAuthorDTO) MapFromAuthorWithBooks(author *core.AuthorWithBooks) {
	d.Id = author.ID
	d.Name = author.Name
	d.Books = MapFromBooks(author.Books)
}
