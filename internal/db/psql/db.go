package psql

import (
	"context"

	exDB "gitlab.ozon.dev/ergossteam/homework-3/internal/db"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
)

type DB struct {
	database *psql.Database
}

func NewDB(ctx context.Context) *DB {
	return &DB{}
}

func (db *DB) Connect(ctx context.Context, uri string) error {
	psqlDb, err := psql.NewDB(ctx, uri)

	if err != nil {
		return err
	}

	db.database = psqlDb
	return nil
}

func (db *DB) Close(ctx context.Context) error {
	if db.database == nil {
		return exDB.ErrDatabaseAlreadyClosed
	}

	db.database.GetPool(ctx).Close()
	db.database = nil
	return nil
}

func (db *DB) Authors() exDB.AuthorRepo {
	return NewAuthorRepo(db.database)
}

func (db *DB) Books() exDB.BookRepo {
	return NewBookRepo(db.database)
}
