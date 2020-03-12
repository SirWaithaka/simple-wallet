package accounts

import (
	"fmt"
	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
	"wallet/account"
)

type param struct {
	Amount uint `json:"amount"`
}

// Deposit allows user to deposit or credit their
// account.
func Deposit(interactor account.Interactor) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {
		var userDetails = ctx.Locals("userDetails").(map[string]string)
		userId := userDetails["userId"]

		var p param
		_ = ctx.BodyParser(&p)

		balance, err := interactor.Deposit(uuid.FromStringOrNil(userId), p.Amount)
		if err != nil {
			errHTTP := ErrResponse(err)
			_ = ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		_ = ctx.JSON(map[string]interface{}{
			"message": fmt.Sprintf("Amount successfully deposited new balance %v", balance),
			"balance": balance,
			"userId": userId,
		})
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
