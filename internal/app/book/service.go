package book

import "context"

type BookRepo interface {
	Create(ctx context.Context, authorModel *BookRow) (int64, error)
	GetById(ctx context.Context, id int64) (*BookRow, error)
	Update(ctx context.Context, id int64, authorModel *BookRow) error
	Delete(ctx context.Context, id int64) error
}

type BookService struct {
	repo BookRepo
}

func NewBookService(repo BookRepo) *BookService {
	return &BookService{
		repo: repo,
	}
}
