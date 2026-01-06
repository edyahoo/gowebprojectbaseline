package domain

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserExists         = errors.New("user already exists")
	ErrInvalidEmail       = errors.New("invalid email")
	ErrTenantNotFound     = errors.New("tenant not found")
	ErrInvalidTenant      = errors.New("invalid tenant")
)
