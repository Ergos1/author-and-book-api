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

type BookRepo struct {
	db *psql.Database
}

func NewBookRepo(db *psql.Database) db.BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (br *BookRepo) Create(ctx context.Context, bookModel *models.Book) (int64, error) {
	var id int64
	err := br.db.ExecQueryRow(ctx, "INSERT INTO books(id, name, rating, author_id) VALUES($1, $2, $3, $4) RETURNING id", bookModel.Id, bookModel.Name, bookModel.Rating, bookModel.AuthorID).Scan(&id)
	if err != nil && err.(*pgconn.PgError).Code == "23505" {
		return id, db.ErrDuplicate
	}

	return id, err
}

func (br *BookRepo) GetById(ctx context.Context, id int64) (*models.Book, error) {
	var book models.Book
	err := br.db.Get(ctx, &book, "SELECT id, name, rating, author_id FROM books WHERE id=$1", id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, db.ErrObjectNotFound
	}

	return &book, nil
}

func (br *BookRepo) Update(ctx context.Context, id int64, bookModel *models.Book) error {
	res, err := br.db.Exec(ctx, "UPDATE books SET name=$1, rating=$2, author_id=$3 WHERE id=$4", bookModel.Name, bookModel.Rating, bookModel.AuthorID, id)
	if res.RowsAffected() == 0 {
		return db.ErrObjectNotFound
	}

	return err
}

func (br *BookRepo) Delete(ctx context.Context, id int64) error {
	res, err := br.db.Exec(ctx, "DELETE FROM books WHERE id=$1", id)
	if res.RowsAffected() == 0 {
		return db.ErrObjectNotFound
	}

	return err
}

func (br *BookRepo) GetByAuthorId(ctx context.Context, authorID int64) ([]*models.Book, error) {
	var books []*models.Book

	err := br.db.Select(ctx, &books, "SELECT id, name, rating, author_id FROM books WHERE author_id=$1", authorID)
	if err != nil {
		return nil, err
	}

	return books, nil
}
