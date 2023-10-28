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
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/kafka"
	kafkarequestlogger "gitlab.ozon.dev/ergossteam/homework-3/internal/kafka_request_logger"
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
		log.Fatal(err)
	}
	defer db.Close(ctx)

	authorRepo := author.NewAuthorRepoPsql(db)
	authorService := author.NewAuthorService(authorRepo)

	bookRepo := book.NewBookRepoPsql(db)
	bookService := book.NewBookService(bookRepo)

	service := core.NewService(authorService, bookService)

	brokers := []string{
		"localhost:9091",
		"localhost:9092",
		"localhost:9093",
	}

	producer, err := kafka.NewProducer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := kafka.NewConsumer(brokers)
	if err != nil {
		log.Fatal(err)
	}

	requestLogger := kafkarequestlogger.NewKafkaRequestLogger(producer)

	var srv Server = http.NewServer(ctx,
		http.WithAddress(cfg.Server.Address),
		http.WithMount("/", handlers.NewBaseHandler().Routes()),
		http.WithMount("/authors", handlers.NewAuthorHandler(service, requestLogger).Routes()),
		http.WithMount("/books", handlers.NewBookHandler(service).Routes()),
		http.WithKafkaProducer(producer),
		http.WithKafkaConsumer(consumer),
	)

	if err := srv.Run(); err != nil {
		log.Printf("[MAIN] Error while running server: %v", err)
	}

	return nil
}
