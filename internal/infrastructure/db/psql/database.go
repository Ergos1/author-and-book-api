package psql

import (
	"context"
	"errors"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrDatabaseAlreadyClosed = errors.New("database is already closed")
var ErrDatabaseAlreadyConnected = errors.New("database is already connected")

type PGX interface {
	DBops
}

type DBops interface {
	Connect(ctx context.Context, uri string) error
	Close(ctx context.Context) error
	GetPool(_ context.Context) *pgxpool.Pool
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
}

type Database struct {
	cluster *pgxpool.Pool
}

func (db *Database) Connect(ctx context.Context, uri string) error {
	if db.cluster != nil {
		return ErrDatabaseAlreadyConnected
	}

	pool, err := pgxpool.Connect(ctx, uri)
	if err != nil {
		return err
	}

	db.cluster = pool
	return nil
}

func (db *Database) Close(ctx context.Context) error {
	if db.cluster == nil {
		return ErrDatabaseAlreadyClosed
	}

	db.GetPool(ctx).Close()
	return nil
}

func (db *Database) GetPool(_ context.Context) *pgxpool.Pool {
	return db.cluster
}

func (db *Database) SetPool(_ context.Context, pool *pgxpool.Pool) error {
	if db.cluster != nil {
		db.cluster.Close()
	}

	db.cluster = pool
	return nil
}

func (db *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

func (db *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

func (db *Database) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, query, args...)
}

func (db *Database) ExecQueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}
