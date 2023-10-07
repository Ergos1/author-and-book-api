package dtos

type CreateAuthorDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type UpdateAuthorDTO struct {
	Name string `json:"name"`
}
