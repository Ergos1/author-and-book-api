package tests

import (
	"context"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/tests/postgres"
	"gitlab.ozon.dev/ergossteam/homework-3/tests/server"
)

var (
	db  *postgres.TDB
	srv *server.TestServer
)

func init() {
	cfg := config.NewConfig()
	db = postgres.NewTestDB(context.Background(), cfg.Database.Uri())
	srv = server.NewTestService(cfg.Server.Address, db.DB)
}
