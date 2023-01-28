package usecases

import (
	"errors"
	"time"

	"github.com/julianojj/digital_wallet/internal/domain/repository"
)

type GetAccount struct {
	AccountRepository repository.AccountRepository
}

type GetAccountOutput struct {
	AccountID    string        `json:"account_id"`
	Balance      float64       `json:"balance"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	ID        string  `json:"id"`
	Operation string  `json:"operation"`
	Amount    float64 `json:"amount"`
	Date      string  `json:"date"`
}

func NewGetAccount(accountRepository repository.AccountRepository) *GetAccount {
	return &GetAccount{
		AccountRepository: accountRepository,
	}
}

func (g *GetAccount) Execute(accountID string) (*GetAccountOutput, error) {
	account, err := g.AccountRepository.FindById(accountID)
	if err != nil {
		return nil, errors.New("account not found")
	}
	var transactions []Transaction
	for _, transaction := range account.Transactions {
		transactions = append(transactions, Transaction{
			ID:        transaction.ID,
			Operation: transaction.Type,
			Amount:    transaction.Amount,
			Date:      transaction.CreatedAt.Format(time.RFC3339),
		})
	}
	return &GetAccountOutput{
		AccountID:    account.ID,
		Balance:      account.GetBalance(),
		Transactions: transactions,
	}, nil
}
