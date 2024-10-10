package models

import "fmt"

// AppError represents a custom error type for the application
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// NewAppError creates a new AppError
func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Predefined error codes
const (
	ErrCodeBadRequest          = 400
	ErrCodeUnauthorized        = 401
	ErrCodeForbidden           = 403
	ErrCodeNotFound            = 404
	ErrCodeInternalServerError = 500
)

// Common errors
var (
	ErrBadRequest          = NewAppError(ErrCodeBadRequest, "Invalid request")
	ErrUnauthorized        = NewAppError(ErrCodeUnauthorized, "Unauthorized access")
	ErrForbidden           = NewAppError(ErrCodeForbidden, "Access forbidden")
	ErrNotFound            = NewAppError(ErrCodeNotFound, "Resource not found")
	ErrInternalServerError = NewAppError(ErrCodeInternalServerError, "Internal server error")
)

// Custom error creation functions
func ErrInvalidInput(field string) *AppError {
	return NewAppError(ErrCodeBadRequest, fmt.Sprintf("Invalid input for field: %s", field))
}

func ErrDuplicateEntry(field string) *AppError {
	return NewAppError(ErrCodeBadRequest, fmt.Sprintf("Duplicate entry for field: %s", field))
}

func ErrResourceNotFound(resource string) *AppError {
	return NewAppError(ErrCodeNotFound, fmt.Sprintf("%s not found", resource))
}
