package user

import (
	"fmt"

	"simple-wallet/app/models"
)

// ErrHashPassword
type ErrHashPassword struct {
	password string
	message  string
}

func (err ErrHashPassword) Error() string {
	return err.message
}

// ErrUnauthorized ...
type ErrUnauthorized struct {
	Message string
}

func (err ErrUnauthorized) Error() string {
	return err.Message
}

// ErrUserNotFound ...
type ErrUserNotFound struct {
	message string
}

func (err ErrUserNotFound) Error() string {
	return err.message
}

// ErrUserExists returned when adding a user with
// phone number or email number that are already in the db.
type ErrUserExists struct {
	message string
	inUser  models.User
	outUser models.User
}

func (err *ErrUserExists) Error() string {
	err.message = "user exists"

	if err.outUser.Email == err.inUser.Email {
		err.message = fmt.Sprintf("user with email %s exists", err.inUser.Email)
	} else if err.outUser.PhoneNumber == err.inUser.PhoneNumber {
		err.message = fmt.Sprintf("user with phone number %s exists", err.inUser.PhoneNumber)
	}

	return err.message
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
