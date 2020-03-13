package accounts

import (
	"net/http"
	"reflect"
	"wallet/account"
)

type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func ErrResponse(err error) ErrHTTP {

	switch err.(type) {
	case account.ErrUserHasAccount:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

	case account.ErrAccountAccess:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusForbidden,
		}

	case account.ErrAmountBelowMinimum:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		}

	case account.ErrNotEnoughBalance:
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
