package main

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"

	"goprojstructtest/internal/http/routes"
	"goprojstructtest/internal/platform/config"
	"goprojstructtest/internal/platform/httpserver"
	"goprojstructtest/internal/platform/logging"
	"goprojstructtest/internal/platform/session"
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

	// Create session store
	sessionStore := session.NewInMemoryStore()

	// Start session cleanup goroutine
	go func() {
		ticker := time.NewTicker(15 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			sessionStore.CleanupExpired()
			logger.Debug("Cleaned up expired sessions")
		}
	}()

	r := chi.NewRouter()

	routes.Setup(r, logger, renderer, sessionStore, cfg)

	serverListenAddress := ":" + cfg.ServerAddr
	server := httpserver.New(serverListenAddress, r, logger)
	if err := server.Start(); err != nil {
		logger.Error("Server error", "error", err)
	}
}
