package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"gitlab.ozon.dev/ergossteam/homework-3/internal/config"
)

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
	fmt.Println(cfg)

	return nil
}
