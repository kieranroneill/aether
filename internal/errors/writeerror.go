package errors

import "aether/internal/constants"

type WriteError struct {
	Code    int
	Error   error
	Message string
}

func NewWriteError(message string, error error) *WriteError {
	return &WriteError{
		Code:    constants.WriteError,
		Error:   error,
		Message: message,
	}
}
