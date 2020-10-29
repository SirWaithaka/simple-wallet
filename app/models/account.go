package models

import (
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
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
	AccTypeSavings = AccountType("savings")
	AccTypeCurrent = AccountType("current")
	AccTypeUtility = AccountType("utility")
)

// Account entity definition
type Account struct {
	gorm.Model

	ID          uuid.UUID
	Balance     float64       `gorm:"column:balance"`
	Status      AccountStatus `gorm:"column:status"`
	AccountType AccountType   `gorm:"column:account_type"`
	UserID      uuid.UUID     `gorm:"column:user_id;not null;unique"` // a user can only have one account
}
