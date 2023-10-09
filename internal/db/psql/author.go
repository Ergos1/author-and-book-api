package psql

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/db"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/models"
)

type AuthorRepo struct {
	db *psql.Database
}

func NewAuthorRepo(db *psql.Database) db.AuthorRepo {
	return &AuthorRepo{
		db: db,
	}
}

func (ar *AuthorRepo) Create(ctx context.Context, authorModel *models.Author) (int64, error) {
	var id int64
	err := ar.db.ExecQueryRow(ctx, "INSERT INTO authors(id, name) VALUES($1, $2) RETURNING id", authorModel.Id, authorModel.Name).Scan(&id)
	if err != nil && err.(*pgconn.PgError).Code == "23505" {
		return id, db.ErrDuplicate
	}

	return id, err
}

func (ar *AuthorRepo) GetById(ctx context.Context, id int64) (*models.Author, error) {
	var author models.Author
	err := ar.db.Get(ctx, &author, "SELECT id, name FROM authors WHERE id=$1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, db.ErrObjectNotFound
	}

	return &author, nil
}

func (ar *AuthorRepo) Update(ctx context.Context, id int64, authorModel *models.Author) error {
	res, err := ar.db.Exec(ctx, "UPDATE authors SET name=$1 WHERE id=$2", authorModel.Name, id)
	if res.RowsAffected() == 0 {
		return db.ErrObjectNotFound
	}

	return err
}

func (ar *AuthorRepo) Delete(ctx context.Context, id int64) error {
	res, err := ar.db.Exec(ctx, "DELETE FROM authors WHERE id=$1", id)
	if res.RowsAffected() == 0 {
		return db.ErrObjectNotFound
	}

	return err
}
