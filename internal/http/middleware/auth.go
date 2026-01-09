package middleware

import (
	"log/slog"
	"net/http"

	"goprojstructtest/internal/appctx"
	"goprojstructtest/internal/domain"
	"goprojstructtest/internal/platform/session"
)

// RequireAuth validates the session and adds user to context
func RequireAuth(store session.SessionStore, logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				logger.Warn("No session cookie found")
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			sessionData, err := store.Get(cookie.Value)
			if err != nil {
				logger.Warn("Invalid or expired session", "error", err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			user := &domain.User{
				ID:       sessionData.UserID,
				TenantID: sessionData.TenantID,
				Role:     sessionData.Role,
			}

			ctx := appctx.WithUser(r.Context(), user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// RequireRole checks that the authenticated user has the required role
func RequireRole(requiredRole domain.Role, logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := appctx.User(r.Context())
			if user == nil {
				logger.Warn("No user in context for role check")
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			if user.Role != requiredRole {
				logger.Warn("User does not have required role",
					"userRole", user.Role,
					"requiredRole", requiredRole)
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
