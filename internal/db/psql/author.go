package psql

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/db"
	pkg_psql "gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/models"
)

type AuthorRepo struct {
	psqlDB *pkg_psql.Database
}

func NewAuthorRepo(db *pkg_psql.Database) db.AuthorRepo {
	return &AuthorRepo{
		psqlDB: db,
	}
}

func (ar *AuthorRepo) Create(ctx context.Context, authorModel *models.Author) error {
	return nil
}

func (ar *AuthorRepo) GetById(ctx context.Context, id int64) (*models.Author, error) {
	return nil, nil
}

func (ar *AuthorRepo) Update(ctx context.Context, id int64, authorModel *models.Author) error {
	return nil
}

func (ar *AuthorRepo) Delete(ctx context.Context, id int64) error {
	return nil
}
