package users

import (
	"net/http"
	"reflect"

	"simple-wallet/app/auth"
	"simple-wallet/app/user"
)

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

func ErrResponse(err error) ErrHTTP {
	var errHTTP ErrHTTP
	switch err.(type) {

	case ErrInvalidParams:
		errHTTP = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

	case *user.ErrUserExists:
		var e = err.(*user.ErrUserExists)
		errHTTP = ErrHTTP{
			Error:   reflect.TypeOf(*e).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

	case user.ErrUserNotFound:
		errHTTP = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

	case auth.ErrTokenParsing:
		errHTTP = ErrHTTP{
			Error:   "ErrServerError",
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}

	case user.ErrUnauthorized:
		errHTTP = ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}

	default:
		errHTTP = ErrHTTP{
			Error:   "ErrUnknown",
			Message: "something unfortunate happened!",
			Status:  http.StatusInternalServerError,
		}
	}

	return errHTTP
}
