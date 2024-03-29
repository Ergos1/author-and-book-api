package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	author_pb "gitlab.ozon.dev/ergossteam/homework-3/pkg/api/grpc/v1/author"
	"gitlab.ozon.dev/ergossteam/homework-3/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ErrNotImplementedCommand = errors.New("not implemented command")
	ErrIdNotGiven            = errors.New("id not given")
	ErrNameNotGiven          = errors.New("name not given")
)

type CMD int64

const (
	GetById CMD = iota
	Create
	Update
	Delete
	Default
)

func (cmd *CMD) String() string {
	return fmt.Sprintf("%d", *cmd)
}

func (cmd *CMD) Set(value string) error {
	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return err
	}
	*cmd = CMD(i)
	return nil
}

type Meta struct {
	id   int64
	name string
}

var meta Meta

func main() {
	cmd := Default
	flag.Var(&cmd, "cmd", "The command to run (0=GetById, 1=Create, 2=Update, 3=Delete)")
	flag.Int64Var(&meta.id, "id", -1, "author id for commands")
	flag.StringVar(&meta.name, "name", "", "author name for commands")
	flag.Parse()

	ctx := context.Background()
	if err := run(ctx, cmd); err != nil {
		logger.Errorf(ctx, "error: %v", err)
	}
}

func run(ctx context.Context, cmd CMD) error {
	ctx, cancels := signal.NotifyContext(ctx, os.Interrupt)
	defer cancels()

	zapLogger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	zapLogger = zapLogger.With(zap.String("component", "grpc_client"))
	logger.SetGlobal(zapLogger)

	cfg := config.NewConfig()
	conn, err := grpc.Dial(cfg.Server.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	client := author_pb.NewAuthorServiceClient(conn)

	switch cmd {
	case GetById:
		return CMDGetByID(ctx, client)
	case Create:
		return CMDCreate(ctx, client)
	case Delete:
		return CMDDelete(ctx, client)
	case Update:
		return CMDUpdate(ctx, client)
	default:
		return ErrNotImplementedCommand
	}
}

func CMDGetByID(ctx context.Context, client author_pb.AuthorServiceClient) error {
	if meta.id == -1 {
		return ErrIdNotGiven
	}

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "get author by id")))

	logger.Infof(ctx, "getting author by id: %v\n", meta.id)
	resp, err := client.GetByID(ctx, &author_pb.GetByIDRequest{Id: meta.id})
	if err != nil {
		return err
	}

	logger.Infof(ctx, "response: %s\n", resp)
	return nil
}

func CMDCreate(ctx context.Context, client author_pb.AuthorServiceClient) error {
	if meta.id == -1 {
		return ErrIdNotGiven
	}

	if meta.name == "" {
		return ErrNameNotGiven
	}

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "create author with id")))

	logger.Infof(ctx, "creating author with id: %v; name: %v\n", meta.id, meta.name)
	resp, err := client.Create(ctx, &author_pb.CreateRequest{Id: meta.id, Name: meta.name})
	if err != nil {
		return err
	}

	logger.Infof(ctx, "response: %s\n", resp)
	return nil
}

func CMDUpdate(ctx context.Context, client author_pb.AuthorServiceClient) error {
	if meta.id == -1 {
		return ErrIdNotGiven
	}

	if meta.name == "" {
		return ErrNameNotGiven
	}

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "update author by id")))

	logger.Infof(ctx, "updating author with id: %v; name: %v\n", meta.id, meta.name)
	resp, err := client.Update(ctx, &author_pb.UpdateRequest{Id: meta.id, Name: meta.name})
	if err != nil {
		return err
	}

	logger.Infof(ctx, "response: %s\n", resp)
	return nil
}

func CMDDelete(ctx context.Context, client author_pb.AuthorServiceClient) error {
	if meta.id == -1 {
		return ErrIdNotGiven
	}

	l := logger.FromContext(ctx)
	ctx = logger.ToContext(ctx, l.With(zap.String("method", "delete author by id")))

	logger.Infof(ctx, "deleting author by id: %v\n", meta.id)
	resp, err := client.Delete(ctx, &author_pb.DeleteRequest{Id: meta.id})
	if err != nil {
		return err
	}

	logger.Infof(ctx, "response: %s\n", resp)
	return nil
}
