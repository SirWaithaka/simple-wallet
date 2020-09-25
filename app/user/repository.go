package user

import (
	"errors"
	"fmt"
	"log"

	"simple-wallet/app/models"
	"simple-wallet/app/storage"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Add(models.User) (models.User, error)
	Delete(models.User) error
	GetByID(uuid.UUID) (models.User, error)
	GetByEmail(string) (models.User, error)
	GetByPhoneNumber(string) (models.User, error)
	Update(models.User) error
}

func NewRepository(db *storage.Database) Repository {
	return &repository{database: db}
}

type repository struct {
	database *storage.Database
}

// Add a user if already not in db.
func (r repository) Add(user models.User) (models.User, error) {
	var u models.User

	// check if user does not exist by email and phone number
	result := r.database.Where(models.User{Email: user.Email}).Or(models.User{PhoneNumber: user.PhoneNumber}).First(&u)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("user %v", u)
		return u, &ErrUserExists{inUser: user, outUser: u}
	}
	// add user to db with given email
	result = r.database.Where(models.User{Email: user.Email}).Assign(user).FirstOrCreate(&u)
	if err := result.Error; err != nil {
		return models.User{}, NewErrUnexpected(err)
	}

	return u, nil
}

func (r repository) Delete(user models.User) error {
	result := r.database.Delete(&user)
	return NewErrUnexpected(result.Error)
}

func (r repository) GetByID(id uuid.UUID) (models.User, error) {
	var user models.User

	result := r.database.Where("id = ?", id.String()).First(&user)

	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		msg := fmt.Sprintf("user with id %v not found", id.String())
		return models.User{}, ErrUserNotFound{message: msg}
	}

	if err := result.Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r repository) GetByEmail(email string) (models.User, error) {
	var user models.User

	// perform query
	result := r.database.Where(models.User{Email: email}).First(&user)

	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		msg := fmt.Sprintf("user with email %v not found", email)
		return models.User{}, ErrUserNotFound{message: msg}
	}

	// if any other error, return.
	if err := result.Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r repository) GetByPhoneNumber(phoneNo string) (models.User, error) {
	var user models.User

	result := r.database.Where(models.User{PhoneNumber: phoneNo}).First(&user)

	// check if no record found.
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		msg := fmt.Sprintf("user with phone number %v not found", phoneNo)
		return models.User{}, ErrUserNotFound{message: msg}
	}

	if err := result.Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r repository) Update(user models.User) error {
	var u models.User
	result := r.database.Model(&u).Omit("id").Updates(user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
