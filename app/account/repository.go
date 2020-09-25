package account

import (
	"errors"
	"fmt"
	"time"

	"simple-wallet/app/models"
	"simple-wallet/app/storage"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(userId uuid.UUID) (models.Account, error)
	GetBalance(userId uuid.UUID) (*models.Account, error)
	Deposit(userId uuid.UUID, amount uint) (*models.Account, error)
	Withdraw(userId uuid.UUID, amount uint) (*models.Account, error)
}

type repository struct {
	database *storage.Database
}

func NewRepository(db *storage.Database) Repository {
	return &repository{database: db}
}

// Create a now account for userId
func (r repository) Create(userId uuid.UUID) (models.Account, error) {
	// check if user has no account already
	var acc models.Account
	result := r.database.Where(models.Account{UserID: userId}).First(&acc)
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// if a record found we return custom error
		err := ErrUserHasAccount{userId: acc.UserID.String(), accountId: acc.ID.String()}
		return models.Account{}, err
	}

	// create account
	newAcc := zeroAccount(userId)
	result = r.database.Where(models.Account{UserID: userId}).Assign(newAcc).FirstOrCreate(&acc)
	if err := result.Error; err != nil {
		return models.Account{}, NewErrUnexpected(err)
	}

	return acc, nil
}

// GetBalance for account with userId
func (r repository) GetBalance(userId uuid.UUID) (*models.Account, error) {
	acc, err := r.isAccAccessible(userId)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

// Deposit amount into account with userId
func (r repository) Deposit(userId uuid.UUID, amount uint) (*models.Account, error) {
	acc, err := r.isAccAccessible(userId)
	if err != nil {
		return nil, err
	}

	// update balance with amount: add amount
	amtF := acc.Balance + float64(amount)
	result := r.database.Model(acc).Updates(models.Account{Balance: amtF})
	if err = result.Error; err != nil {
		return nil, NewErrUnexpected(err)
	}

	return acc, nil
}

// Withdraw amount from account with userId
func (r repository) Withdraw(userId uuid.UUID, amount uint) (*models.Account, error) {
	acc, err := r.isAccAccessible(userId)
	if err != nil {
		return nil, err
	}

	// check that balance is more than amount
	if acc.Balance < float64(amount) {
		return nil, ErrNotEnoughBalance{
			Message: fmt.Sprintf("cannot withdraw %v, your account has %v", amount, acc.Balance),
		}
	}

	// update balance with amount: subtract amount
	amtF := acc.Balance - float64(amount)
	result := r.database.Model(acc).Updates(models.Account{Balance: amtF})
	if err = result.Error; err != nil {
		return nil, NewErrUnexpected(err)
	}

	return acc, nil
}

func (r repository) isAccAccessible(userId uuid.UUID) (*models.Account, error) {
	var acc models.Account
	result := r.database.Where(models.Account{UserID: userId}).First(&acc)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err := ErrAccountAccess{reason: "Not Created. Report issue"}
		return nil, err
	}

	if acc.Status == models.StatusFrozen || acc.Status == models.StatusSuspended {
		return nil, ErrAccountAccess{reason: acc.Status}
	}

	return &acc, nil
}

func zeroAccount(userId uuid.UUID) *models.Account {
	id, _ := uuid.NewV4()

	return &models.Account{
		ID:              id,
		Balance:         0,
		Status:          models.StatusActive,
		AccountType:     models.TypeCurrent,
		UserID:          userId,
		CreationDate:    time.Now(),
		LastUpdatedDate: time.Now(),
	}
}
