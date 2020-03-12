package transaction

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"wallet/data"
)

func parseToTransaction(newTx data.NewTransaction) *Transaction {
	return &Transaction{
		ID:        uuid.NewV4(),
		Type:      newTx.TxType,
		Timestamp: time.Now(),
		Amount:    newTx.Amount,
		UserID:    newTx.UserID,
		AccountID: newTx.AccountID,
	}
}
