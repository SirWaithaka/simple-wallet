package user

import (
	"fmt"
	"log"
	"wallet"

	"golang.org/x/crypto/bcrypt"
)

type Interactor interface {
	AuthenticateByEmail(email, password string) (SignedUser, error)
	AuthenticateByPhoneNumber(email, password string) (SignedUser, error)
	Register(*User) (User, error)
}

func NewInteractor(config wallet.Config, userRepo Repository) Interactor {
	return &interactor{
		config: config,
		repository: userRepo,
	}
}

type interactor struct {
	config wallet.Config
	repository Repository
}

func (ui interactor) AuthenticateByEmail(email, password string) (SignedUser, error) {
	user, err := ui.repository.GetByEmail(email)
	if err != nil {
		return SignedUser{}, err
	}

	token, err := ui.authenticateUser(&user, password)
	if err != nil {
		return SignedUser{}, err
	}

	return SignedUser{ UserID: user.ID.String(), Token: token}, nil
}

func (ui interactor) AuthenticateByPhoneNumber(phoneNo, password string) (SignedUser, error) {
	user, err := ui.repository.GetByPhoneNumber(phoneNo)
	if err != nil {
		return SignedUser{}, err
	}

	token, err := ui.authenticateUser(&user, password)
	if err != nil {
		return SignedUser{}, err
	}

	return SignedUser{ UserID: user.ID.String(), Token: token}, nil
}

// generate a auth token and return
func (ui interactor) authenticateUser(user *User, password string) (string, error) {
	// validate password
	if isValidPassword := ui.ComparePasswords(user.Password, []byte(password)); !isValidPassword {
		msg := fmt.Sprintf("user or password invalid!")
		return "", ErrUnauthorized{message: msg}
	}

	// create a token
	tok := generateToken(user)
	token, err := getTokenString(ui.config.Secret, tok)
	if err != nil {
		return "", err
	}

	return token, nil
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
