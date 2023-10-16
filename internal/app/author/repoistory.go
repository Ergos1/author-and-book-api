package author

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type PsqlConnector interface {
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type AuthorRepoPsql struct {
	db PsqlConnector
}

func NewAuthorRepoPsql(db PsqlConnector) *AuthorRepoPsql {
	return &AuthorRepoPsql{
		db: db,
	}
}

func (r *AuthorRepoPsql) Create(ctx context.Context, authorModel *AuthorRow) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO authors(id, name) VALUES($1, $2) RETURNING id", authorModel.ID, authorModel.Name).Scan(&id)
	if err != nil && err.(*pgconn.PgError).Code == "23505" {
		return id, ErrAuthorDuplicate
	}

	return id, err
}

func (r *AuthorRepoPsql) GetById(ctx context.Context, id int64) (*AuthorRow, error) {
	var author AuthorRow
	err := r.db.Get(ctx, &author, "SELECT id, name FROM authors WHERE id=$1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrAuthorNotFound
	}

	return &author, nil
}

func (r *AuthorRepoPsql) Update(ctx context.Context, id int64, authorModel *AuthorRow) error {
	res, err := r.db.Exec(ctx, "UPDATE authors SET name=$1 WHERE id=$2", authorModel.Name, id)
	if res.RowsAffected() == 0 {
		return ErrAuthorNotFound
	}

	return err
}

func (r *AuthorRepoPsql) Delete(ctx context.Context, id int64) error {
	res, err := r.db.Exec(ctx, "DELETE FROM authors WHERE id=$1", id)
	if res.RowsAffected() == 0 {
		return ErrAuthorNotFound
	}

	return err
}
