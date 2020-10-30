package middleware

import (
	"net/http"
	"reflect"
	"strings"

	"simple-wallet/app/auth"
	"simple-wallet/app/errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrHTTP(err error) ErrHTTP {

	switch err.(type) {
	case errors.Unauthorized:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
	default:
		return ErrHTTP{
			Error:   reflect.TypeOf(err).Name(),
			Message: err.Error(),
			Status:  http.StatusUnauthorized,
		}
	}
}

func AuthByBearerToken(secret string) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		// check that the header is actually set
		header := ctx.Get("Authorization")
		if header == "" {
			err := errors.Unauthorized{Message: "authorization header not set"}
			return ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(err))
		}

		// check that the token value in header is set
		bearer := strings.Split(header, " ")
		if len(bearer) < 2 || bearer[1] == "" {
			err := errors.Unauthorized{Message: "authentication token not set"}
			return ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(err))
		}

		var claims auth.TokenClaims
		token, err := auth.ParseToken(bearer[1], secret, &claims)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				errUnauthorized := errors.Unauthorized{Message: "invalid signature on token"}
				return ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(errUnauthorized))
			}

			errUnauthorized := errors.Unauthorized{Message: "token has expired or is invalid"}
			return ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(errUnauthorized))
		}
		if valid := auth.ValidateToken(token); !valid {
			return ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(errors.Unauthorized{Message: "invalid token"}))
		}

		userDetails := map[string]string{
			"userId": claims.User.UserId,
			"email":  claims.User.Email,
		}
		ctx.Locals("userDetails", userDetails)

		return ctx.Next()
	}
}
