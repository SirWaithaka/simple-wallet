package models

import (
	"github.com/gofrs/uuid"
)

// Names according to national id/passport,
// id/passport no, D.O.B, creation date,
// last update date, email,phone, Password/PIN,
// Country (So you need a country table),
// reset pin/pass count, max reset pin/pass count,
// account status(active,dormant,frozen,suspended etc),etc
type User struct {
	ID          uuid.UUID `json:"userId" gorm:"primary_key"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email" gorm:"not null;unique"`
	PhoneNumber string    `json:"phoneNumber" gorm:"not null;unique"`
	PassportNo  string    `json:"passportNumber"`
	Password    string    `json:"password" gorm:"not null"`
}
