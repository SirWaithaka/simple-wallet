package errors

import "fmt"

// ErrUserHasAccount ...
type ErrUserHasAccount struct {
	message   string
	AccountId string
	UserId    string
}

func (err ErrUserHasAccount) Error() string {
	return fmt.Sprintf("user %v has account with id %v", err.UserId, err.AccountId)
}

// ErrAccountAccess ...
type ErrAccountAccess struct {
	Reason  string
	message string
}


func (err ErrAccountAccess) Error() string {
	msg := fmt.Sprintf("couldn't access account. %v", err.Reason)
	return msg
}

// ErrAmountBelowMinimum
type ErrAmountBelowMinimum struct {
	Message string
}

func (err ErrAmountBelowMinimum) Error() string {
	return err.Message
}

// ErrNotEnoughBalance
type ErrNotEnoughBalance struct {
	Message string
}

func (err ErrNotEnoughBalance) Error() string {
	return err.Message
}
