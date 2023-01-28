package memory

import (
	"errors"

	"github.com/julianojj/digital_wallet/internal/domain/entity"
)

type AccountRepositoryMemory struct {
	Accounts []*entity.Account
}

func NewAccountRepositoryMemory() *AccountRepositoryMemory {
	return &AccountRepositoryMemory{
		Accounts: make([]*entity.Account, 0),
	}
}

func (a *AccountRepositoryMemory) Save(account *entity.Account) error {
	a.Accounts = append(a.Accounts, account)
	return nil
}

func (a *AccountRepositoryMemory) FindById(accountID string) (*entity.Account, error) {
	for _, account := range a.Accounts {
		if account.ID == accountID {
			return account, nil
		}
	}
	return nil, errors.New("account not found")
}
