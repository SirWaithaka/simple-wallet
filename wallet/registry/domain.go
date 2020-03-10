package registry

import (
	"wallet"
	"wallet/storage"
	"wallet/user"
)

type Domain struct {
	User user.Interactor
}

func NewDomain(config wallet.Config, database *storage.Database) *Domain {
	userRepo := user.NewRepository(database)

	return &Domain{
		User: user.NewInteractor(config, userRepo),
	}
}