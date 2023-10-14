package psql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context, uri string) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		return nil, err
	}

	database := newDatabase(pool)
	go database.CloseGracefully(ctx)

	return database, nil
}
