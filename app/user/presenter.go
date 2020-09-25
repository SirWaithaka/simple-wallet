package user

import (
	"simple-wallet/app/data"
	"simple-wallet/app/models"
)

type SignedUser struct {
	UserID string `json:"userId"`
	Token string `json:"token"`
}

func RegistrationResponse(user *models.User) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"message": "user created",
		"user": map[string]interface{} {
			"userId": user.ID,
			"email": user.Email,
		},
	}
}

func parseToNewUser(user models.User) data.UserContract {
	return data.UserContract{
		UserID: user.ID,
	}
}