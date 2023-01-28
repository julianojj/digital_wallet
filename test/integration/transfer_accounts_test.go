package integration

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/digital_wallet/internal/application/usecases"
	"github.com/julianojj/digital_wallet/internal/domain/entity"
	"github.com/julianojj/digital_wallet/internal/domain/service"
	"github.com/julianojj/digital_wallet/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldTransferAccountsIfAccountFromNotFound(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
	transactionRepository := memory.NewTransactionRepositoryMemory()
	transferService := service.NewTransferService(transactionRepository)
	transferAccounts := usecases.NewTransferAccounts(accountRepository, transferService)
	err := transferAccounts.Execute(usecases.TransferAccountsInput{
		From:   "a188a603-1725-47fb-9f6a-5b1ec0aa91cc",
		To:     "145695f3-7438-4911-a69d-51074b21c8ec",
		Amount: 100.00,
	})
	assert.EqualError(t, err, "account not found")
}

func TestNotShouldTransferAccountsIfAccountToNotFound(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
	transactionRepository := memory.NewTransactionRepositoryMemory()
	transferService := service.NewTransferService(transactionRepository)
	currentTime := time.Now()
	accountFrom, _ := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	accountFrom.Credit(1000.00)
	accountRepository.Save(accountFrom)
	transferAccounts := usecases.NewTransferAccounts(accountRepository, transferService)
	err := transferAccounts.Execute(usecases.TransferAccountsInput{
		From:   accountFrom.ID,
		To:     "145695f3-7438-4911-a69d-51074b21c8ec",
		Amount: 100.00,
	})
	assert.EqualError(t, err, "account not found")
}

func TestNotShouldTransferAccountsIfAmountMustBeGreaterThanZero(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
	transactionRepository := memory.NewTransactionRepositoryMemory()
	transferService := service.NewTransferService(transactionRepository)
	currentTime := time.Now()
	accountFrom, _ := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	accountFrom.Credit(1000.00)
	accountTo, _ := entity.NewAccount(uuid.NewString(), "260", "0001", "1234567-8", currentTime, currentTime)
	accountRepository.Save(accountFrom)
	accountRepository.Save(accountTo)
	transferAccounts := usecases.NewTransferAccounts(accountRepository, transferService)
	err := transferAccounts.Execute(usecases.TransferAccountsInput{
		From:   accountFrom.ID,
		To:     accountTo.ID,
		Amount: 0.00,
	})
	assert.EqualError(t, err, "amount must be greater than zero")
}

func TestNotShouldTransferAccountsIfAccountFromNoHaveEnoughbalance(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
    transactionRepository := memory.NewTransactionRepositoryMemory()
    transferService := service.NewTransferService(transactionRepository)
	currentTime := time.Now()
	accountFrom, _ := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	accountFrom.Credit(1000.00)
	accountTo, _ := entity.NewAccount(uuid.NewString(), "260", "0001", "1234567-8", currentTime, currentTime)
	accountRepository.Save(accountFrom)
	accountRepository.Save(accountTo)
	transferAccounts := usecases.NewTransferAccounts(accountRepository, transferService)
	err := transferAccounts.Execute(usecases.TransferAccountsInput{
		From:   accountFrom.ID,
		To:     accountTo.ID,
		Amount: 1000.01,
	})
	assert.EqualError(t, err, "account no have enough balance")
}

func TestShouldTransferAccounts(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
    transactionRepository := memory.NewTransactionRepositoryMemory()
    transferService := service.NewTransferService(transactionRepository)
	currentTime := time.Now()
	accountFrom, _ := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	accountFrom.Credit(1000.00)
	accountTo, _ := entity.NewAccount(uuid.NewString(), "260", "0001", "1234567-8", currentTime, currentTime)
	accountRepository.Save(accountFrom)
	accountRepository.Save(accountTo)
	transferAccounts := usecases.NewTransferAccounts(accountRepository, transferService)
	err := transferAccounts.Execute(usecases.TransferAccountsInput{
		From:   accountFrom.ID,
		To:     accountTo.ID,
		Amount: 100.00,
	})
	transactionFrom, _:= transactionRepository.FindByAccount(accountFrom.ID)
	transactionTo, _:= transactionRepository.FindByAccount(accountTo.ID)
	assert.NoError(t, err)
	assert.Len(t, transactionFrom, 2)
	assert.Len(t, transactionTo, 1)
}
