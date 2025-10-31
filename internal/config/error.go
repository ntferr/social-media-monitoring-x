package config

import "fmt"

type ErrorConfig struct {
	Err     error
	Details string
}

func NewErrorConfig(err error, details string) *ErrorConfig {
	return &ErrorConfig{
		Err:     err,
		Details: details,
	}
}

func (e *ErrorConfig) Error() string {
	if e.Err != nil {
		return e.Details
	}
	return fmt.Sprintf("%s: %v", e.Details, e.Err)
}

func (e *ErrorConfig) Unwrap() error {
	return e.Err
}
