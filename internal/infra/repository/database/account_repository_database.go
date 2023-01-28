package database

import (
	"database/sql"

	"github.com/julianojj/digital_wallet/internal/domain/entity"
)

type AccountRepositoryDatabase struct {
	DB *sql.DB
}

func NewAccountRepositoryDatabase(db *sql.DB) *AccountRepositoryDatabase {
	return &AccountRepositoryDatabase{
		DB: db,
	}
}

func (r *AccountRepositoryDatabase) Save(account *entity.Account) error {
	_, err := r.DB.Exec("INSERT INTO Accounts(Id, BankCode, Branch, AccountNumber) VALUES ($1, $2, $3, $4)", account.ID, account.BankCode, account.Branch, account.AccountNumber)
	return err
}

func (r *AccountRepositoryDatabase) FindById(accountID string) (*entity.Account, error) {
	var account entity.Account
	errAccount := r.DB.QueryRow("SELECT Id, BankCode, Branch, AccountNumber, CreatedAt, UpdatedAt FROM Accounts WHERE Id = $1", accountID).Scan(&account.ID, &account.BankCode, &account.Branch, &account.AccountNumber, &account.CreatedAt, &account.UpdatedAt)
	if errAccount != nil {
		return nil, errAccount
	}
	rows, errTransactions := r.DB.Query("select Id, AccountId, Type, Amount FROM Transactions WHERE AccountId = $1", account.ID)
	if errTransactions != nil {
		return nil, errTransactions
	}
	for rows.Next() {
		var transaction entity.Transaction
		errTransaction := rows.Scan(&transaction.ID, &transaction.AccountID, &transaction.Type, &transaction.Amount)
		if errTransaction != nil {
			return nil, errTransaction
		}
		account.Transactions = append(account.Transactions, &transaction)
	}
	return &account, nil
}
