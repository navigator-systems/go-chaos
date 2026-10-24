package config

import (
	"errors"
	"fmt"
)

// ChaosError represents a custom error with operation context
type ChaosError struct {
	Op  string // Operation that failed
	Err error  // Original error
}

func (e *ChaosError) Error() string {
	if e.Op == "" {
		return e.Err.Error()
	}
	return fmt.Sprintf("%s: %v", e.Op, e.Err)
}

func (e *ChaosError) Unwrap() error {
	return e.Err
}

// NewError creates a new JRXError
func NewError(op string, err error) *ChaosError {
	return &ChaosError{
		Op:  op,
		Err: err,
	}
}

// Common errors
var (
	ErrFileNotFound = errors.New("file not found")
	ErrFileCantRead = errors.New("file cannot be read")
)
