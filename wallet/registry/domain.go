package registry

import (
	"wallet"
	"wallet/account"
	"wallet/storage"
	"wallet/transaction"
	"wallet/user"
)

type Domain struct {
	Account account.Interactor
	Transaction transaction.Interactor
	User user.Interactor
}

func NewDomain(config wallet.Config, database *storage.Database, channels *Channels) *Domain {
	accRepo := account.NewRepository(database)
	userRepo := user.NewRepository(database)
	txnRepo := transaction.NewRepository(database)

	return &Domain{
		Account: account.NewInteractor(accRepo, channels.ChannelNewUsers, channels.ChannelNewTransactions),
		Transaction: transaction.NewInteractor(txnRepo, channels.ChannelNewTransactions),
		User: user.NewInteractor(config, userRepo, channels.ChannelNewUsers),
	}
}