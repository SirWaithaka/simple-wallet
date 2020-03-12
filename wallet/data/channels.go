package data

import uuid "github.com/satori/go.uuid"

type NewRegisteredUser struct {
	UserID uuid.UUID
}

type ChanNewUsers struct {
	Channel chan NewRegisteredUser
	Reader <-chan NewRegisteredUser
	Writer chan<- NewRegisteredUser
}
