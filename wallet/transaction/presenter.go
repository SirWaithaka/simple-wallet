package transaction

import (
	"time"

	"github.com/gofrs/uuid"

	"wallet/data"
)

func parseToTransaction(newTx data.TransactionContract) *Transaction {
	id, _ := uuid.NewV4()

	return &Transaction{
		ID:        id,
		Type:      newTx.TxType,
		Timestamp: time.Now(),
		Amount:    newTx.Amount,
		UserID:    newTx.UserID,
		AccountID: newTx.AccountID,
	}
}
