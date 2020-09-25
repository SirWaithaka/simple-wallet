package account

import (
	"github.com/gofrs/uuid"
	"time"

	"wallet/data"
)

func parseTransactionDetails(userId uuid.UUID, acc Account, txType string, timestamp time.Time) *data.TransactionContract {
	return &data.TransactionContract{
		UserID:    userId,
		AccountID: acc.ID,
		Amount:    acc.Balance,
		TxType:    txType,
		Timestamp: timestamp,
	}
}
