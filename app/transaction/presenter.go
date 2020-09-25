package transaction

import (
	"time"

	"simple-wallet/app/data"

	"github.com/gofrs/uuid"
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
