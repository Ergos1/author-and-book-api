package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	"gitlab.ozon.dev/ergossteam/homework-3/pkg/commander"
)

type Service interface {
	CreateAuthor(ctx context.Context, request core.CreateAuthorRequest) (int64, error)
	GetAuthorById(ctx context.Context, id int64) (*author.Author, error)
}

var ErrBadArgs = errors.New("bad args")
var ErrBadFilePath = errors.New("bad file path")
var ErrCannotCreateFile = errors.New("cannot create file")

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}

func getService(ctx context.Context, db psql.PGX) Service {
	authorRepo := author.NewAuthorRepoPsql(db)
	authorService := author.NewAuthorService(authorRepo)

	bookRepo := book.NewBookRepoPsql(db)
	bookService := book.NewBookService(bookRepo)

	service := core.NewService(authorService, bookService)
	return service
}

func run(ctx context.Context) error {
	log.SetPrefix("main")

	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	cfg := config.NewConfig()

	db := psql.NewDB(ctx)
	if err := db.Connect(ctx, cfg.Database.Uri()); err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	service := getService(ctx, db)

	cmder := &commander.Commander{
		Commands:  make([]commander.Command, 0),
		OutWriter: os.Stdout,
	}

	cmder.AddCommand(ctx, commander.Command{
		Name:        "create-author",
		Description: "Creates author",
		Flags: []commander.Flag{
			{
				Name:        "author_id",
				Description: "id of author",
			},
			{
				Name:        "author_name",
				Description: "name of author",
			},
		},
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
				return ErrBadArgs
			}

			id, err := service.CreateAuthor(ctx, core.CreateAuthorRequest{
				ID:   authorID,
				Name: authorName,
			})
			if err != nil {
				return err
			}

			fmt.Printf("Created account id = %v", id)

			return nil
		},
	})

	cmder.AddCommand(ctx, commander.Command{
		Name:        "get-author",
		Description: "Get author by id",
		Flags: []commander.Flag{
			{
				Name:        "author_id",
				Description: "id of author",
			},
		},
		Handler: func(ctx context.Context, args []string) error {
			authorID, err := cmder.Flags().GetInt64("author_id")
			if err != nil {
				return err
			}

			if authorID == 0 {
				return ErrBadArgs
			}

			author, err := service.GetAuthorById(ctx, authorID)
			if err != nil {
				return err
			}

			fmt.Printf("Author = %v", *author)
			return nil
		},
	})

	cmder.AddCommand(ctx, commander.Command{
		Name:        "gofmt",
		Description: "format file",
		Flags: []commander.Flag{
			{
				Name:        "file_path",
				Description: "relative path of file",
			},
		},
		Handler: func(ctx context.Context, args []string) error {
			filePath, err := cmder.Flags().GetString("file_path")
			if err != nil {
				return err
			}

			if filePath == "" {
				return ErrBadArgs
			}

			fileContent, err := os.ReadFile(filePath)
			if err != nil {
				return ErrBadFilePath
			}

			writeFile, err := os.Create("output.txt")
			if err != nil {
				return ErrCannotCreateFile
			}

			lines := make([]string, 0)

			scanner := bufio.NewScanner(strings.NewReader(string(fileContent)))
			for scanner.Scan() {
				line := scanner.Text()
				lines = append(lines, line)
			}

			isStart := true
			formatedLines := make([]string, 0, len(lines))

			for index, line := range lines {
				if isStart {
					line = "\t" + line
					isStart = false
				}

				if line == "" {
					isStart = true
					formatedLines = append(formatedLines, "")
					continue
				}

				if (index == len(lines)-1 || lines[index+1] == "") && string(line[len(line)-1]) != "." {
					line = line + "."
				}
				formatedLines = append(formatedLines, line)
			}

			for _, line := range formatedLines {
				writeFile.WriteString(line + "\n")
			}

			return nil
		},
	})

	cmder.Flags().Int64("author_id", 0, "author id")
	cmder.Flags().String("author_name", "", "author name")
	cmder.Flags().String("file_path", "", "relative path of file")

	if err := cmder.Run(ctx); err != nil {
		return err
	}

	return nil
}
