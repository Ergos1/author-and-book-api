package book

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

type BookRepoPsql struct {
	db PsqlConnector
}

func NewBookRepoPsql(db PsqlConnector) *BookRepoPsql {
	return &BookRepoPsql{
		db: db,
	}
}

func (r *BookRepoPsql) Create(ctx context.Context, bookModel *BookRow) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx, "INSERT INTO books(id, name, rating, author_id) VALUES($1, $2, $3, $4) RETURNING id", bookModel.ID, bookModel.Name, bookModel.Rating, bookModel.AuthorID).Scan(&id)
	if err != nil && err.(*pgconn.PgError).Code == "23505" {
		return id, ErrBookDuplicate
	}

	return id, err
}

func (r *BookRepoPsql) GetById(ctx context.Context, id int64) (*BookRow, error) {
	var book BookRow
	err := r.db.Get(ctx, &book, "SELECT id, name, rating, author_id FROM books WHERE id=$1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrBookNotFound
	}

	return &book, nil
}

func (r *BookRepoPsql) Update(ctx context.Context, id int64, bookModel *BookRow) error {
	res, err := r.db.Exec(ctx, "UPDATE books SET name=$1, rating=$2, author_id=$3 WHERE id=$4", bookModel.Name, bookModel.Rating, bookModel.AuthorID, id)
	if res.RowsAffected() == 0 {
		return ErrBookNotFound
	}

	return err
}

func (r *BookRepoPsql) Delete(ctx context.Context, id int64) error {
	res, err := r.db.Exec(ctx, "DELETE FROM books WHERE id=$1", id)
	if res.RowsAffected() == 0 {
		return ErrBookNotFound
	}

	return err
}

func (r *BookRepoPsql) GetByAuthorId(ctx context.Context, authorID int64) ([]*BookRow, error) {
	var books []*BookRow

	err := r.db.Select(ctx, &books, "SELECT id, name, rating, author_id FROM books WHERE author_id=$1", authorID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
