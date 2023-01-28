package database

import (
	"database/sql"

	"github.com/julianojj/digital_wallet/internal/domain/entity"
)

type TransactionRepositoryDatabase struct {
	DB *sql.DB
}

func NewTransactionRepositoryDatabase(db *sql.DB) *TransactionRepositoryDatabase {
	return &TransactionRepositoryDatabase{
		DB: db,
	}
}

func (t *TransactionRepositoryDatabase) Save(transaction *entity.Transaction) error {
	_, errTransaction := t.DB.Exec("INSERT INTO Transactions (Id, AccountId, Type, Amount) VALUES($1, $2, $3, $4)", transaction.ID, transaction.AccountID, transaction.Type, transaction.Amount)
	if errTransaction != nil {
		return errTransaction
	}
	return nil
}

func (t *TransactionRepositoryDatabase) FindByAccount(accountID string) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction
	rows, err := t.DB.Query("SELECT Id, AccountId, Type, Amount FROM Transactions WHERE AccountId = $1", accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Type, &transaction.Amount)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, &transaction)
	}
	return transactions, nil
}
