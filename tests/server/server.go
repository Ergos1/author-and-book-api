package server

import (
	"context"
	"sync"
	"testing"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	myhttp "gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http/handlers"
)

type TestServer struct {
	ctx    context.Context
	cancel context.CancelFunc

	Server  *myhttp.Server
	Service *core.Service
	sync.Mutex
}

func NewTestService(address string, db psql.PGX) *TestServer {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	authorRepo := author.NewAuthorRepoPsql(db)
	authorService := author.NewAuthorService(authorRepo)

	bookRepo := book.NewBookRepoPsql(db)
	bookService := book.NewBookService(bookRepo)

	service := core.NewService(authorService, bookService)

	srv := myhttp.NewServer(ctx,
		myhttp.WithAddress(address),
		myhttp.WithMount("/authors", handlers.NewAuthorHandler(service).Routes()),
		myhttp.WithMount("/service", handlers.NewBookHandler(service).Routes()),
	)

	return &TestServer{
		ctx:     ctx,
		cancel:  cancel,
		Server:  srv,
		Service: service,
	}
}

func (ts *TestServer) SetUp(t *testing.T, args ...interface{}) error {
	t.Helper()
	ts.Lock()
	go ts.Server.Run()

	return nil
}

func (ts *TestServer) TearDown() {
	ts.cancel()
	ts.Unlock()
}
