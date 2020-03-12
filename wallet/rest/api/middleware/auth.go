package middleware

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"

	"wallet/user"
)

type ErrHTTP struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewErrHTTP(err error) ErrHTTP {

	switch err.(type) {
	case user.ErrUnauthorized:
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

func AuthByBearerToken(secret string) func(*fiber.Ctx) {

	return func(ctx *fiber.Ctx) {

		// check that the header is actually set
		header := ctx.Get("Authorization")
		if header == "" {
			err := user.ErrUnauthorized{Message: "authorization header not set"}
			_ = ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(err))
			return
		}

		// check that the token value in header is set
		bearer := strings.Split(header, " ")
		if len(bearer) < 2 || bearer[1] == "" {
			err := user.ErrUnauthorized{Message: "authentication token not set"}
			_ = ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(err))
			return
		}

		var claims user.TokenClaims
		token, err := user.ParseToken(bearer[1], secret, &claims)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				errUnauthorized := user.ErrUnauthorized{Message: "invalid signature on token"}
				_ = ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(errUnauthorized))
				return
			}

			errUnauthorized := user.ErrUnauthorized{Message: "token has expired or is invalid"}
			_ = ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(errUnauthorized))
			return
		}
		if valid := user.ValidateToken(token); !valid {
			_ = ctx.Status(http.StatusUnauthorized).JSON(NewErrHTTP(user.ErrUnauthorized{Message: "invalid token"}))
			return
		}

		userDetails := map[string]string{
			"userId": claims.User.UserId,
			"email": claims.User.Email,
		}
		ctx.Locals("userDetails", userDetails)

		ctx.Next()
	}
}
