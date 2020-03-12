package registry

import (
	"wallet"
	"wallet/account"
	"wallet/storage"
	"wallet/user"
)

type Domain struct {
	Account account.Interactor
	User user.Interactor
}

func NewDomain(config wallet.Config, database *storage.Database, channels *Channels) *Domain {
	accRepo := account.NewRepository(database)
	userRepo := user.NewRepository(database)

	return &Domain{
		Account: account.NewInteractor(accRepo, channels.ChannelNewUsers),
		User: user.NewInteractor(config, userRepo, channels.ChannelNewUsers),
	}
}