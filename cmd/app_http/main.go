package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
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

	db, err := psql.NewDB(ctx, cfg.Database.Uri())
	if err != nil {
		log.Fatalf("[MAIN] Error while connecting db: %v", err)
	}

	defer func() {
		err := db.Close(ctx)
		if err != nil {
			log.Printf("[MAIN] Error while closing db: %v", err)
		}
	}()

	authorRepo := author.NewAuthorRepoPsql(db)
	authorService := author.NewAuthorService(authorRepo)

	bookRepo := book.NewBookRepoPsql(db)
	bookService := book.NewBookService(bookRepo)

	service := core.NewService(authorService, bookService)

	var srv Server = http.NewServer(ctx,
		http.WithAddress(cfg.Server.Address),
		http.WithMount("/", handlers.NewBaseHandler().Routes()),
		http.WithMount("/authors", handlers.NewAuthorHandler(service).Routes()),
		http.WithMount("/books", handlers.NewBookHandler(service).Routes()),
	)

	if err := srv.Run(); err != nil {
		log.Printf("[MAIN] Error while running server: %v", err)
	}

	return nil
}
