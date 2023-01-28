package memory

import (
	"github.com/julianojj/digital_wallet/internal/domain/entity"
)

type TransactionRepositoryMemory struct {
	Transactions []*entity.Transaction
}

func NewTransactionRepositoryMemory() *TransactionRepositoryMemory {
	return &TransactionRepositoryMemory{
		Transactions: make([]*entity.Transaction, 0),
	}
}

func (t *TransactionRepositoryMemory) Save(transaction *entity.Transaction) error {
	t.Transactions = append(t.Transactions, transaction)
	return nil
}

func (t *TransactionRepositoryMemory) FindByAccount(accountID string) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	for _, transaction := range t.Transactions {
		if transaction.AccountID == accountID {
			transactions = append(transactions, transaction)
		}
	}
	return transactions, nil
}
