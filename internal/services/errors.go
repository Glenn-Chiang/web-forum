package services

import "fmt"

type ValidationError struct {
	Message string
}

func (err *ValidationError) Error() string {
	return err.Message
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

type AlreadyInUseError struct {
	Field string
}

func (err *AlreadyInUseError) Error() string {
	return fmt.Sprintf("%s already in use", err.Field)
}

func NewAlreadyInUseError(field string) *AlreadyInUseError {
	return &AlreadyInUseError{Field: field}
}
