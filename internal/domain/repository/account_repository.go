package repository

import "github.com/julianojj/digital_wallet/internal/domain/entity"

type AccountRepository interface {
	Save(account *entity.Account) error
	FindById(accountID string) (*entity.Account, error)
}
