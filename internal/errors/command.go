package errs

import (
	"errors"
	"fmt"
)

var (
	ErrCollectionFileRequired = errors.New("collection file is required (use -c or --collection)")
	ErrRunsOrDurationRequired = errors.New("one or both of --runs or --duration must be specified and greater than 0")
	ErrConcurrentWorkers      = errors.New("concurrent workers must be greater than 0")
)

func ErrInvalidLogLevel(s string) error {
	return fmt.Errorf("invalid log level '%s', must be one of: debug, info, warn, error", s)
}

func ErrInvalidFormat(f string) error {
	return fmt.Errorf("invalid format '%s' - must be one of: html, json, csv", f)
}

func ErrValidationFailed(err error) error {
	return fmt.Errorf("validation failed: %w", err)
}
