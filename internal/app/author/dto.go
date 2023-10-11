package author

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (a *Author) MapFromModel(author AuthorRow) *Author {
	a.ID = author.ID
	a.Name = author.Name

	return a
}

func (a *Author) MapToModel() AuthorRow {
	return AuthorRow{
		ID:   a.ID,
		Name: a.Name,
	}
}

func MapToModels(authors []Author) []AuthorRow {
	authorRows := make([]AuthorRow, len(authors))
	for _, author := range authors {
		authorRows = append(authorRows, author.MapToModel())
	}

	return authorRows
}
