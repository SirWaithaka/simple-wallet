package transaction

import (
	uuid "github.com/satori/go.uuid"
	"time"

	"wallet/storage"
)

type Repository interface {
	Add(Transaction) (Transaction, error)
	GetTransactions(userId uuid.UUID, from time.Time, limit int) (*[]Transaction, error)
}

type repository struct {
	database *storage.Database
}

func NewRepository(db *storage.Database) Repository {
	return &repository{db}
}

func (r repository) Add(tx Transaction) (Transaction, error) {
	result := r.database.Create(&tx)
	if err := result.Error; err != nil {
		return Transaction{}, NewErrUnexpected(err)
	}

	return tx, nil
}

func (r repository) GetTransactions(userId uuid.UUID, from time.Time, limit int) (*[]Transaction, error) {
	var transactions []Transaction

	result := r.database.Where(
		Transaction{UserID: userId},
	).Where(
		"timestamp <= ?", from,
	).Order("timestamp desc").Limit(limit).Find(&transactions)

	if err := result.Error; err != nil {
		return nil, err
	}

	return &transactions, nil
}
