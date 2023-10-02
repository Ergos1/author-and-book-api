package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
	"gitlab.ozon.dev/ergossteam/homework-3/internal/transport/http"
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
	var srv Server = http.NewServer(ctx, http.WithAddress(cfg.Server.Address))

	if err := srv.Run(); err != nil {
		log.Printf("[MAIN] Error: %v", err)
	}

	return nil
}
