package registry

import (
	"wallet/storage"
	"wallet/user"
)

type Domain struct {
	User user.Interactor
}

func NewDomain(database *storage.Database) *Domain {
	userRepo := user.NewRepository(database)

	return &Domain{
		User: user.NewInteractor(userRepo),
	}
}