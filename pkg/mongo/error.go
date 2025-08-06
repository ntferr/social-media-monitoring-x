package mongo

import "fmt"

type ErrorMongo struct {
	Err     error
	Details string
}

func NewErrorMongo(err error, details string) *ErrorMongo {
	return &ErrorMongo{
		Err:     err,
		Details: details,
	}
}

func (e *ErrorMongo) Error() string {
	if e.Err == nil {
		return e.Details
	}
	return fmt.Sprintf("%s: %v", e.Details, e.Err)
}

func (e *ErrorMongo) Unwrap() error {
	return e.Err
}
