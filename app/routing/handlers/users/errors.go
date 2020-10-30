package users

// ErrInvalidParams ...
type ErrInvalidParams struct {
	message string
}

func (err ErrInvalidParams) Error() string {
	return err.message
}

type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
