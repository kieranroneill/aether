package errors

import "aether/internal/constants"

type ReadError struct {
	Code    int
	Error   error
	Message string
}

func NewReadError(message string, error error) *ReadError {
	return &ReadError{
		Code:    constants.ReadError,
		Error:   error,
		Message: message,
	}
}
