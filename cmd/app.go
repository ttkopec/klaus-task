package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/ttkopec/klaus-task/internal/api"
	"github.com/ttkopec/klaus-task/internal/config"
	"github.com/ttkopec/klaus-task/internal/db"
)

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg := config.NewConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.LogLevel,
	}))
	dbService, err := db.NewService(*cfg, logger)
	if err != nil {
		logger.Error("failed to establish db connection", slog.String("error", err.Error()))
		panic("failed to establish db connection")
	}
	defer dbService.Close()
	srv := api.NewServer(dbService, logger)

	api.Start(ctx, srv, *cfg, logger)
}
