package author

import "context"

func (s *AuthorService) GetByID(ctx context.Context, id int64) (*Author, error) {
	authorRow, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	author := &Author{}
	author.MapFromModel(*authorRow)
	return author, err
}
