package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// AccountStatus (active,dormant,frozen,suspended)
type AccountStatus string

// AccountType (savings,current,utility)
type AccountType string

const (
	StatusActive    = AccountStatus("active")
	StatusDormant   = AccountStatus("dormant")
	StatusFrozen    = AccountStatus("frozen")
	StatusSuspended = AccountStatus("suspended")
)

const (
	// different types of accounts a user could hold
	// we will use current account only.
	TypeSavings = AccountType("savings")
	TypeCurrent = AccountType("current")
	TypeUtility = AccountType("utility")
)

// Account entity definition
type Account struct {
	ID          uuid.UUID     `json:"accountId"`
	Balance     float64       `json:"balance"`
	Status      AccountStatus `json:"status"`
	AccountType AccountType   `json:"accountType"`
	UserID      uuid.UUID     `json:"userId"`

	CreationDate    time.Time `json:"creationDate"`
	LastUpdatedDate time.Time `json:"lastUpdatedDate"`
}
