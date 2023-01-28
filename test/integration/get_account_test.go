package integration

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/digital_wallet/internal/application/usecases"
	"github.com/julianojj/digital_wallet/internal/domain/entity"
	"github.com/julianojj/digital_wallet/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldGetAccountIfAccountNotFound(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
	getAccount := usecases.NewGetAccount(accountRepository)
	account, err := getAccount.Execute("a188a603-1725-47fb-9f6a-5b1ec0aa91cc")
	assert.EqualError(t, err, "account not found")
	assert.Nil(t, account)
}

func TestShouldGetAccount(t *testing.T) {
	accountRepository := memory.NewAccountRepositoryMemory()
	currentTime := time.Now()
	account, _ := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	account.Credit(1000.00)
	accountRepository.Save(account)
	getAccount := usecases.NewGetAccount(accountRepository)
	accountOutput, err := getAccount.Execute(account.ID)
	assert.NoError(t, err)
	assert.Equal(t, account.ID, accountOutput.AccountID)
	assert.Equal(t, 1000.00, accountOutput.Balance)
}
