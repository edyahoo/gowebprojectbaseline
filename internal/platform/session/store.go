package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"sync"
	"time"

	"goprojstructtest/internal/domain"
)

var (
	ErrSessionNotFound = errors.New("session not found")
	ErrSessionExpired  = errors.New("session expired")
)

// SessionData holds the server-side session information
type SessionData struct {
	UserID    domain.UserID
	Role      domain.Role
	TenantID  domain.TenantID
	ExpiresAt time.Time
}

// SessionStore defines the interface for session storage
type SessionStore interface {
	Create(userID domain.UserID, role domain.Role, tenantID domain.TenantID, duration time.Duration) (string, error)
	Get(sessionID string) (*SessionData, error)
	Delete(sessionID string) error
}

// InMemoryStore implements SessionStore using an in-memory map
type InMemoryStore struct {
	sessions map[string]*SessionData
	mu       sync.RWMutex
}

// NewInMemoryStore creates a new in-memory session store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		sessions: make(map[string]*SessionData),
	}
}

// Create generates a new session with a cryptographically secure session ID
func (s *InMemoryStore) Create(userID domain.UserID, role domain.Role, tenantID domain.TenantID, duration time.Duration) (string, error) {
	sessionID, err := generateSessionID()
	if err != nil {
		return "", err
	}

	s.mu.Lock()
	s.sessions[sessionID] = &SessionData{
		UserID:    userID,
		Role:      role,
		TenantID:  tenantID,
		ExpiresAt: time.Now().Add(duration),
	}
	s.mu.Unlock()

	return sessionID, nil
}

// Get retrieves session data by session ID, returns error if not found or expired
func (s *InMemoryStore) Get(sessionID string) (*SessionData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, exists := s.sessions[sessionID]
	if !exists {
		return nil, ErrSessionNotFound
	}

	if time.Now().After(data.ExpiresAt) {
		return nil, ErrSessionExpired
	}

	return data, nil
}

// Delete removes a session from the store
func (s *InMemoryStore) Delete(sessionID string) error {
	s.mu.Lock()
	delete(s.sessions, sessionID)
	s.mu.Unlock()
	return nil
}

// CleanupExpired removes all expired sessions from the store
func (s *InMemoryStore) CleanupExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	for id, data := range s.sessions {
		if now.After(data.ExpiresAt) {
			delete(s.sessions, id)
		}
	}
}

// generateSessionID creates a cryptographically secure random session ID
func generateSessionID() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
