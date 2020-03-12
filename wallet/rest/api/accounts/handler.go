package account

import (
	"github.com/gofiber/fiber"
	"wallet/account"
)

// Deposit allows user to deposit or credit their
// account.
func Deposit(interactor account.Interactor) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {

	}
}

// Withdraw allows user to withdraw or debit their
// account.
func Withdraw(interactor account.Interactor) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {

	}
}

// MiniStatement returns a small short summary of the
// most recent transactions on an account.
func MiniStatement() func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {

	}
}
