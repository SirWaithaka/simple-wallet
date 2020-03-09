package user

import (
	"log"

	uuid "github.com/satori/go.uuid"

	"wallet/storage"
)


type Repository interface {
	Add(User) (User, error)
	Delete(User) error
	GetByID(id uuid.UUID) (User, error)
	GetByEmail(email string) (User, error)
	Update(User) error
}


func NewRepository(db *storage.Database) Repository {
	return &repository{database:db}
}

type repository struct {
	database *storage.Database
}

// Add a user if already not in db.
func (r repository) Add(user User) (User, error) {
	var u User

	// check if user does not exist by email and phone number
	var notExists bool
	notExists = r.database.Where(User{Email: user.Email}).Or(User{PhoneNumber: user.PhoneNumber}).First(&u).RecordNotFound()
	if !notExists {
		log.Printf("user %v", u)
		return u, &ErrUserExists{inUser: user, outUser: u}
	}
	// add user to db with given email
	result := r.database.Where(User{Email: user.Email,}).Assign(user).FirstOrCreate(&u)
	if err := result.Error; err != nil {
		return User{}, NewErrUnexpected(err)
	}

	return u, nil
}

func (r repository) Delete(user User) error {
	result := r.database.Delete(&user)
	return result.Error
}

func (r repository) GetByID(id uuid.UUID) (User, error) {
	var user User
	result := r.database.Where("id = ?", id.String()).First(&user)
	if err := result.Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r repository) GetByEmail(email string) (User, error) {
	var user User
	result := r.database.Where("email = ?", email).First(&user)
	if err := result.Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r repository) Update(user User) error {
	var u User
	result := r.database.Model(&u).Omit("id").Update(user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
