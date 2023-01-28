package entity

import (
	"time"

	"github.com/julianojj/digital_wallet/internal/application/exceptions"
)

type Transaction struct {
	ID        string
	AccountID string
	Type      string
	Amount    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransaction(
	id string,
	accountID string,
	transactionType string,
	amount float64,
	createdAt time.Time,
	updatedAt time.Time,
) (*Transaction, error) {
	transaction := &Transaction{
		ID:        id,
		AccountID: accountID,
		Type:      transactionType,
		Amount:    amount,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.ID == "" {
		return exceptions.NewValidationError("transaction id is required")
	}
	if t.AccountID == "" {
		return exceptions.NewValidationError("account id is required")
	}
	if t.Type == "" {
		return exceptions.NewValidationError("transaction type is required")
	}
	if t.Amount <= 0 {
		return exceptions.NewValidationError("amount must be greater than zero")
	}
	return nil
}
