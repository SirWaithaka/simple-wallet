package user


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
