package account

import (
	uuid "github.com/satori/go.uuid"
	"time"

	"wallet/data"
)

func parseTransactionDetails(userId uuid.UUID, acc Account, txType string, timestamp time.Time) *data.NewTransaction {
	return &data.NewTransaction{
		UserID:    userId,
		AccountID: acc.ID,
		Amount:    acc.Balance,
		TxType:    txType,
		Timestamp: timestamp,
	}
}
