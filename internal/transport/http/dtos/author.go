package dtos

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
