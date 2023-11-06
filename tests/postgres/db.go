package postgres

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"testing"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
)

type TDB struct {
	uri string
	ctx context.Context

	DB psql.PGX
	sync.Mutex
}

func NewTestDB(ctx context.Context, uri string) *TDB {
	db := psql.NewDB(ctx)
	return &TDB{
		ctx: ctx,
		uri: uri,
		DB:  db,
	}
}

func (d *TDB) SetUp(t *testing.T, args ...interface{}) error {
	t.Helper()
	d.Lock()
	err := d.DB.Connect(d.ctx, d.uri)
	if err != nil {
		return err
	}

	d.Truncate()
	return nil
}

func (d *TDB) TearDown() {
	d.Truncate()

	d.DB.Close(d.ctx)
	defer d.Unlock()
}

func (d *TDB) Truncate() {
	var tables []string
	err := d.DB.Select(d.ctx, &tables, "SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_type='BASE TABLE' AND table_name != 'goose_db_version'")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	if len(tables) == 0 {
		panic("run migration plz")
	}
	q := fmt.Sprintf("Truncate table %s", strings.Join(tables, ","))
	if _, err := d.DB.Exec(d.ctx, q); err != nil {
		panic(err)
	}
}
