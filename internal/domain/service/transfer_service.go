package service

import (
	"errors"

	"github.com/julianojj/digital_wallet/internal/domain/entity"
	"github.com/julianojj/digital_wallet/internal/domain/repository"
)

type TransferService struct {
	TransactionRepository repository.TransactionRepository
}

func NewTransferService(
	transactionRepository repository.TransactionRepository,
) *TransferService {
	return &TransferService{
		TransactionRepository: transactionRepository,
	}
}

func (t *TransferService) Transfer(from *entity.Account, to *entity.Account, amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if from.GetBalance() < amount {
		return errors.New("account no have enough balance")
	}
	from.Debit(amount)
	to.Credit(amount)
	for _, transaction := range from.Transactions {
		t.TransactionRepository.Save(transaction)
	}
	for _, transaction := range to.Transactions {
		t.TransactionRepository.Save(transaction)
	}
	return nil
}
