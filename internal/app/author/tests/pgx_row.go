package tests

import "github.com/jackc/pgconn"

type GoodPgxRow struct{}

func NewGoodPgxRow() *GoodPgxRow {
	return &GoodPgxRow{}
}

func (pr *GoodPgxRow) Scan(dest ...interface{}) error {
	return nil
}

type DuplicatePgxRow struct{}

func NewDuplicatePgxRow() *DuplicatePgxRow {
	return &DuplicatePgxRow{}
}

func (pr *DuplicatePgxRow) Scan(dest ...interface{}) error {
	return &pgconn.PgError{
		Code: "23505",
	}
}
