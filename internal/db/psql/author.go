package psql

import (
	"context"
	"database/sql"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/db"
	pkg_psql "gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/models"
)

type AuthorRepo struct {
	db *pkg_psql.Database
}

func NewAuthorRepo(db *pkg_psql.Database) db.AuthorRepo {
	return &AuthorRepo{
		db: db,
	}
}

func (ar *AuthorRepo) Create(ctx context.Context, authorModel *models.Author) (int64, error) {
	var id int64
	err := ar.db.ExecQueryRow(ctx, "INSERT INTO authors(name) VALUES($1) RETURNING id", authorModel.Name).Scan(&id)
	return id, err
}

func (ar *AuthorRepo) GetById(ctx context.Context, id int64) (*models.Author, error) {
	var author models.Author
	err := ar.db.Get(ctx, &author, "SELECT id, name FROM authors WHERE id=$1", id)
	if err == sql.ErrNoRows {
		return nil, ErrObjectNotFound
	}

	return &author, nil
}

func (ar *AuthorRepo) Update(ctx context.Context, id int64, authorModel *models.Author) error {
	_, err := ar.db.Exec(ctx, "UPDATE authors SET name=$1 WHERE id=$2", authorModel.Name, id)
	return err
}

func (ar *AuthorRepo) Delete(ctx context.Context, id int64) error {
	_, err := ar.db.Exec(ctx, "DELETE FROM authors WHERE id=$1", id)
	return err
}
