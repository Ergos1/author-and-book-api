//go:generate mockgen -source=./repository.go -destination=./mocks/repository.go -package=mock_repository
package author

import (
	"context"
	"errors"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
)

type AuthorRepo interface {
	Create(ctx context.Context, authorModel *AuthorRow) (int64, error)
	GetById(ctx context.Context, id int64) (*AuthorRow, error)
	Update(ctx context.Context, id int64, authorModel *AuthorRow) error
	Delete(ctx context.Context, id int64) error
}

type AuthorRepoPsql struct {
	db psql.PGX
}

func NewAuthorRepoPsql(db psql.PGX) *AuthorRepoPsql {
	return &AuthorRepoPsql{
		db: db,
	}
}

func (r *AuthorRepoPsql) Create(ctx context.Context, authorModel *AuthorRow) (int64, error) {
	var id int64
	err := r.db.Create(ctx, &id, "INSERT INTO authors(id, name) VALUES($1, $2) RETURNING id", authorModel.ID, authorModel.Name)
	if err != nil && err.(*pgconn.PgError).Code == "23505" {
		return id, ErrAuthorDuplicate
	}

	return id, err
}

func (r *AuthorRepoPsql) GetById(ctx context.Context, id int64) (*AuthorRow, error) {
	var author AuthorRow
	err := r.db.Get(ctx, &author, "SELECT id, name FROM authors WHERE id=$1", id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAuthorNotFound
		}

		return nil, err
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
