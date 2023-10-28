package book

type Book struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Rating   int64  `json:"rating"`
	AuthorID int64  `json:"author_id"`
}

func (b *Book) MapFromModel(book BookRow) *Book {
	b.ID = book.ID
	b.Name = book.Name
	b.Rating = book.Rating
	b.AuthorID = book.AuthorID

	return b
}

func (b *Book) MapToModel() BookRow {
	return BookRow{
		ID:       b.ID,
		Name:     b.Name,
		Rating:   b.Rating,
		AuthorID: b.AuthorID,
	}
}

func MapToModels(books []Book) []BookRow {
	bookRows := make([]BookRow, 0, len(books))
	for _, book := range books {
		bookRows = append(bookRows, book.MapToModel())
	}

	return bookRows
}

func MapFromModels(bookRows []*BookRow) []*Book {
	books := make([]*Book, 0, len(bookRows))
	for _, bookRow := range bookRows {
		book := &Book{}
		book.MapFromModel(*bookRow)
		books = append(books, book)
	}

	return books
}
