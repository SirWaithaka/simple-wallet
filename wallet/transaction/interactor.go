package transaction

import (
	"log"
	"wallet/data"
)

type Interactor interface {
	AddTransaction(Transaction) error
	//GetStatement()
}

type interactor struct {
	repository Repository
	transChannel data.ChanNewTransactions
}

func NewInteractor(repository Repository, transChan data.ChanNewTransactions) Interactor {
	intr := &interactor{
		repository:repository,
		transChannel: transChan,
	}

	go intr.listenOnTransactions()

	return intr
}

func(i interactor) AddTransaction(tx Transaction) error {
	_, err := i.repository.Add(tx)
	if err != nil {
		// if we get an error we are going to add the
		// transaction into a buffer object that will
		// retry adding the transaction at a later time

		return err
	}
	return nil
}

func (i interactor) listenOnTransactions() {
	for {
		select {
		case tx := <- i.transChannel.Reader:
			transaction := parseToTransaction(tx)

			err := i.AddTransaction(*transaction)
			if err != nil {
				e := err.(*ErrUnexpected)
				log.Println(e.Debug())
				return
			}
			log.Printf("Transaction %v has been successfully added.", transaction.ID)
		}
	}
}