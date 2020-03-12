package account

import (
	"fmt"
)

// ErrUserHasAccount ...
type ErrUserHasAccount struct {
	message   string
	accountId string
	userId    string
}

func (err ErrUserHasAccount) Error() string {
	return fmt.Sprintf("user %v has account with id %v", err.userId, err.accountId)
}

// ErrAccountAccess ...
type ErrAccountAccess struct {
	reason  string
	message string
}


func (err ErrAccountAccess) Error() string {
	msg := fmt.Sprintf("couldn't access account. %v", err.reason)
	return msg
}

// ErrAmountBelowMinimum
type ErrAmountBelowMinimum struct {
	Message string
}

func (err ErrAmountBelowMinimum) Error() string {
	return err.Message
}

// NewErrUnexpected create a new ErrUnexpected error for when doing
// queries or anything else for the type user.
func NewErrUnexpected(err error) *ErrUnexpected {
	return &ErrUnexpected{
		debug: err.Error(),
	}
}

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
