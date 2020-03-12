package user

import (
	"fmt"
	"log"
	"wallet"
	"wallet/data"

	"golang.org/x/crypto/bcrypt"
)

type Interactor interface {
	AuthenticateByEmail(email, password string) (SignedUser, error)
	AuthenticateByPhoneNumber(email, password string) (SignedUser, error)
	Register(*User) (User, error)
}

func NewInteractor(config wallet.Config, userRepo Repository, usersChan data.ChanNewUsers) Interactor {
	return &interactor{
		config: config,
		repository: userRepo,
		userChannel: usersChan,
	}
}

type interactor struct {
	userChannel data.ChanNewUsers
	config wallet.Config
	repository Repository
}

// AuthenticateByEmail verifies a user by provided unique email address
func (ui interactor) AuthenticateByEmail(email, password string) (SignedUser, error) {
	// search for user by email.
	user, err := ui.repository.GetByEmail(email)
	if err != nil {
		return SignedUser{}, err
	}

	// if user exists authenticate them.
	token, err := ui.authenticateUser(&user, password)
	if err != nil {
		return SignedUser{}, err
	}

	return SignedUser{ UserID: user.ID.String(), Token: token}, nil
}

// AuthenticateByPhoneNumber verifies a user by provided unique phone number
func (ui interactor) AuthenticateByPhoneNumber(phoneNo, password string) (SignedUser, error) {
	// search for user by phone number.
	user, err := ui.repository.GetByPhoneNumber(phoneNo)
	if err != nil {
		return SignedUser{}, err
	}

	// if user exists authenticate them.
	token, err := ui.authenticateUser(&user, password)
	if err != nil {
		return SignedUser{}, err
	}

	return SignedUser{ UserID: user.ID.String(), Token: token}, nil
}

// given a user object and password  verify password then generate
// an auth token string and return.
func (ui interactor) authenticateUser(user *User, password string) (string, error) {
	// validate password
	if isValidPassword := ui.comparePasswords(user.Password, []byte(password)); !isValidPassword {
		msg := fmt.Sprintf("user or password invalid!")
		return "", ErrUnauthorized{Message: msg}
	}

	// create a token
	tok := generateToken(user)
	token, err := getTokenString(ui.config.Secret, tok)
	if err != nil {
		return "", err
	}

	return token, nil
}

// verify hashed password and plain text password if they match.
func (ui interactor) comparePasswords(hashedPassword string, password []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, password)
	if err != nil {
		log.Printf("err comparing passwords %v", err)
		return false
	}
	return true
}

// take a plain text password and hash it.
func (ui interactor) hashUserPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		hashErr := ErrHashPassword{password: string(password), message: err.Error()}
		return "", &hashErr
	}
	return string(hash), nil
}

// Register takes in a user object and adds the user to db.
func (ui interactor) Register(user *User) (User, error) {
	// hash user password before adding to db.
	passwordHash, err := ui.hashUserPassword([]byte(user.Password))
	if err != nil {
		return User{}, err
	}

	// change password to hashed string
	user.Password = passwordHash
	u, err :=  ui.repository.Add(*user)
	if err != nil {
		return User{}, err
	}

	// tell channel listeners that a new user has been created.
	ui.postNewUserToChannel(&u)
	return u, nil
}

// take the newly created user and post them to channel
// that listens for newly created user and acts upon them
// like creating an account for them automatically.
func (ui interactor) postNewUserToChannel(user *User) {
	newUser := parseToNewUser(*user)
	go func() { ui.userChannel.Writer <- newUser }()
}