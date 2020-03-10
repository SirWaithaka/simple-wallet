package user

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
