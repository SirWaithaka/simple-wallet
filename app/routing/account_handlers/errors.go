package account_handlers

import (
	"net/http"
	"reflect"

	"simple-wallet/app/errors"
)

type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ErrResponse(err error) ErrHTTP {

	switch err.(type) {
	case errors.ErrAccountAccess:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusForbidden,
		}

	case errors.ErrNotEnoughBalance:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusForbidden,
		}

	default:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  500,
		}
	}
}
