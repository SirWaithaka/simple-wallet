package transaction

import "fmt"

// ErrUnexpected ...
type ErrUnexpected struct {
	debug string
}

// Error returns a custom error that we can show to user.
func (err *ErrUnexpected) Error() string {
	return fmt.Sprintf("unexpected error occurred")
}

// Debug returns the original error that we can log.
func (err *ErrUnexpected) Debug() string {
	return err.debug
}

// NewErrUnexpected create a new ErrUnexpected error for when doing
// queries or anything else for the type user.
func NewErrUnexpected(err error) *ErrUnexpected {
	return &ErrUnexpected{
		debug: err.Error(),
	}
}
