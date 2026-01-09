package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/go-chi/chi/v5"

	"goprojstructtest/internal/http/routes"
	"goprojstructtest/internal/platform/config"
	"goprojstructtest/internal/platform/httpserver"
	"goprojstructtest/internal/platform/logging"
	"goprojstructtest/internal/render"
)

const LevelSuperDebug slog.Level = -8

func main() {
	cfg := config.Load()

	logger := logging.NewLogger(cfg.LogLevel)
	logger.Info("Configured logger at level", "level", cfg.LogLevel)
	logger.Info("Starting application", "env", cfg.Env)
	logger.Log(
		context.Background(),
		LevelSuperDebug,
		"Loaded configuration",
		"config", cfg,
	)

	renderer, err := render.New()
	if err != nil {
		log.Fatal("Failed to initialize renderer:", err)
	}

	r := chi.NewRouter()

	routes.Setup(r, logger, renderer)

	serverListenAddress := ":" + cfg.ServerAddr
	server := httpserver.New(serverListenAddress, r, logger)
	if err := server.Start(); err != nil {
		logger.Error("Server error", "error", err)
	}
}
