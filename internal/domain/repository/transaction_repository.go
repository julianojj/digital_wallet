package repository

import "github.com/julianojj/digital_wallet/internal/domain/entity"

type TransactionRepository interface {
	Save(transaction *entity.Transaction) error
	FindByAccount(accountID string) ([]*entity.Transaction, error)
}
