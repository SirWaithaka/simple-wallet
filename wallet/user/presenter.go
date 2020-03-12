package user

import "wallet/data"

type SignedUser struct {
	UserID string `json:"userId"`
	Token string `json:"token"`
}

func RegistrationResponse(user *User) map[string]interface{} {
	return map[string]interface{}{
		"status": "success",
		"message": "user created",
		"user": map[string]interface{} {
			"userId": user.ID,
			"email": user.Email,
		},
	}
}

func parseToNewUser(user User) data.UserContract {
	return data.UserContract{
		UserID: user.ID,
	}
}