package users

import (
	"encoding/json"
	"fmt"
	"log"
)

type LoginParams struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

type RegistrationParams struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	PassportNo  string `json:"passportNumber"`
	Password    string `json:"password"`
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

