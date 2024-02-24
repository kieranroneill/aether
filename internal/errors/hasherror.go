package errors

import "aether/internal/constants"

type HashError struct {
	Code    int
	Error   error
	Message string
}

func NewHashError(message string, error error) *HashError {
	return &HashError{
		Code:    constants.HashError,
		Error:   error,
		Message: message,
	}
}
