package users

import (
	"fmt"
	"net/http"

	"simple-wallet/app/models"
	"simple-wallet/app/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
)

func createUserObject(params RegistrationParams) *models.User {
	id, _ := uuid.NewV4()

	var userObj = models.User{
		ID:          id,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		Email:       params.Email,
		PhoneNumber: params.PhoneNumber,
		Password:    params.Password,
		PassportNo:  params.PassportNo,
	}
	return &userObj

}

func Authenticate(userDomain user.Interactor) fiber.Handler {

	return func(ctx *fiber.Ctx) error {
		var params LoginParams
		_ = ctx.BodyParser(&params)

		if params.Email == "" && params.PhoneNumber == "" {
			err := ErrResponse(ErrInvalidParams{
				message: fmt.Sprintf("provide email or phone number to sign in."),
			})
			return ctx.Status(err.Status).JSON(err)
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
			return ctx.Status(err.Status).JSON(err)
		}

		// return token
		_ = ctx.Status(http.StatusOK).JSON(signedUser)

		return nil
	}
}

func Register(userDomain user.Interactor) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		var params RegistrationParams
		_ = ctx.BodyParser(&params)

		_, err := ValidateRegisterParams(&params)
		if err != nil {
			errHTTP := ErrResponse(err)
			return ctx.Status(errHTTP.Status).JSON(errHTTP)
		}

		newUser := createUserObject(params)
		// register user
		u, err := userDomain.Register(newUser)
		if err != nil {
			errHTTP := ErrResponse(err)
			return ctx.Status(errHTTP.Status).JSON(errHTTP)
		}

		// we use a presenter to reformat the response of user.
		_ = ctx.JSON(user.RegistrationResponse(&u))

		return nil
	}
}
