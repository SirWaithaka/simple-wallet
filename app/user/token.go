package user

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenDetails struct {
	UserId string `json:"userId"`
	Email  string `json:"email"`
}

type TokenClaims struct {
	User TokenDetails `json:"user"`

	jwt.StandardClaims
}

func generateToken(user *User) *jwt.Token {

	expirationTime := time.Now().Add(6 * time.Hour).Unix()
	issuedAt := time.Now().Unix()
	claims := TokenClaims{
		User: TokenDetails{
			UserId: user.ID.String(),
			Email:  user.Email,
		},

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
			IssuedAt:  issuedAt,
		},
	}

	// We build a token, we give it and expiry of 6 hours.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token
}

func getTokenString(secret string, token *jwt.Token) (string, error) {
	str, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", ErrTokenParsing{message: err.Error()}
	}
	return str, nil
}

func ParseToken(token, secret string, claims *TokenClaims) (*jwt.Token, error) {
	tok, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	return tok, err
}

func ValidateToken(tok *jwt.Token) bool {
	if !tok.Valid {
		return false
	}
	return true
}
