package auth

// ErrTokenParsing ...
type ErrTokenParsing struct {
	message string
}

func (err ErrTokenParsing) Error() string {
	return err.message
}

