package account

import (
	"time"

	"github.com/gofrs/uuid"
)

const (
	StatusActive    = "active"
	StatusFrozen    = "frozen"
	StatusSuspended = "suspended"
)

const (
	// different types of accounts a user could hold
	// we will use current account only.
	TypeSavings = "savings"
	TypeCurrent = "current"
	TypeUtility = "utility"
)

type Account struct {
	ID          uuid.UUID `json:"accountId"`
	Balance     float64   `json:"balance"`
	Status      string    `json:"status"`
	AccountType string    `json:"accountType"`
	UserID      uuid.UUID `json:"userId"`

	CreationDate    time.Time `json:"creationDate"`
	LastUpdatedDate time.Time `json:"lastUpdatedDate"`
}
