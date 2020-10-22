package auth

// TokenParsingError ...
type TokenParsingError struct {
	message string
}

func (err TokenParsingError) Error() string {
	return err.message
}

