package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/opentracing/opentracing-go"
	jaeger_config "github.com/uber/jaeger-client-go/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/author"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/book"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/app/core"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/infrastructure/db/psql"
	impl "gitlab.ozon.dev/ergossteam/homework-3/internal/transport/grpc"
	author_pb "gitlab.ozon.dev/ergossteam/homework-3/pkg/api/grpc/v1/author"
	"gitlab.ozon.dev/ergossteam/homework-3/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Print(err)
		// logger.Errorf(ctx, "%v", err)
	}
}

func run(ctx context.Context) error {
	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zapLogger = zapLogger.With(zap.String("component", "grpc_server"))
	logger.SetGlobal(zapLogger)

	jaeger_cfg := jaeger_config.Configuration{
		ServiceName: "GRPC_SERVER",
		Sampler: &jaeger_config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaeger_config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
		},
	}
	tracer, closer, err := jaeger_cfg.NewTracer()
	if err != nil {
		return err
	}
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	cfg := config.NewConfig()
	db := psql.NewDB(ctx)
	if err := db.Connect(ctx, cfg.Database.Uri()); err != nil {
		logger.Errorf(ctx, "%v", err)
	}
	defer db.Close(ctx)

	authorRepo := author.NewAuthorRepoPsql(db)
	authorService := author.NewAuthorService(authorRepo)

	bookRepo := book.NewBookRepoPsql(db)
	bookService := book.NewBookService(bookRepo)

	service := core.NewService(authorService, bookService)

	server := grpc.NewServer()
	author_pb.RegisterAuthorServiceServer(server, impl.New(service))

	lis, err := net.Listen("tcp", cfg.Server.Address)
	if err != nil {
		return err
	}

	go func() {
		<-ctx.Done()
		server.GracefulStop()
	}()

	return server.Serve(lis)
}
