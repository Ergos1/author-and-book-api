package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	"gitlab.ozon.dev/ergossteam/homework-3/pkg/commander"
)

type Service interface {
	CreateAuthor(ctx context.Context, request core.CreateAuthorRequest) (int64, error)
}

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}

func getService(ctx context.Context) Service {
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
	return service
}

func run(ctx context.Context) error {
	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	service := getService(ctx)

	cmder := &commander.Commander{
		Commands:  make([]commander.Command, 0),
		OutWriter: os.Stdout,
	}

	cmder.AddCommand(ctx, commander.Command{
		Name:        "create-account",
		Description: "Creates account",
		Handler: func(ctx context.Context, args []string) error {
			authorID, err := cmder.Flags().GetInt64("author_id")
			if err != nil {
				return err
			}

			authorName, err := cmder.Flags().GetString("author_name")
			if err != nil {
				return err
			}

			if authorID == 0 || authorName == "" {
				return errors.New("bad args")
			}

			id, err := service.CreateAuthor(ctx, core.CreateAuthorRequest{
				ID:   authorID,
				Name: authorName,
			})
			if err != nil {
				return err
			}

			fmt.Println(id)

			return nil
		},
	})

	cmder.Flags().Int64("author_id", 0, "author name")
	cmder.Flags().String("author_name", "", "author name")

	if err := cmder.Run(ctx); err != nil {
		return err
	}

	return nil
}
