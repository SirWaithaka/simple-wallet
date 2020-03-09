package user

import (

	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"

	"wallet/user"
)

type RegistrationParams struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	PassportNo  string `json:"passportNumber"`
	Password    string `json:"password"`
}


func createUserObject(params RegistrationParams) *user.User {

	var userObj = user.User{
		ID: uuid.NewV4(),
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		PhoneNumber: params.PhoneNumber,
		Password: params.Password,
		PassportNo: params.PassportNo,
	}
	return &userObj

}

func Authenticate(userDomain user.Interactor) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {

	}
}

func Register(userDomain user.Interactor) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {

		var params RegistrationParams
		_ = ctx.BodyParser(&params)

		_, err := ValidateRegisterParams(&params)
		if err != nil {
			errHTTP := ErrResponse(err)
			_ = ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		newUser := createUserObject(params)
		// register user
		u, err := userDomain.Register(newUser)
		if err != nil {
			errHTTP := ErrResponse(err)
			_ = ctx.Status(errHTTP.Status).JSON(errHTTP)
			return
		}

		// we use a presenter to reformat the response of user.
		_ = ctx.JSON(user.RegistrationResponse(&u))
	}
}
