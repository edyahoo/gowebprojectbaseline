package errors

import "net/http"

type HTTPError struct {
	Code    int
	Message string
}

var (
	ErrBadRequest          = &HTTPError{http.StatusBadRequest, "Bad request"}
	ErrUnauthorized        = &HTTPError{http.StatusUnauthorized, "Unauthorized"}
	ErrForbidden           = &HTTPError{http.StatusForbidden, "Forbidden"}
	ErrNotFound            = &HTTPError{http.StatusNotFound, "Not found"}
	ErrInternalServerError = &HTTPError{http.StatusInternalServerError, "Internal server error"}
)

func ToHTTP(err error) *HTTPError {
	return ErrInternalServerError
}
