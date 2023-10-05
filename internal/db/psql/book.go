package psql

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/db"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/models"
	"gitlab.ozon.dev/ergossteam/homework-3/pkg/db/psql"
)

type BookRepo struct {
	db *psql.Database
}

func NewBookRepo(db *psql.Database) db.BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (br *BookRepo) Create(ctx context.Context, authorModel *models.Book) error {
	return nil
}

func (br *BookRepo) GetById(ctx context.Context, id int64) (*models.Book, error) {
	return nil, nil
}

func (br *BookRepo) Update(ctx context.Context, id int64, authorModel *models.Book) error {
	return nil
}

func (br *BookRepo) Delete(ctx context.Context, id int64) error {
	return nil
}
