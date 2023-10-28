package psql

import (
	"context"
)

func NewDB(ctx context.Context) *Database {
	return &Database{
		cluster: nil,
	}
}
