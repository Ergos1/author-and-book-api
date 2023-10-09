package db

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/models"
)

type DB interface {
	Connect(ctx context.Context, uri string) error
	Close(ctx context.Context) error

	Authors() AuthorRepo
	Books() BookRepo
}

type AuthorRepo interface {
	Create(ctx context.Context, authorModel *models.Author) (int64, error)
	GetById(ctx context.Context, id int64) (*models.Author, error)
	Update(ctx context.Context, id int64, authorModel *models.Author) error
	Delete(ctx context.Context, id int64) error
}

type BookRepo interface {
	Create(ctx context.Context, bookModel *models.Book) (int64, error)
	GetById(ctx context.Context, id int64) (*models.Book, error)
	Update(ctx context.Context, id int64, bookModel *models.Book) error
	Delete(ctx context.Context, id int64) error
	GetByAuthorId(ctx context.Context, authorID int64) ([]*models.Book, error)
}
