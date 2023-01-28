package usecases

import (
	"github.com/julianojj/digital_wallet/internal/domain/repository"
	"github.com/julianojj/digital_wallet/internal/domain/service"
)

type TransferAccounts struct {
	AccountRepository repository.AccountRepository
	TransferService   *service.TransferService
}

type TransferAccountsInput struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func NewTransferAccounts(
	accountRepository repository.AccountRepository,
	transferService *service.TransferService,
) *TransferAccounts {
	return &TransferAccounts{
		AccountRepository: accountRepository,
		TransferService:   transferService,
	}
}

func (t *TransferAccounts) Execute(input TransferAccountsInput) error {
	existingAccountFrom, errExintingAccountFrom := t.AccountRepository.FindById(input.From)
	if errExintingAccountFrom != nil {
		return errExintingAccountFrom
	}
	existingAccountTo, errExintingAccountTo := t.AccountRepository.FindById(input.To)
	if errExintingAccountTo != nil {
		return errExintingAccountTo
	}
	err := t.TransferService.Transfer(existingAccountFrom, existingAccountTo, input.Amount)
	if err != nil {
		return err
	}
	return nil
}
