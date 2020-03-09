package user

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"wallet/user"
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
			Status:  400,
		}

	case *user.ErrUserExists:
		var e = err.(*user.ErrUserExists)
		errHTTP = ErrHTTP{
			Error:   reflect.TypeOf(*e).Name(),
			Message: err.Error(),
			Status:  400,
		}
		
	default:
		errHTTP = ErrHTTP{
			Error:   "ErrUnknown",
			Message: "something unfortunate happened!",
			Status:  500,
		}
	}

	return errHTTP
}

// ValidateRegisterParams checks for required parameters are not zero values.
func ValidateRegisterParams(params *RegistrationParams) (bool, error) {

	// definitely not the best way to convert struct to map
	// performance is really affected for many requests.
	// it works for now.
	var pMap map[string]interface{}
	in, _ := json.Marshal(params)
	_ = json.Unmarshal(in, &pMap)

	var err error
	for k, v := range pMap {
		// ignore passportNumber field in struct.
		if k == "passportNumber" {
			continue
		}

		// if value in map is empty we return an error
		if v == "" {
			log.Printf("checking %v value passed %v\n", k, v)
			err = ErrInvalidParams{message: fmt.Sprintf("%v is required", k)}
			break
		}
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
