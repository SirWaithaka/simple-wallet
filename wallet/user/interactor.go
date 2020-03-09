package user

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Interactor interface {
	Register(*User) (User, error)
}

func NewInteractor(userRepo Repository) Interactor {
	return &interactor{
		repository: userRepo,
	}
}

type interactor struct {
	repository Repository
}

func (ui interactor) ComparePasswords(hashedPassword string, password []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Printf("err comparing passwords %v", err)
		return false
	}
	return true
}

func (ui interactor) HashUserPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		hashErr := ErrHashPassword{password: string(password), message: err.Error()}
		return "", &hashErr
	}
	return string(hash), nil
}

func (ui interactor) Register(user *User) (User, error) {
	// hash user password before adding to db.
	passwordHash, err := ui.HashUserPassword([]byte(user.Password))
	if err != nil {
		return User{}, err
	}
	user.Password = passwordHash
	return ui.repository.Add(*user)
}
