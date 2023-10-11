package author

import "context"

type AuthorRepo interface {
	Create(ctx context.Context, authorModel *AuthorRow) (int64, error)
	GetById(ctx context.Context, id int64) (*AuthorRow, error)
	Update(ctx context.Context, id int64, authorModel *AuthorRow) error
	Delete(ctx context.Context, id int64) error
}

type AuthorService struct {
	repo AuthorRepo
}

func NewAuthorService(repo AuthorRepo) *AuthorService {
	return &AuthorService{
		repo: repo,
	}
}
