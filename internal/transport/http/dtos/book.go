package dtos

type CreateBookDTO struct {
	Id       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Rating   int64  `json:"rating" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

type UpdateBookDTO struct {
	Name     string `json:"name" validate:"required"`
	Rating   int64  `json:"rating" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}

type ReadBookDTO struct {
	Id       int64  `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Rating   int64  `json:"rating" validate:"required"`
	AuthorID int64  `json:"author_id" validate:"required"`
}
