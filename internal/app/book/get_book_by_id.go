package book

import "context"

func (s *BookService) GetByID(ctx context.Context, id int64) (*Book, error) {
	bookRow, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	book := &Book{}
	book.MapFromModel(*bookRow)
	return book, err
}
