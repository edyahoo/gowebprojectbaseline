package appctx

import (
	"context"
	"log/slog"

	"goprojstructtest/internal/domain"
)

type contextKey string

const (
	userKey   contextKey = "user"
	tenantKey contextKey = "tenant"
	loggerKey contextKey = "logger"
)

func WithUser(ctx context.Context, user *domain.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *domain.User {
	if user, ok := ctx.Value(userKey).(*domain.User); ok {
		return user
	}
	return nil
}

func WithTenant(ctx context.Context, tenantID domain.TenantID) context.Context {
	return context.WithValue(ctx, tenantKey, tenantID)
}

func TenantID(ctx context.Context) domain.TenantID {
	if id, ok := ctx.Value(tenantKey).(domain.TenantID); ok {
		return id
	}
	return 0
}

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func Logger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
