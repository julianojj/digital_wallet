package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/digital_wallet/internal/application/exceptions"
)

type Account struct {
	ID            string
	BankCode      string
	Branch        string
	AccountNumber string
	Transactions  []*Transaction
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewAccount(
	id string,
	bankCode string,
	branch string,
	accountNumber string,
	created time.Time,
	updatedAt time.Time,
) (*Account, error) {
	account := &Account{
		ID:            id,
		BankCode:      bankCode,
		Branch:        branch,
		AccountNumber: accountNumber,
		CreatedAt:     created,
		UpdatedAt:     updatedAt,
	}
	err := account.Validate()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (a *Account) Validate() error {
	if a.ID == "" {
		return exceptions.NewValidationError("account id is required")
	}
	if a.BankCode == "" {
		return exceptions.NewValidationError("bank code is required")
	}
	if a.Branch == "" {
		return exceptions.NewValidationError("branch is required")
	}
	if a.AccountNumber == "" {
		return exceptions.NewValidationError("account number is required")
	}
	return nil
}

func (a *Account) GetBalance() float64 {
	balance := 0.00
	for _, transaction := range a.Transactions {
		if transaction.Type == "Credit" {
			balance += transaction.Amount
		}
		if transaction.Type == "Debit" {
			balance -= transaction.Amount
		}
	}
	return balance
}

func (a *Account) Credit(amount float64) error {
	currentTime := time.Now()
	transaction, err := NewTransaction(uuid.NewString(), a.ID, "Credit", amount, currentTime, currentTime)
	if err != nil {
		return err
	}
	a.Transactions = append(a.Transactions, transaction)
	return nil
}

func (a *Account) Debit(amount float64) error {
	currentTime := time.Now()
	transaction, err := NewTransaction(uuid.NewString(), a.ID, "Debit", amount, currentTime, currentTime)
	if err != nil {
		return err
	}
	a.Transactions = append(a.Transactions, transaction)
	return nil
}
