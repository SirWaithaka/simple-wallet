package errors

// ErrUserNotFound
type ErrUserNotFound string

func (e ErrUserNotFound) Error() string {
	return string(e)
}

// ErrUserExists returned when adding a user with
// phone number or email number that are already in the db.
type ErrUserExists struct {
	Err error
}

func (e ErrUserExists) Error() string {
	return "user already exists"
}

// ErrInvalidCredentials
type ErrInvalidCredentials ErrorT

func (e ErrInvalidCredentials) Error() string {
	return string(ErrorInvalidUsernameOrPassword)
}

func (e ErrInvalidCredentials) Debug() error {
	return e.Err
}

// PasswordHashError
type PasswordHashError ErrorT

func (e PasswordHashError) Error() string {
	return "error hashing password"
}

func (e PasswordHashError) Debug() error {
	return e.Err
}
