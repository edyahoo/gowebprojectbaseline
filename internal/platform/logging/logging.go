package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const LevelSuperDebug slog.Level = -8

func NewLogger(level string) *slog.Logger {
	var logLevel slog.Level
	fmt.Printf("Setting App Log Level to: %s\n", level)
	switch strings.ToLower(level) {
	case "superdebug", "trace":
		logLevel = LevelSuperDebug
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	levelVar := new(slog.LevelVar)
	levelVar.Set(logLevel)

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     levelVar,
		AddSource: true, // ← file:line included
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Intercept source attribute
			if a.Key == slog.SourceKey {
				if src, ok := a.Value.Any().(*slog.Source); ok {
					src.File = filepath.Base(src.File) // keep filename only
					return slog.Any(a.Key, src)
				}
			}
			return a
		},
	})
	/*
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})
	*/
	return slog.New(handler)
}
