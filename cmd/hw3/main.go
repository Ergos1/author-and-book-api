package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/db/psql"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http/handlers"
)

type Server interface {
	Run() error
}

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	cfg := config.NewConfig()

	db := psql.NewDB(ctx)
	if err := db.Connect(ctx, cfg.Database.Uri()); err != nil {
		log.Fatalf("[MAIN] Error while connecting db: %v", err)
	}
	defer func() {
		err := db.Close(ctx)
		if err != nil {
			log.Printf("[MAIN] Error while closing db: %v", err)
		}
	}()

	var srv Server = http.NewServer(ctx,
		http.WithAddress(cfg.Server.Address),
		http.WithMount("/", handlers.NewBaseHandler().Routes()),
		http.WithMount("/authors", handlers.NewAuthorHandler(db).Routes()),
		http.WithMount("/books", handlers.NewBookHandler(db).Routes()),
	)

	if err := srv.Run(); err != nil {
		log.Printf("[MAIN] Error while running server: %v", err)
	}

	return nil
}
