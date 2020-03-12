package users

import (
	"fmt"
	"github.com/gofiber/fiber"
	uuid "github.com/satori/go.uuid"
	"net/http"

	"wallet/user"
)

func createUserObject(params RegistrationParams) *user.User {

	var userObj = user.User{
		ID:          uuid.NewV4(),
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
		Password:    params.Password,
		PassportNo:  params.PassportNo,
	}
	return &userObj

}

func Authenticate(userDomain user.Interactor) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {
		var params LoginParams
		_ = ctx.BodyParser(&params)

		if params.Email == "" && params.PhoneNumber ==  "" {
			err := ErrResponse(ErrInvalidParams{
				message: fmt.Sprintf("provide email or phone number to sign in."),
			})
			_  = ctx.Status(err.Status).JSON(err)
			return
		}

		var authErr error
		var signedUser user.SignedUser
		switch {
		case params.Email != "":
			// if email parameter not empty authenticate by email.
			signedUser, authErr = userDomain.AuthenticateByEmail(params.Email, params.Password)
		case params.PhoneNumber != "":
			// if phone number parameter not empty authenticate by phone number.
			signedUser, authErr = userDomain.AuthenticateByPhoneNumber(params.PhoneNumber, params.Password)
		default:
			authErr = nil
		}

		// if there is an error authenticating user.
		if authErr != nil {
			err := ErrResponse(authErr)
			_ = ctx.Status(err.Status).JSON(err)
			return
		}

		// return token
		_ = ctx.Status(http.StatusOK).JSON(signedUser)
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
